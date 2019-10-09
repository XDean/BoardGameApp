package handler

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/goex/xgo"
	_const "github.com/xdean/miniboardgame/go/server/const"
	"github.com/xdean/miniboardgame/go/server/game"
	"github.com/xdean/miniboardgame/go/server/model"
	"github.com/xdean/miniboardgame/go/server/web/handler/openid"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	USERID   = 1
	USERNAME = "username"
	USERPWD  = "user123456"

	USERID2   = 2
	USERNAME2 = "username2"
	USERPWD2  = "user1234562"

	USERID3   = 3
	USERNAME3 = "username3"
	USERPWD3  = "user1234563"

	ADMINID   = 4
	ADMINNAME = "adminname"
	ADMINPWD  = "admin123456"
	ROOMID    = 1

	GAME_ID   = "test-game"
	GAME_NAME = "Test Game"
)

var (
	USER, USER2, USER3, ADMIN   = new(model.User), new(model.User), new(model.User), new(model.User)
	USER_PROFILE, ADMIN_PROFILE = new(model.Profile), new(model.Profile)
	ROOM                        = new(model.Room)

	GAME = &game.Game{
		Id:      GAME_ID,
		Name:    GAME_NAME,
		Player:  game.Range{Min: 2, Max: 3},
		Options: []game.Option{},
		NewEvent: func() game.Event {
			e := game.BaseEvent{}
			e.ResponseStream = make(chan game.Response, 5)
			return e
		},
		OnEvent: func(event game.Event) game.Response {
			switch event.(type) {
			case game.NewGameEvent:
				return event.GetUser().ID
			}
			return event
		},
	}
)

func init() {
	initVars()
	game.Register(GAME)
}

type (
	DBSetup      func(db *gorm.DB)
	ContextSetup func(echo.Context)
	Params       map[string]string

	Setup interface {
		mark()
	}

	TestHttp struct {
		test     *testing.T
		handler  echo.HandlerFunc
		request  Request
		response Response
		setups   []Setup
	}

	TestHttpSeries struct {
		test     *testing.T
		setups   []Setup
		children []TestHttp
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

func (DBSetup) mark() {
}

func (ContextSetup) mark() {
}

func WithUser(t *testing.T, user *model.User) DBSetup {
	return func(db *gorm.DB) { assert.NoError(t, user.CreateAccount(db)) }
}

func WithProfile(t *testing.T, profile *model.Profile) DBSetup {
	return func(db *gorm.DB) { assert.NoError(t, profile.Save(db)) }
}

func WithCreateRoom(t *testing.T, room *model.Room, host *model.User) DBSetup {
	return func(db *gorm.DB) {
		player := new(model.Player)
		err := player.GetByUserID(db, host.ID)
		assert.NoError(t, err)
		err = room.CreateByHost(db, player)
		assert.NoError(t, err)
	}
}

func WithLogin(t *testing.T, user *model.User) ContextSetup {
	return func(c echo.Context) { c.Set(_const.USER, user) }
}

func WithInRoom(t *testing.T, room *model.Room) ContextSetup {
	return func(c echo.Context) { c.Set(_const.ROOM, room) }
}

func WithOpenid() ContextSetup {
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
	*USER = model.User{
		ID:       USERID,
		Username: USERNAME,
		Password: USERPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
	*USER2 = model.User{
		ID:       USERID2,
		Username: USERNAME2,
		Password: USERPWD2,
		Roles: []model.Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
	*USER3 = model.User{
		ID:       USERID3,
		Username: USERNAME3,
		Password: USERPWD3,
		Roles: []model.Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
	*ADMIN = model.User{
		ID:       ADMINID,
		Username: ADMINNAME,
		Password: ADMINPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_ADMIN,
			},
		},
	}
	*USER_PROFILE = model.Profile{
		UserID:    USERID,
		Sex:       model.Male,
		Nickname:  "usernick",
		AvatarURL: "userurl",
	}
	*ADMIN_PROFILE = model.Profile{
		UserID:    USERID,
		Sex:       model.Male,
		Nickname:  "adminname",
		AvatarURL: "adminurl",
	}
	*ROOM = model.Room{
		ID:          ROOMID,
		GameId:      GAME_ID,
		RoomName:    "room name",
		PlayerCount: 3,
	}
}

func (t TestHttp) Run() {
	initVars()

	// prepare db
	defer resetDB()()

	setupDB(dbContext, t.setups)

	t.doRun(dbContext)
}

func (s TestHttpSeries) Run() {
	initVars()

	// prepare db
	defer resetDB()()

	setupDB(dbContext, s.setups)

	for _, t := range s.children {
		t.test = s.test
		t.setups = append(s.setups, t.setups...)
		t.doRun(dbContext)
	}
}

func (t TestHttp) doRun(tx *gorm.DB) {
	// prepare request and response
	t.request = defaultRequest(t.request)
	t.response = defaultResponse(t.response)
	// create request and response object and echo context
	req := t.genRequest()
	rec := httptest.NewRecorder()
	c := echoContext.NewContext(req, rec)
	// setup echo context
	c.Set(_const.DATABASE, tx)
	t.setupRequestParam(c)
	t.setupContext(c)
	// handle
	err := t.doHttp(c)
	echoContext.HTTPErrorHandler(err, c)
	// assert
	t.doAssert(err, rec, tx)
}

func (t TestHttp) doAssert(err error, rec *httptest.ResponseRecorder, tx *gorm.DB) {
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
	expectBody := t.response.Body
	if expectBody != nil {
		actualResponse := make(xecho.J)
		err := json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.NoError(t.test, err)

		expectResponse := make(xecho.J)
		expectJson, err := json.Marshal(expectBody)
		assert.NoError(t.test, err)
		err = json.Unmarshal(expectJson, &expectResponse)
		assert.NoError(t.test, err)

		if ok, err := xgo.StructContain(actualResponse, expectResponse); !ok {
			if err == nil {
				assert.Fail(t.test, "Body not as expected: ", t.response.Body)
			} else {
				assert.Fail(t.test, "Body: "+err.Error())
			}
		}
	}
	// extra
	if t.response.Extra != nil {
		t.response.Extra(tx, rec)
	}
}

func (t TestHttp) doHttp(c echo.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			e, ok := r.(xecho.BreakError)
			if !ok {
				panic(r)
			}
			err = e.Actual
		}
	}()
	return t.handler(c)
}

func (t TestHttp) setupContext(c echo.Context) {
	for _, setup := range t.setups {
		if t, ok := setup.(ContextSetup); ok {
			t(c)
		}
	}
}

func setupDB(db *gorm.DB, setups []Setup) {
	for _, setup := range setups {
		if t, ok := setup.(DBSetup); ok {
			t(db)
		}
	}
}

func (t TestHttp) setupRequestParam(c echo.Context) {
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
}

func (t TestHttp) genRequest() (req *http.Request) {
	if t.request.Body != nil {
		var js []byte
		js, err := json.Marshal(t.request.Body)
		body := strings.NewReader(string(js))
		assert.NoError(t.test, err)
		req = httptest.NewRequest(t.request.Method, t.request.Path, body)
	} else {
		req = httptest.NewRequest(t.request.Method, t.request.Path, nil)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
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
