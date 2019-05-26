package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const ROOMID uint = 1

func TestRoom_CreateByHost(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	player := new(Player)
	err := player.GetByUserID(db, USERID)
	assert.NoError(t, err)
	assert.Equal(t, Player{
		UserID:      USERID,
		RoomID:      0,
		Room:        nil,
		State:       OUT_OF_GAME,
		StateString: OUT_OF_GAME.String(),
		Seat:        0,
	}, *player)

	room := &Room{
		ID:          ROOMID,
		GameName:    "game",
		RoomName:    "room",
		PlayerCount: 3,
	}
	err = room.CreateByHost(db, NewUser())
	assert.NoError(t, err)
	assert.Equal(t, room.Players[0].Room, room)

	err = player.GetByUserID(db, USERID)
	assert.NoError(t, err)
	assert.Equal(t, player.Room.Players[0], player)
	assert.Equal(t, ROOMID, player.RoomID)
	assert.Equal(t, "game", player.Room.GameName)
	assert.Equal(t, "room", player.Room.RoomName)

	room = new(Room)
	err = room.FindByID(db, ROOMID)
	assert.NoError(t, err)
	assert.Equal(t, room.Players[0].Room, room)
	assert.Equal(t, ROOMID, room.ID)
	assert.Equal(t, "game", room.GameName)
	assert.Equal(t, "room", room.RoomName)

	rooms, err := FindRoomsByGame(db, "game", Unbound)
	assert.NoError(t, err)
	assert.Len(t, rooms, 1)
}
