package model

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestUser_SetRoles(t *testing.T) {
	user := NewUser()
	roles := []string{"role1", "role2"}
	user.SetRoles(roles)
	result := user.GetRoleStrings()
	assert.Equal(t, roles, result)
}
