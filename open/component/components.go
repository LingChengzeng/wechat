package component

import (
	"fmt"
	"net/url"
)

// 微信接口
const wechatComponentAPIHost = "https://api.weixin.qq.com/cgi-bin/component"

var (
	apiComponentToken      = wechatComponentAPIHost + "/api_component_token"
	apiCreatePreAuthCode   = wechatComponentAPIHost + "/api_create_preauthcode?component_access_token=%s"
	apiQueryAuth           = wechatComponentAPIHost + "/api_query_auth?component_access_token=%s"
	apiAuthorizerToken     = wechatComponentAPIHost + "/api_authorizer_token?component_access_token=%s"
	apiGetAuthorizerInfo   = wechatComponentAPIHost + "/api_get_authorizer_info?component_access_token=%s"
	apiGetAuthorizerOption = wechatComponentAPIHost + "/api_get_authorizer_option?component_access_token=%s"
	apiSetAuthorizerOption = wechatComponentAPIHost + "/api_set_authorizer_option?component_access_token=%s"

	oauthUrl = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s"
)

type WechatComponent interface {
	GetRegularApi() APIRegular
	GetNormalApi() APINormal
	OAuthUrl(redirectUrl, preAuthCode string) string
	GetCipher() (IOCipher, error)
}

func New(appId, appSecret, cryptoKey, token string) WechatComponent {
	return &WechatThird{
		appId:     appId,
		appSecret: appSecret,
		cryptoKey: cryptoKey,
		token:     token,
	}
}

// 微信第三方公众号平台 实现 ServeHTTP interface
type WechatThird struct {
	appId     string // 第三方应用id
	appSecret string // 第三方应用secret
	cryptoKey string // 公众号消息加解密Key
	token     string // 公众号消息校验Token
}

func (this *WechatThird) GetRegularApi() APIRegular {
	return &apiRegular{
		wt: this,
	}
}

func (this *WechatThird) GetNormalApi() APINormal {
	return &apiNormal{
		wt: this,
	}
}

func (this *WechatThird) GetCipher() (IOCipher, error) {
	return NewCipher(this.token, this.cryptoKey, this.appId)
}

func (this *WechatThird) OAuthUrl(redirectUrl, preAuthCode string) string {
	url, _ := UrlEncoded(redirectUrl)
	return fmt.Sprintf(oauthUrl, this.appId, preAuthCode, url)
}

func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
