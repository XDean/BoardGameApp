package handler

import (
	"encoding/json"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	USERNAME  = "username"
	USERPWD   = "user123456"
	ADMINNAME = "adminname"
	ADMINPWD  = "admin123456"
)

var (
	USER = &model.User{
		Username: USERNAME,
		Password: USERPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_USER,
			},
		},
	}
	ADMIN = &model.User{
		Username: ADMINNAME,
		Password: ADMINPWD,
		Roles: []model.Role{
			{
				Name: _const.ROLE_ADMIN,
			},
		},
	}
)

type (
	Setup func(echo.Context)

	TestHttp struct {
		test    *testing.T
		handler echo.HandlerFunc
		req     Request
		res     Response
		setups  []Setup
	}

	Request struct {
		Method string
		Path   string
		Body   interface{}
	}

	Response struct {
		Code        int
		Error       bool
		ErrorDetail string
	}
)

func WithUser(t *testing.T, user *model.User) Setup {
	return func(c echo.Context) {
		err := user.CreateAccount(GetDB(c))
		assert.NoError(t, err)
	}
}

func WithLogin(t *testing.T, user *model.User) Setup {
	return func(c echo.Context) {
		c.Set(_const.USER, user)
	}
}

func (t TestHttp) Run() {
	t.req = defaultRequest(t.req)
	t.res = defaultResponse(t.res)

	tx := dbContext.Begin()
	defer tx.Rollback()

	var body *strings.Reader
	var err error
	if t.req.Body != nil {
		var js []byte
		js, err = json.Marshal(t.req.Body)
		body = strings.NewReader(string(js))
		assert.NoError(t.test, err)
	}
	req := httptest.NewRequest(t.req.Method, t.req.Path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := echoContext.NewContext(req, rec)
	c.Set(_const.DATABASE, tx)
	if t.setups != nil {
		for _, setup := range t.setups {
			setup(c)
		}
	}
	err = t.handler(c)
	if t.res.Error {
		assert.Error(t.test, err)
		if t.res.ErrorDetail != "" {
			assert.EqualError(t.test, err, t.res.ErrorDetail)
		}
		echoContext.HTTPErrorHandler(err, c)
	} else {
		assert.NoError(t.test, err)
	}
	assert.Equal(t.test, t.res.Code, rec.Code)
}

func defaultRequest(request Request) Request {
	if request.Method == "" {
		request.Method = echo.GET
	}
	if request.Path == "" {
		request.Path = "/mock-path"
	}
	return request
}

func defaultResponse(response Response) Response {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}
	return response
}
