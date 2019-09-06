package main

import (
	"github.com/xdean/miniboardgame/go/games/guobiao"
	"github.com/xdean/miniboardgame/go/games/ocr"
	"github.com/xdean/miniboardgame/go/wechat/state"
)

func init() {
	state.Register(guobiao.State)
	state.Register(ocr.State)
}
