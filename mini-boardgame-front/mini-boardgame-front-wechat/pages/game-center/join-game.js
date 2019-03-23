// pages/index/join-game/join-game.js

import * as util from '../../utils/util.js'

var roomId = null
var getRoomTaskId = null

Page({
  data: {
    room: null
  },
  onLoad: function(options) {

  },
  inputRoomId: function(e) {
    var self = this
    roomId = e.detail.value
    if (getRoomTaskId)
      clearTimeout(getRoomTaskId)
    getRoomTaskId = setTimeout(() => {
      var id = roomId
      util.request({
        url: 'game/room/search',
        method: 'POST',
        data: {
          'roomId': id
        },
        success: e => {
          console.debug(e.data)
          if (e.data.rooms && e.data.rooms.length == 1) {
            self.setData({
              room: e.data.rooms[0]
            })
          }
        }
      })
    }, 1000)
  },
  joinGame: function(e){
    wx.showLoading({
      title: '加入中...',
      mask: true,
    })
    util.request({
      url: 'game/room/join',
      method: 'POST',
      data: {
        'roomId': roomId
      },
      success: function () {
        wx.hideLoading()
        wx.redirectTo({
          url: 'waiting-room',
        })
      }
    })
  }
})