package component

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lingchengzeng/wechat/tools"
)

type ApiError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type APINormal interface {
	GetPublicInfo(accessToken, authCode string) (*PublicInfo, error)
	GetAuthAccessToken(accessToken, appId, refreshToken string) (*PublicToken, error)
	GetAuthProfile(accessToken, appId string) (*PublicProfile, error)
	GetAuthOption(accessToken, appId, option string) (*PublicOption, error)
	SetAuthOption(accessToken, appId, optionName, optionValue string) error
}

type apiNormal struct {
	wt *WechatThird
}

type PublicInfo struct {
	AuthorizationInfo struct {
		AppId        string  `json:"authorizer_appid"`
		AccessToken  string  `json:"authorizer_access_token"`
		ExpiresIn    float64 `json:"expires_in"`
		RefreshToken string  `json:"authorizer_refresh_token"`
		FuncInfo     []struct {
			Funcscope struct {
				Id int64 `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

// 4. 使用授权码换取公众号或小程序的接口调用凭据和授权信息
// 该API用于使用授权码换取授权公众号或小程序的授权信息，并换取authorizer_access_token
// 和authorizer_refresh_token。 授权码的获取，需要在用户在第三方平台授权页中完成授权流程后，
// 在回调URI中通过URL参数提供给第三方平台方。请注意，由于现在公众号或小程序可以自定义选择部分
// 权限授权给第三方平台，因此第三方平台开发者需要通过该接口来获取公众号或小程序具体授权了哪些权
// 限，而不是简单地认为自己声明的权限就是公众号或小程序授权的权限。
func (this *apiNormal) GetPublicInfo(accessToken, authCode string) (*PublicInfo, error) {
	postData := struct {
		Component_appid    string `json:"component_appid"`
		Authorization_code string `json:"authorization_code"`
	}{
		Component_appid:    this.wt.appId,
		Authorization_code: authCode,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, errors.New("GetPublicInfo 参数格式转换错误")
	}

	body, err := tools.PostString(fmt.Sprintf(apiQueryAuth, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}

	var rs PublicInfo
	if err = json.Unmarshal([]byte(body), &rs); err != nil {
		return nil, err
	}

	return &rs, nil
}

// authorizer access token and refresh token
type PublicToken struct {
	AccessToken  string  `json:"authorizer_access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	RefreshToken string  `json:"authorizer_refresh_token"`
}

// 5.获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
// 该API用于在授权方令牌（authorizer_access_token）失效时，
// 可用刷新令牌（authorizer_refresh_token）获取新的令牌。
func (this *apiNormal) GetAuthAccessToken(accessToken, appId, refreshToken string) (*PublicToken, error) {
	postData := struct {
		Component_appid          string `json:"component_appid"`
		Authorizer_appid         string `json:"authorizer_appid"`
		Authorizer_refresh_token string `json:"authorizer_refresh_token"`
	}{
		Component_appid:          this.wt.appId,
		Authorizer_appid:         appId,
		Authorizer_refresh_token: refreshToken,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, errors.New("GetAuthAccessToken 参数格式转换错误")
	}

	body, err := tools.PostString(fmt.Sprintf(apiAuthorizerToken, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}

	var rs PublicToken
	if err = json.Unmarshal([]byte(body), &rs); err != nil {
		return nil, err
	}

	return &rs, nil
}

// 小程序返回的数据结构
type PublicProfile struct {
	AuthorizerInfo struct {
		NickName        string `json:"nick_name"`
		HeadImg         string `json:"head_img"`
		ServiceTypeInfo struct {
			Id int64 `json:"id"`
		}
		VerifyTypeInfo struct {
			Id int64 `json:"id"`
		}
		UserName      string `json:"user_name"`
		PrincipalName string `json:"principal_name"`
		BusinessInfo  struct {
			OpenStore string `json:"open_store"`
			OpenScan  string `json:"open_scan"`
			OpenPay   string `json:"open_pay"`
			OpenCard  string `json:"open_card"`
			OpenShake string `json:"open_shake"`
		}
		Alias           string `json:"alias"`
		QR              string `json:"qrcode_url"`
		Signature       string `json: "signature"`
		MiniProgramInfo struct {
			Network struct {
				RequestDomain   []string `json: "RequestDomain"`
				WsRequestDomain []string `json: "WsRequestDomain"`
				UploadDomain    []string `json: "UploadDomain"`
				DownloadDomain  []string `json: "DownloadDomain"`
			} `json:"network"`
			Categories []struct {
				First  string `json: "first"`
				Second string `json: "second"`
			} `json: "categories"`
			VisitStatus string `json: "visit_status"`
		}
	} `json:"authorizer_info"`
	AuthorizationInfo struct {
		AppId    string `json:"authorization_appid"`
		FuncInfo []struct {
			Funcscope struct {
				Id int64 `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

// 6.获取授权方的帐号基本信息
// 该API用于获取授权方的基本信息，包括头像、昵称、帐号类型、认证类型、微信号、原始ID和二维码图
// 片URL。
// 需要特别记录授权方的帐号类型，在消息及事件推送时，对于不具备客服接口的公众号，需要在5秒内立
// 即响应；而若有客服接口，则可以选择暂时不响应，而选择后续通过客服接口来发送消息触达粉丝。
// appId为平台APPID
func (this *apiNormal) GetAuthProfile(accessToken, authorizer_appid string) (*PublicProfile, error) {
	postData := struct {
		Component_appid  string `json:"component_appid"`
		Authorizer_appid string `json:"authorizer_appid"`
	}{
		Component_appid:  this.wt.appId,    // 平台APPID
		Authorizer_appid: authorizer_appid, // 授权方公众号APPID
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, errors.New("GetAuthProfile 参数格式转换错误")
	}

	body, err := tools.PostString(fmt.Sprintf(apiGetAuthorizerInfo, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}

	// TODO 解析返回数据
	var rs PublicProfile
	if err = json.Unmarshal(body, &rs); err != nil {
		return nil, err
	}

	return &rs, nil
}

type PublicOption struct {
	AppId       string `json:"authorizer_appid"`
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
}

// 7.获取授权方的选项设置信息
// 该API用于获取授权方的公众号或小程序的选项设置信息，如：地理位置上报，语音识别开关，多客服开
// 关。注意，获取各项选项设置信息，需要有授权方的授权，详见权限集说明。
func (this *apiNormal) GetAuthOption(accessToken, appId, option string) (*PublicOption, error) {
	postData := struct {
		Component_appid  string `json:"component_appid"`
		Authorizer_appid string `json:"authorizer_appid"`
		Option_name      string `json:"option_name"`
	}{
		Component_appid:  this.wt.appId,
		Authorizer_appid: appId,
		Option_name:      option,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, errors.New("GetAuthOption 参数格式转换错误")
	}

	body, err := tools.PostString(fmt.Sprintf(apiGetAuthorizerOption, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}

	var rs PublicOption
	if err = json.Unmarshal(body, &rs); err != nil {
		return nil, err
	}

	return &rs, nil
}

// 8.设置授权方的选项信息
// 该API用于设置授权方的公众号或小程序的选项信息，如：地理位置上报，语音识别开关，多客服开关。
// 注意，设置各项选项设置信息，需要有授权方的授权，详见权限集说明。
func (this *apiNormal) SetAuthOption(accessToken, appId, optionName, optionValue string) error {
	postData := struct {
		Component_appid  string `json:"component_appid"`
		Authorizer_appid string `json:"authorizer_appid"`
		Option_name      string `json:"option_name"`
		Option_value     string `json:"option_value"`
	}{
		Component_appid:  this.wt.appId,
		Authorizer_appid: appId,
		Option_name:      optionName,
		Option_value:     optionValue,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return errors.New("SetAuthOption 参数格式转换错误")
	}

	body, err := tools.PostString(fmt.Sprintf(apiSetAuthorizerOption, accessToken), string(jsonData))
	if err != nil {
		return err
	}

	var rs ApiError
	if err = json.Unmarshal([]byte(body), &rs); err != nil {
		return err
	}

	// TODO 需要测试和调整
	return json.Unmarshal(body, &rs)
}
