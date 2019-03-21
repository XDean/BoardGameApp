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
      var count = this.data.playerCountList[e.detail.value];
      this.setData({
        playerCountChoice: e.detail.value,
        'config.playerCount': count
      })
    }
  },
  observers: {
    'config': function(newValue) {
      console.log(this.data.config)
      var index = this.data.playerCountList.findIndex(x => x == newValue.playerCount)
      this.setData({
        playerCountChoice: index
      })
    }
  }
})