//index.js
//获取应用实例
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
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.userInfoReadyCallback = res => {
        this.updateUserInfo(res.userInfo)
      }
    } else {
      // 在没有 open-type=getUserInfo 版本的兼容处理
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
  //事件处理函数
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