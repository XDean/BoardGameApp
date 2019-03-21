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
              util.request({
                url: 'user/profile',
                method: 'POST',
                data: JSON.stringify({
                  profile: {
                    nickname: res.userInfo.nickName,
                    male: res.userInfo.gender,
                    avatarUrl: res.userInfo.avatarUrl
                  }
                })
              })
            }
          })
        }
      }
    })
  },
  globalData: {
    userInfo: null,
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