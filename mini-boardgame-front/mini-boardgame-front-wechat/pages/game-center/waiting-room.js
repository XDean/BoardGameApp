// pages/game-center/waiting-room.js
import * as util from '../../utils/util.js'
import * as tooltip from '../../utils/toolTip.js'

const app = getApp()

Page({
  data: {
    myId: null,
    room: {},
    gameName: {},
    seatList: [],
    socket: {}
  },
  onLoad: function() {
    var self = this
    tooltip.init(this)
    util.request({
      url: 'game/room',
      success: function(e) {
        console.log(e.data)
        self.setData({
          myId: e.data.playerId,
          room: e.data.room,
          gameName: app.globalData.gameList.find(x => x.id == e.data.room.gameName).name
        })
        self.updateSeatList()
        self.connectSocket()
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
  connectSocket() {
    var socket = util.connectSocket({
      url: `game/room/${this.data.room.id}`
    })
    this.data.socket = socket
    socket.onMessage(this.handleMessage)
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
  clickSeat(e) {
    var self = this;
    var seat = Number(e.currentTarget.id)
    var player = this.data.room.players.find(x => x.seat == seat)
    if (player) {
      if (player.id == this.data.myId)
        return
      wx.showActionSheet({
        itemList: [`和${player.profile.nickname}换座位`],
        success: index => {
          self.data.socket.send({
            topic: 'CHANGE_SEAT_REQUEST',
            attributes: {
              TO_SEAT: seat
            }
          })
          tooltip.showToolTip('info', '正在请求换座位', 2000)
        }
      })
    } else {
      this.data.socket.send({
        topic: 'CHANGE_SEAT_REQUEST',
        attributes: {
          TO_SEAT: seat
        }
      })
    }
  },
  handleMessage: function(msg) {
    var self = this;
    switch (msg.topic) {
      case 'PLAYER_JOIN':
        var player = {
          id: msg.attributes.USER_ID,
          seat: msg.attributes.SEAT,
          ready: false,
          profile: {}
        }
        this.data.room.players.push(player)
        this.updateSeatList()
        util.request({
          url: 'user/profile',
          data: {
            id: player.id
          },
          success: e => {
            player.profile = e.data.profile
            self.updateSeatList()
          }
        })
        break;
      case 'PLAYER_EXIT':
        var id = msg.attributes.USER_ID
        var index = this.data.room.players.findIndex(x => x.id == id)
        console.log(this.data.room.players[index])
        if (index) {
          this.data.room.players.splice(index, 1)
          this.updateSeatList()
        }
        break;
      case 'CHANGE_SEAT':
        var fromSeat = msg.attributes.FROM_SEAT
        var toSeat = msg.attributes.TO_SEAT
        var fromPlayer = this.data.room.players.find(x => x.seat == fromSeat)
        var toPlayer = this.data.room.players.find(x => x.seat == toSeat)
        fromPlayer && (fromPlayer.seat = toSeat)
        toPlayer && (toPlayer.seat = fromSeat)
        this.updateSeatList()
        break;
    }
  }
})