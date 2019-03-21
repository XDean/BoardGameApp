// pages/index/create-game.js

const app = getApp()

Page({
  data: {
    userInfo: {},
    gameList: {},
    gameChoice: 0
  },
  onLoad: function(options) {
    this.setData({
      userInfo: app.globalData.userInfo,
      gameList: app.globalData.gameList
    })
  },
  selectGame:function(e) {
    this.setData({
      gameChoice: e.detail.value
    })
  }
})