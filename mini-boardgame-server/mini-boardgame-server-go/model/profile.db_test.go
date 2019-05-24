package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewProfile() *Profile {
	return &Profile{
		UserID:    USERID,
		Nickname:  "test-nickname",
		Male:      true,
		AvatarURL: "test-url",
	}
}

func TestProfile_FindByUserID(t *testing.T) {
	db := testDB.Begin()
	defer db.Rollback()

	profile := new(Profile)
	err := profile.FindByUserID(db, USERID)
	assert.Error(t, err)

	profile = NewProfile()
	err = profile.Save(db)
	assert.NoError(t, err)

	profile = new(Profile)
	err = profile.FindByUserID(db, USERID)
	assert.NoError(t, err)
}
