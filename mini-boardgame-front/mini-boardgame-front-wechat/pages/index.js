const app = getApp()

Page({
  data: {
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo')
  },
  onLoad: function() {
    if (app.globalData.userInfo) {
      // user info is ready
    } else if (this.data.canIUse) {
      app.userInfoReadyCallback = res => {
        this.updateUserInfo(res.userInfo)
      }
    } else {
      wx.getUserInfo({
        success: res => {
          this.updateUserInfo(res.userInfo)
        }
      })
    }
  },
  getUserInfo: function(e) {
    this.updateUserInfo(e.detail.userInfo)
  },
  createGame: function() {
    wx.navigateTo({
      url: 'game-center/create-game'
    })
  },
  joinGame: function() {
    wx.navigateTo({
      url: 'game-center/join-game'
    })
  },
  updateUserInfo: function (userInfo) {
    console.log(userInfo)
    this.setData({
      userInfo: userInfo,
      hasUserInfo: true
    })
    app.globalData.userInfo = userInfo
    app.globalData.hasUserInfo = true
  }
})