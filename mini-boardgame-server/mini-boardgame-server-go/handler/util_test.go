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
	USERID    = 1
	USERNAME  = "username"
	USERPWD   = "user123456"
	ADMINID   = 2
	ADMINNAME = "adminname"
	ADMINPWD  = "admin123456"
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
	err = t.handler(c)

	// assert error
	if t.response.Error {
		assert.Error(t.test, err)
		if t.response.ErrorDetail != "" {
			assert.EqualError(t.test, err, t.response.ErrorDetail)
		}
		echoContext.HTTPErrorHandler(err, c)
	} else {
		assert.NoError(t.test, err)
	}

	// assert code
	assert.Equal(t.test, t.response.Code, rec.Code)

	// assert body
	if t.response.Body != nil {
		var js []byte
		js, err = json.Marshal(t.response.Body)
		assert.NoError(t.test, err)
		assert.JSONEq(t.test, rec.Body.String(), string(js))
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
