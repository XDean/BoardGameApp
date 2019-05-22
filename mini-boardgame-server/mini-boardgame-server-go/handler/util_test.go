package handler

import (
	"encoding/json"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

var echoContext *echo.Echo

func init() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	model.DB = db

	err = model.Config(db)
	if err != nil {
		panic(err)
	}

	echoContext = echo.New()
	echoContext.Validator = NewValidator()
}

type Request struct {
	Method string
	Path   string
	Body   interface{}
}

type Response struct {
	Code        int
	Error       bool
	ErrorDetail string
}

func testHttp(test *testing.T, request Request, response Response, fn echo.HandlerFunc) {
	request = defaultRequest(request)

	var body *strings.Reader
	var err error
	if request.Body != nil {
		var js []byte
		js, err = json.Marshal(request.Body)
		body = strings.NewReader(string(js))
		assert.NoError(test, err)
	}
	req := httptest.NewRequest(request.Method, request.Path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := echoContext.NewContext(req, rec)

	err = fn(c)
	if response.Error {
		assert.Error(test, err)
		if response.ErrorDetail != "" {
			assert.EqualError(test, err, response.ErrorDetail)
		}
		echoContext.HTTPErrorHandler(err, c)
	}
	assert.Equal(test, response.Code, rec.Code)
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
