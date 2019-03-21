// pages/game-center/waiting-room.js
import * as util from '../../utils/util.js'

const app = getApp()

Page({
  data: {
    room: {},
    gameName: {}
  },
  onLoad: function() {
    var self = this
    util.request({
      url: 'game/room',
      success: function(e) {
        console.log(e.data)
        self.setData({
          room: e.data.room,
          gameName: app.globalData.gameList.find(x => x.id == e.data.room.gameName).name
        })
      },
      badRequest: function(e) {
        wx.navigateBack()
      }
    })
  },
  onUnload: function() {

  }
})