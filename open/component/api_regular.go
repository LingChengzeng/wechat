package component

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lingchengzeng/wechat/tools"
)

type APIRegular interface {
	GetAccessToken(ticket string) (string, float64, error)
	GetPreAuthCode(accessToken string) (string, float64, error)
}

type apiRegular struct {
	wt *WechatThird
}

// 2.获取第三方平台component_access_token
// 第三方平台component_access_token是第三方平台的下文中接口的调用凭据,也叫做令牌
func (this *apiRegular) GetAccessToken(ticket string) (string, float64, error) {
	postData := struct {
		Component_appid         string `json:"component_appid"`
		Component_appsecret     string `json:"component_appsecret"`
		Component_verify_ticket string `json:"component_verify_ticket"`
	}{
		Component_appid:         this.wt.appId,
		Component_appsecret:     this.wt.appSecret,
		Component_verify_ticket: ticket,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return "", 0, errors.New("GetAccessToken 参数格式转换错误")
	}

	body, err := tools.PostString(apiComponentToken, string(jsonData))
	if err != nil {
		return "", 0, err
	}

	var msg map[string]interface{}
	if err = json.Unmarshal([]byte(body), &msg); err != nil {
		return "", 0, err
	}

	msgMap := tools.NewMap(msg)
	return msgMap.GetToString("component_access_token"), msgMap.GetToFloat64("expires_in"), nil
}

// 3.获取预授权码pre_auth_code,用于公众号oauth
// 该API用于获取预授权码。预授权码用于公众号或小程序授权时的第三方平台方安全验证。
func (this *apiRegular) GetPreAuthCode(accessToken string) (string, float64, error) {
	postData := struct {
		Component_appid string `json:"component_appid"`
	}{
		Component_appid: this.wt.appId,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return "", 0, errors.New("GetPreAuthCode 参数格式转换错误")
	}

	body, err := tools.PostString(fmt.Sprintf(apiCreatePreAuthCode, accessToken), string(jsonData))
	if err != nil {
		return "", 0, err
	}

	var msg map[string]interface{}
	if err = json.Unmarshal([]byte(body), &msg); err != nil {
		return "", 0, err
	}

	msgMap := tools.NewMap(msg)
	return msgMap.GetToString("pre_auth_code"), msgMap.GetToFloat64("expires_in"), nil
}
