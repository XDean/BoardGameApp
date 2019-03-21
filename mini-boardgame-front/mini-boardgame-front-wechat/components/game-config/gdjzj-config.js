// pages/components/game-config/gdjzj-config.js
Component({
  properties: {
    readOnly: Boolean,
    config: {
      type: Object,
      value: {
        playerCount: 6
      }
    }
  },
  data: {
    playerCountList: [6, 7, 8],
    playerCountChoice: 0
  },
  methods: {
    selectPlayerCount: function(e) {
      this.setData({
        playerCountChoice: e.detail.value
      })
      config.playerCount = playerCountList[e.detail.value]
    }
  }
})