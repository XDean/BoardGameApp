const serverUrl = "https://xdean.cn:8081/"
var accessToken = null
var sessionId = null

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
      if (find)
        accessToken = find.value
      find = e.cookies.find(x => x.name === 'JSESSIONID')
      if (find)
        sessionId = find.value
      if (e.statusCode == 200 && input.success)
        input.success(e)
    },
    fail: input.fail,
    complete: input.complete
  })
}

module.exports = {
  serverUrl: serverUrl,
  request: request
}