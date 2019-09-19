package model

import (
	"fmt"
	"github.com/xdean/miniboardgame/go/server/model/space"
)

func (r *Room) EventHostId() string {
	return fmt.Sprintf("ROOM-%d", r.ID)
}

func (r *Room) SendEvent(e space.Message) {
	space.SendEvent(r, e)
}

func (r *Room) Listen() space.Subscription {
	return space.Listen(r)
}

func (r *Room) Do(f func()) {
	space.Do(r, f)
}

func (r *Room) DoAndWait(f func()) {
	space.DoAndWait(r, f)
}
