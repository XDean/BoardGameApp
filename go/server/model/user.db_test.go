package model

import (
	"github.com/stretchr/testify/assert"
	_const "github.com/xdean/miniboardgame/go/server/const"
	"testing"
)

const (
	USERID   = 1
	USERNAME = "username"
	USERPWD  = "user123456"
)

func NewUser() *User {
	return &User{
		ID:       USERID,
		Username: USERNAME,
		Password: USERPWD,
		Roles: []Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
}

func TestUser_CreateAccount(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	user := NewUser()
	err := user.CreateAccount(db)
	assert.NoError(t, err)

	err = user.CreateAccount(db)
	assert.Error(t, err)
}

func TestUser_FindByID(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	user := new(User)
	err := user.FindByID(db, USERID)
	assert.Error(t, err)

	err = NewUser().save(db)
	assert.NoError(t, err)
	err = user.FindByID(db, USERID)
	assert.NoError(t, err)
}

func TestUser_FindByUsername(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	user := new(User)
	err := user.FindByUsername(db, USERNAME)
	assert.Error(t, err)

	err = NewUser().save(db)
	assert.NoError(t, err)
	err = user.FindByUsername(db, USERNAME)
	assert.NoError(t, err)
}

func TestUser_MatchPassword(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	user := NewUser()
	err := user.CreateAccount(db)
	assert.NoError(t, err)

	yes := user.MatchPassword(USERPWD)
	assert.True(t, yes)
}

func TestUser_ChangePassword(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	user := NewUser()
	err := user.CreateAccount(db)
	assert.NoError(t, err)

	err = user.ChangePassword(db, USERPWD, "newpwd")
	assert.NoError(t, err)

	yes := user.MatchPassword("newpwd")
	assert.True(t, yes)

	err = user.ChangePassword(db, "wrong", "wrong")
	assert.Error(t, err)
}

func TestUserExistById(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	yes, err := UserExistById(db, USERID)
	assert.NoError(t, err)
	assert.False(t, yes)

	user := NewUser()
	err = user.CreateAccount(db)
	assert.NoError(t, err)

	yes, err = UserExistById(db, USERID)
	assert.NoError(t, err)
	assert.True(t, yes)
}

func TestUserExistByUsername(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	yes, err := UserExistByUsername(db, USERNAME)
	assert.NoError(t, err)
	assert.False(t, yes)

	user := NewUser()
	err = user.CreateAccount(db)
	assert.NoError(t, err)

	yes, err = UserExistByUsername(db, USERNAME)
	assert.NoError(t, err)
	assert.True(t, yes)
}
