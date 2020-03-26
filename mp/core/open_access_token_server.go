package core

import (
	"github.com/lingchengzeng/wechat/util"
	"net/http"
	"net/url"
)
//开放平台使用第三方token
type DefaultOpenAccessTokenServer struct {
	appId      string
	appSecret  string
	token      string
	httpClient *http.Client
}

func NewDefaultOpenAccessTokenServer(appId, appSecret,token string, httpClient *http.Client) (srv *DefaultOpenAccessTokenServer) {
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	srv = &DefaultOpenAccessTokenServer{
		appId:                    url.QueryEscape(appId),
		appSecret:                url.QueryEscape(appSecret),
		httpClient:               httpClient,
		token:                    token,
	}

	return
}

func (srv *DefaultOpenAccessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {}

func (srv *DefaultOpenAccessTokenServer) Token() (token string, err error) {
	return srv.token,nil
}

func (srv *DefaultOpenAccessTokenServer) RefreshToken(currentToken string) (token string, err error) {
	if len(currentToken) != 0 {
		return srv.token,nil
	}
	return "",nil
}
