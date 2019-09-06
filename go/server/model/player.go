package model

type PlayerState uint

const (
	OUT_OF_GAME PlayerState = 0
	NOT_READY   PlayerState = 1
	READY       PlayerState = 2
	HOST        PlayerState = 3
)

type Player struct {
	UserID uint        `gorm:"primary_key"`
	RoomID uint        `gorm:"default:0"`
	Room   *Room       `gorm:"-"` //gorm can't handle this, need put manually
	State  PlayerState `gorm:"default:0"`
	Seat   uint        `gorm:"default:0"`
}

func (s PlayerState) String() string {
	switch s {
	case OUT_OF_GAME:
		return "OUT_OF_GAME"
	case NOT_READY:
		return "NOT_READY"
	case READY:
		return "READY"
	case HOST:
		return "HOST"
	default:
		return "Unknown"
	}
}

func (p *Player) normalize() {
	// do nothing
}
