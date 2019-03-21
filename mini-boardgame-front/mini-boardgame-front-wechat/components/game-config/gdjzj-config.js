// pages/components/game-config/gdjzj-config.js
Component({
  properties: {},
  data: {
    playerCountList: [6, 7, 8],
    playerCountChoice: 0,
    config: {
      playerCount: 6
    }
  },
  methods: {
    selectPlayerCount: function(e) {
      this.setData({
        playerCountChoice: e.detail.value
      })
    }
  }
})