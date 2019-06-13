package handler

import (
	"encoding/json"
	"errors"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/XDean/MiniBoardgame/web/handler/openid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xdean/goex/xgo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	USERID    = 1
	USERNAME  = "username"
	USERPWD   = "user123456"
	ADMINID   = 2
	ADMINNAME = "adminname"
	ADMINPWD  = "admin123456"
	ROOMID    = 1
)

var (
	USER = &model.User{
		ID:       USERID,
		Username: USERNAME,
		Password: USERPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
	ADMIN = &model.User{
		ID:       ADMINID,
		Username: ADMINNAME,
		Password: ADMINPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_ADMIN,
			},
		},
	}
	USER_PROFILE = &model.Profile{
		UserID:    USERID,
		Sex:       model.Male,
		Nickname:  "usernick",
		AvatarURL: "userurl",
	}
	ADMIN_PROFILE = &model.Profile{
		UserID:    USERID,
		Sex:       model.Male,
		Nickname:  "adminname",
		AvatarURL: "adminurl",
	}
	ROOM = &model.Room{
		ID:          ROOMID,
		GameName:    "game name",
		RoomName:    "room name",
		PlayerCount: 3,
	}
)

type (
	Setup  func(echo.Context)
	Params map[string]string

	TestHttp struct {
		test     *testing.T
		handler  echo.HandlerFunc
		request  Request
		response Response
		setups   []Setup
	}

	Request struct {
		Method string
		Path   string
		Params Params
		Body   interface{}
	}

	Response struct {
		Code        int
		Body        interface{}
		Error       bool
		ErrorDetail string
		Extra       func(*gorm.DB, *httptest.ResponseRecorder)
	}
)

func WithUser(t *testing.T, user *model.User) Setup {
	return func(c echo.Context) {
		err := user.CreateAccount(GetDB(c))
		assert.NoError(t, err)
	}
}

func WithProfile(t *testing.T, profile *model.Profile) Setup {
	return func(c echo.Context) {
		err := profile.Save(GetDB(c))
		assert.NoError(t, err)
	}
}

func WithLogin(t *testing.T, user *model.User) Setup {
	return func(c echo.Context) {
		c.Set(_const.USER, user)
	}
}

func WithRoom(t *testing.T, room *model.Room) Setup {
	return func(c echo.Context) {
		user, err := GetCurrentUser(c)
		assert.NoError(t, err)

		player := new(model.Player)
		err = player.GetByUserID(GetDB(c), user.ID)
		assert.NoError(t, err)

		assert.NoError(t, err)
		err = room.CreateByHost(GetDB(c), player)
		assert.NoError(t, err)
	}
}

func WithOpenid() Setup {
	return func(c echo.Context) {
		openid.Providers = map[string]openid.OpenIdProvider{
			"test": {
				Name: "test",
				Auth: func(token string) (string, error) {
					return token, nil
				},
			},
			"test-fail": {
				Name: "test",
				Auth: func(token string) (string, error) {
					return "", errors.New("openid fail")
				},
			},
		}
	}
}

func initVars() {
	USER = &model.User{
		ID:       USERID,
		Username: USERNAME,
		Password: USERPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
	ADMIN = &model.User{
		ID:       ADMINID,
		Username: ADMINNAME,
		Password: ADMINPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_ADMIN,
			},
		},
	}
	USER_PROFILE = &model.Profile{
		UserID:    USERID,
		Sex:       model.Male,
		Nickname:  "usernick",
		AvatarURL: "userurl",
	}
	ADMIN_PROFILE = &model.Profile{
		UserID:    USERID,
		Sex:       model.Male,
		Nickname:  "adminname",
		AvatarURL: "adminurl",
	}
	ROOM = &model.Room{
		ID:          ROOMID,
		GameName:    "game name",
		RoomName:    "room name",
		PlayerCount: 3,
	}
}

func (t TestHttp) Run() {
	initVars()
	// prepare request and response
	t.request = defaultRequest(t.request)
	t.response = defaultResponse(t.response)

	// prepare db
	tx := dbContext.Begin()
	defer tx.Rollback()

	// format body and create request object
	var body *strings.Reader
	var err error
	var req *http.Request
	if t.request.Body != nil {
		var js []byte
		js, err = json.Marshal(t.request.Body)
		body = strings.NewReader(string(js))
		assert.NoError(t.test, err)
		req = httptest.NewRequest(t.request.Method, t.request.Path, body)
	} else {
		req = httptest.NewRequest(t.request.Method, t.request.Path, nil)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// create response object and echo context
	rec := httptest.NewRecorder()
	c := echoContext.NewContext(req, rec)

	// setup echo context
	c.Set(_const.DATABASE, tx)
	if t.request.Params != nil {
		keys := make([]string, 0)
		values := make([]string, 0)
		for k, v := range t.request.Params {
			keys = append(keys, k)
			values = append(values, v)
		}
		c.SetParamNames(keys...)
		c.SetParamValues(values...)
	}
	if t.setups != nil {
		for _, setup := range t.setups {
			setup(c)
		}
	}

	// handle
	func() {
		defer func() {
			if r := recover(); r != nil {
				e, ok := r.(model.BreakError)
				if !ok {
					panic(r)
				}
				err = e.Actual
			}
		}()
		err = t.handler(c)
	}()
	echoContext.HTTPErrorHandler(err, c)

	// assert error
	if t.response.Error {
		assert.Error(t.test, err)
		if t.response.ErrorDetail != "" {
			assert.EqualError(t.test, err, t.response.ErrorDetail)
		}
	} else {
		assert.NoError(t.test, err)
	}

	// assert code
	assert.Equal(t.test, t.response.Code, rec.Code)

	// assert body
	if t.response.Body != nil {
		actualResponse := make(J)
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t.test, err)
		if ok, err := xgo.StructContain(actualResponse, t.response.Body); !ok {
			assert.Fail(t.test, "Body: "+err.Error())
		}
	}

	// extra
	if t.response.Extra != nil {
		t.response.Extra(tx, rec)
	}
}

func defaultRequest(request Request) Request {
	if request.Method == "" {
		request.Method = echo.GET
	}
	if request.Path == "" {
		request.Path = "/mock-path"
	}
	if request.Params == nil {
		request.Params = make(map[string]string)
	}
	return request
}

func defaultResponse(response Response) Response {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}
	return response
}
