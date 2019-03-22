// pages/game-center/waiting-room.js
import * as util from '../../utils/util.js'
import * as tooltip from '../../utils/toolTip.js'

const app = getApp()

Page({
  data: {
    room: {},
    gameName: {},
    seatList: []
  },
  onLoad: function() {
    var self = this
    tooltip.init(this)
    util.request({
      url: 'game/room',
      success: function(e) {
        console.log(e.data)
        self.setData({
          room: e.data.room,
          gameName: app.globalData.gameList.find(x => x.id == e.data.room.gameName).name
        })
        self.updateSeatList()
      },
      badRequest: function(e) {
        wx.navigateBack()
      }
    })
  },
  onUnload: function() {
    wx.showLoading({
      title: '退出中...',
      mask: true
    })
    util.request({
      url: 'game/room/exit',
      method: 'POST',
      data: '{}',
      complete: function() {
        wx.hideLoading()
      }
    })
  },
  updateSeatList() {
    var room = this.data.room
    var seatList = []
    for (var i = 0; i < room.playerCount; i++) {
      seatList[i] = room.players.find(x => x.seat == i)
    }
    this.setData({
      seatList: seatList
    })
  },
  clickIcon(e) {
    var player = this.data.room.players[e.currentTarget.id]
    tooltip.showToolTip('info', player.profile.nickname, 2000);
  }
})