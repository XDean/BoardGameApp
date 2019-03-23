const serverUrl = "https://xdean.cn:8081/"
const socketUrl = "wss://xdean.cn:8081/"
var accessToken = null
var sessionId = null
var sockets = []

const request = input => {
  if (accessToken)
    wx.request({
      ...input,
      url: serverUrl + input.url,
      header: {
        ...input.header,
        ...(accessToken && {
          "Authorization": accessToken
        }),
        ...(sessionId && {
          "Cookie": `JSESSIONID = ${sessionId}`
        })
      },
      method: (input.method ? input.method : 'GET'),
      success: function(e) {
        console.log(e)
        var find = e.cookies.find(x => x.name === 'JSESSIONID')
        if (find)
          sessionId = find.value
        if (e.statusCode == 200 && input.success)
          input.success(e)
        else if (input.badRequest)
          input.badRequest(e)
      }
    })
  else
    wx.login({
      success: res => {
        wx.request({
          url: serverUrl + `public/login-openid?provider=wechat-mbg&token=${res.code}`,
          method: 'POST',
          success: function(e) {
            console.log(e)
            var find = e.cookies.find(x => x.name === 'access-token')
            if (find)
              accessToken = find.value
            find = e.cookies.find(x => x.name === 'JSESSIONID')
            if (find)
              sessionId = find.value
            request(input)
          }
        })
      }
    })
}

const connectSocket = input => {
  var url = input.url
  var cache = sockets.find(x => x.url == url)
  if (cache) {
    cache.handle++;
    return cache.socket
  }
  var originSocket = wx.connectSocket({
    ...input,
    url: socketUrl + url
  })
  var socket = wrapSocket(originSocket)
  sockets.push({
    url: url,
    handle: 1,
    socket: socket
  })
  return socket
}

const wrapSocket = originSocket => {
  var authenticated = false
  var header = null
  var onOpens = []
  var onMessages = []
  var messageBeforeAuth = []
  var responseFunctions = {}
  var requestId = 0
  var lastMsgMillis = new Date().getTime()
  var heartBeatId = 0
  originSocket.onOpen(e => {
    header = e
    var msg = JSON.stringify({
      id: ++requestId,
      topic: 'AUTHENTICATION',
      attributes: {
        ACCESS_TOKEN: accessToken
      },
      payload: null
    })
    originSocket.send({
      data: msg
    })
  })
  originSocket.onMessage(e => {
    lastMsgMillis = new Date().getTime()
    var msg = JSON.parse(e.data)
    console.log('recieve', msg)
    switch (msg.topic) {
      case 'AUTHENTICATION':
        authenticated = true
        onOpens.forEach(x => x(header))
        messageBeforeAuth.forEach(x => {
          console.log('send', msg)
          originSocket.send({
            data: msg
          })
        })
        messageBeforeAuth = []
        break
      case 'BAD_CREDENTIAL':
      case 'WRONG_FORMAT':
      case 'ERROR':
        console.error(msg)
        break
      case 'ROOM_CANCEL':
        originSocket.close({
          reason: 'Room canceled'
        })
        break;
      default:
        if (msg.id) {
          var responseFunc = responseFunctions[msg.id]
          if (responseFunc) {
            responseFunctions.delete(msg.id)
            responseFunc(msg)
          }
        }
        onMessages.forEach(x => x(msg))
    }
  })
  originSocket.onError(x => {
    console.log('error', x)
  })
  originSocket.onClose(x => {
    console.log('close', x)
    clearInterval(heartBeatId)
  })
  var socket = {
    ...originSocket,
    send: e => {
      var msg = JSON.stringify({
        id: ++requestId,
        attributes: {},
        payload: null,
        ...e
      })
      if (authenticated) {
        console.log('send', msg)
        originSocket.send({
          data: msg
        })
      } else {
        messageBeforeAuth.push(msg)
      }
      return {
        onResponse: func => responseFunctions[msg.id] = func
      }
    },
    onOpen: e => onOpens.push(e),
    onMessage: e => onMessages.push(e)
  }
  heartBeatId = setInterval(() => {
    var current = new Date().getTime()
    if (current - lastMsgMillis > 15000) {
      socket.send({
        topic: 'HEART_BEAT',
        attributes: {
          TIMESTAMP: current
        }
      }).onResponse(msg => {
        if (msg && msg.topic == 'HEART_BEAT') {
          var serverTime = msg.attributes.TIMESTAMP
          if (serverTime - current > 5000)
            console.warn('Heart beat response too slow', socket)
        } else {
          console.error('Heart beat error', msg)
        }
      })
    }
  }, 15000)
  return socket
}

module.exports = {
  serverUrl: serverUrl,
  request: request,
  connectSocket: connectSocket
}