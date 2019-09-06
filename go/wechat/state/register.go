package state

import (
	"github.com/xdean/miniboardgame/go/games/guobiao"
	"github.com/xdean/miniboardgame/go/games/ocr"
)

func init() {
	Register(guobiao.State)
	Register(ocr.State)
}
