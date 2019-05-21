package handler

import (
	"encoding/json"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUp(test *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(test, err)
	defer db.Close()
	model.DB = db

	err = model.Config(db)
	assert.NoError(test, err)

	e := echo.New()
	e.Validator = NewValidator()
	json, err := json.Marshal(J{
		"username": "username",
		"password": "password123",
	})
	assert.NoError(test, err)
	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(string(json)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(test, SignUp(c)) {
		assert.Equal(test, http.StatusCreated, rec.Code)
	}
}
