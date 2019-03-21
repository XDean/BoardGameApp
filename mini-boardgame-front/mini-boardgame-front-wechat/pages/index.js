const app = getApp()

import * as util from '../utils/util.js'

Page({
  data: {
    userInfo: {},
    hasUserInfo: false,
    notInGame: true,
    canIUse: wx.canIUse('button.open-type.getUserInfo')
  },
  onLoad: function() {
    var self = this
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
    util.request({
      url: 'game/room',
      success: function(e) {
        wx.navigateTo({
          url: 'game-center/waiting-room',
        })
      },
      badRequest: function(e) {
        self.setData({
          notInGame: true
        })
      }
    })
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
  updateUserInfo: function(userInfo) {
    console.log(userInfo)
    this.setData({
      userInfo: userInfo,
      hasUserInfo: true
    })
    app.globalData.userInfo = userInfo
    app.globalData.hasUserInfo = true
  }
})