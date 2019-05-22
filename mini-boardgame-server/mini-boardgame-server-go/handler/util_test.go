package handler

import (
	"encoding/json"
	"fmt"
	"github.com/XDean/MiniBoardgame/config"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var echoContext *echo.Echo
var dbContext *gorm.DB

func TestMain(m *testing.M) {
	config.Global.Debug = true
	tmp, err := ioutil.TempFile("", "mini-bg-*.db")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmp.Name())
	fmt.Println("Temp database file:", tmp.Name())
	db, err := gorm.Open("sqlite3", tmp.Name())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = model.ConfigDB(db)
	if err != nil {
		panic(err)
	}

	dbContext = db
	echoContext = echo.New()
	echoContext.Validator = NewValidator()

	result := m.Run()

	os.Exit(result)
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

func testHttp(test *testing.T, fn echo.HandlerFunc, request Request, response Response) {
	tx := dbContext.Begin()
	defer tx.Rollback()

	request = defaultRequest(request)
	response = defaultResponse(response)

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
	c.Set(_const.DATABASE, tx)

	err = fn(c)
	if response.Error {
		assert.Error(test, err)
		if response.ErrorDetail != "" {
			assert.EqualError(test, err, response.ErrorDetail)
		}
		echoContext.HTTPErrorHandler(err, c)
	} else {
		assert.NoError(test, err)
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

func defaultResponse(response Response) Response {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}
	return response
}
