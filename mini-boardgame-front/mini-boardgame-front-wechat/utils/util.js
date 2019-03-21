const serverUrl = "https://xdean.cn:8081/"
var accessToken = null
var sessionId = null

const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : '0' + n
}

const request = input => {
  wx.request({
    url: serverUrl + input.url,
    data: input.data,
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
      var find = e.cookies.find(x => x.name === 'access-token')
      accessToken = find && find.value
      find = e.cookies.find(x => x.name === 'JSESSIONID')
      sessionId = find && find.value
      if (input.success)
        input.success(e)
    },
    fail: input.fail,
    complete: input.complete
  })
}

module.exports = {
  formatTime: formatTime,
  serverUrl: serverUrl,
  request: request
}