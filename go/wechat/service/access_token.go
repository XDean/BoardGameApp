package service

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xdean/miniboardgame/go/wechat/config"
	"github.com/xdean/miniboardgame/go/wechat/model"
	"io/ioutil"
	"net/http"
	"time"
)

func StartAccessTokenTask() {
	go accessTokenTask()
}

var getAccessTokenStream chan bool = make(chan bool)
var accessTokenStream chan string = make(chan string)

func GetAccessToken() string {
	getAccessTokenStream <- true
	return <-accessTokenStream
}

func accessTokenTask() {
	type Response struct {
		model.Error
		Token   string `json:"access_token"`
		Expires int    `json:"expires_in"`
	}
	var accessToken string
	var expire <-chan time.Time = time.After(0)
	for {
		select {
		case <-expire:
			break
		case <-getAccessTokenStream:
			accessTokenStream <- accessToken
			continue
		}
		query := fmt.Sprintf("https://%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			config.Instance.Wechat.Url, config.Instance.Wechat.AppId, config.Instance.Wechat.AppSecret)
		logrus.Debugln("To query access token: ", query)
		resp, err := http.Get(query)
		if err != nil {
			logrus.WithError(err).Error("Fail to query access token")
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.WithError(err).Error("Fail to read access token body")
			} else {
				result := new(Response)
				err := json.Unmarshal(body, &result)
				if err != nil {
					logrus.WithError(err).Error("Fail to parse access token response")
				} else {
					if result.ErrorCode == 0 {
						logrus.Debug("Got access token: ", result.Token)
						accessToken = result.Token
						expire = time.After(time.Second * time.Duration(result.Expires))
					} else {
						logrus.WithField("Code", result.ErrorCode).Error("Fail to get access token")
					}
				}
			}
		}
	}
}
