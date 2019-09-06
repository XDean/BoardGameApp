package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewPlayer() *Player {
	return &Player{
		UserID: USERID,
	}
}

func TestPlayer_FindByUserID(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	player := NewPlayer()
	err := player.save(db)
	assert.NoError(t, err)
}
