package component

import (
	"encoding/json"
	"testing"
)

type PublicProfileMPA struct {
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

type PublicProfileB struct {
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
		Alias string `json:"alias"`
		QR    string `json:"qrcode_url"`
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

func pP() *PublicProfileMPA {
	pro := `{
		"authorizer_info": {
			"nick_name": "微信SDK Demo Special",
			"head_img": "http://wx.qlogo.cn/mmopen/GPy",
			"service_type_info": {
				"id": 2
			},
			"verify_type_info": {
				"id": 0
			},
			"user_name": "gh_eb5e3a772040",
			"principal_name": "腾讯计算机系统有限公司",
			"business_info": {
				"open_store": 0,
				"open_scan": 0,
				"open_pay": 0,
				"open_card": 0,
				"open_shake": 0
			},
			"qrcode_url": "URL",
			"signature": "时间的水缓缓流去",
			"MiniProgramInfo": {
				"network": {
					"RequestDomain": ["https://www.qq.com", "https://www.qq.com"],
					"WsRequestDomain": ["wss://www.qq.com", "wss://www.qq.com"],
					"UploadDomain": ["https://www.qq.com", "https://www.qq.com"],
					"DownloadDomain": ["https://www.qq.com", "https://www.qq.com"]
				},
				"categories": [{
					"first": "资讯",
					"second": "文娱"
				}, {
					"first": "工具",
					"second": "天气"
				}],
				"visit_status": "0"
			}
		},
		"authorization_info": {
			"authorization_appid": "wxf8b4f85f3a794e77",
			"func_info": [{
					"funcscope_category": {
						"id": 17
					}
				},
				{
					"funcscope_category": {
						"id": 18
					}
				},
				{
					"funcscope_category": {
						"id": 19
					}
				}
			]
		}
	}`

	//var rs PublicProfileB
	var rs PublicProfileMPA
	if err := json.Unmarshal([]byte(pro), &rs); err != nil {
		return nil
	}

	return &rs
}

func TestProfile(t *testing.T) {
	rs := pP()

	minJSON, _ := json.Marshal(rs.AuthorizerInfo.MiniProgramInfo)

	t.Log(string(minJSON))
}
