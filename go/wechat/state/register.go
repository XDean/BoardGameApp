package state

import "github.com/xdean/miniboardgame/go/games/guobiao"

func init() {
	Register(guobiao.StateInstance)
}
