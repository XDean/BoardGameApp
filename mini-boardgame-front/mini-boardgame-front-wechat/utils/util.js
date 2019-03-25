const serverUrl = "https://xdean.cn:8081/"
const socketUrl = "wss://xdean.cn:8081/"
var accessToken = null
var sessionId = null

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
  var socket = wrapSocket(input)
  return socket
}

const wrapSocket = input => {
  var data = {
    closed: false,
    authenticated: false,
    header: null,
    onOpens: [],
    onMessages: [],
    onErrors: [],
    onCloses: [],
    messageBeforeAuth: [],
    responseFunctions: {},
    requestId: 0,
    lastMsgMillis: new Date().getTime(),
    heartBeatId: 0,
    originSocket: null,
    retry: input.retry || 100000,
    retryCount: 0,
  }
  var socket = {
    send: function(e) {
      var msg = JSON.stringify({
        id: ++data.requestId,
        attributes: {},
        payload: null,
        ...e
      })
      if (data.authenticated) {
        console.log('send', msg)
        data.originSocket.send({
          data: msg
        })
      } else {
        data.messageBeforeAuth.push(msg)
      }
      return {
        onResponse: func => data.responseFunctions[msg.id] = func
      }
    },
    onOpen: x => {
      if (authenticated)
        x(data.header)
      else
        data.onOpens.push(x)
    },
    onMessage: x => data.onMessages.push(x),
    onError: x => data.onErrors.push(x),
    onClose: x => {
      if (data.closed)
        x()
      else
        data.onCloses.push(x)
    },
    close: function() {
      data.closed = true
      data.originSocket.close()
    },
    connect: function() {
      var originSocket = wx.connectSocket({
        ...input,
        url: socketUrl + input.url,
        header: {
          'Authorization': accessToken
        }
      })
      data.originSocket = originSocket
      originSocket.onOpen(e => {
        data.header = e
        var msg = JSON.stringify({
          id: ++data.requestId,
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
        data.lastMsgMillis = new Date().getTime()
        var msg = JSON.parse(e.data)
        console.log('recieve', msg)
        switch (msg.topic) {
          case 'AUTHENTICATION':
            data.authenticated = true
            data.onOpens.forEach(x => x(data.header))
            data.onOpens = []
            data.messageBeforeAuth.forEach(x => {
              console.log('send', msg)
              data.originSocket.send({
                data: msg
              })
            })
            data.messageBeforeAuth = []
            this.heartBeat()
            break
          case 'BAD_CREDENTIAL':
          case 'WRONG_FORMAT':
          case 'ERROR':
            console.error(msg)
            break
          default:
            if (msg.id) {
              var responseFunc = data.responseFunctions[msg.id]
              if (responseFunc) {
                data.responseFunctions.delete(msg.id)
                responseFunc(msg)
              }
            }
            data.onMessages.forEach(x => x(msg))
        }
      })
      originSocket.onError(x => {
        console.log('error', x)
        data.onErrors.forEach(e => e(x))
        data.originSocket.close()
      })
      originSocket.onClose(x => {
        data.authenticated = false
        data.originSocket = null
        console.log('close', x)
        data.onCloses.forEach(e => e(x))
        clearInterval(data.heartBeatId)
        if (!data.closed) {
          if (data.retry > data.retryCount) {
            console.info('Socket is closed. Reconnect will be attempted in 1 second.', x.reason);
            setTimeout(() => this.connect(), 1000);
          } else {
            console.error('Connect max times, still fail')
          }
        }
      })
    },
    heartBeat: function() {
      data.heartBeatId = setInterval(() => {
        var current = new Date().getTime()
        if (current - data.lastMsgMillis > 15000) {
          this.send({
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
    }
  }
  socket.connect()
  return socket
}

module.exports = {
  serverUrl: serverUrl,
  request: request,
  connectSocket: connectSocket
}