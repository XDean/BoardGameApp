//app.js
import * as util from '/utils/util.js'

App({
  onLaunch: function() {
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          wx.getUserInfo({
            success: res => {
              this.globalData.userInfo = res.userInfo
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
              wx.login({
                success: loginRes => {
                  util.request({
                    url: `public/login-openid?provider=wechat-mbg&token=${loginRes.code}`,
                    method: 'post',
                    success: function(e) {
                      util.request({
                        url: 'user/profile',
                        data: JSON.stringify({
                          profile: {
                            nickname: res.userInfo.name,
                            male: res.userInfo.male,
                            avatarUrl: res.userInfo.avatarUrl
                          }
                        })
                      })
                    }
                  })
                }
              })
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
        id: "lost-cities",
        name: "Lost Cities"
      },
      {
        id: "gdjzj",
        name: "古中局"
      }
    ]
  }
})