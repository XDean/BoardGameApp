//app.js
import * as util from '/utils/util.js'

App({
  onLaunch: function () {
    wx.login({
      success: res => {
        util.request({
          url: `public/login-openid?provider=wechat-mbg&token=${res.code}`,
          method: 'post',
          success: function(e){
            util.request({
              url: 'hello'
            })
          }
        })
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          wx.getUserInfo({
            success: res => {
              this.globalData.userInfo = res.userInfo
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
            }
          })
        }
      }
    })
  },
  globalData: {
    userInfo: null,
    accessToken: null,
    gameList: [{
        id: "lc",
        name: "Lost Cities"
      },
      {
        id: "gdjzj",
        name: "古中局"
      }
    ]
  }
})