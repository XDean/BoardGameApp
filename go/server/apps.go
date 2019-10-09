package main

import (
	"github.com/xdean/miniboardgame/go/games/lostcities"
	"github.com/xdean/miniboardgame/go/games/rps"
	"github.com/xdean/miniboardgame/go/server/game"
)

func init() {
	game.Register(lostcities.Instance)
	game.Register(rps.Instance)
}
