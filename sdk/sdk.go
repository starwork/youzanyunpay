package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/yuyan2077/youzanyunpay/context"
	"github.com/yuyan2077/youzanyunpay/util"
	"strings"
)

const youzanyunAPIURL = "https://open.youzan.com/api"

//SDK
type SDK struct {
	*context.Context
}

type InvokeResp struct {
	Response      []byte `json:"response"`
	ErrorResponse []byte `json:"error_response"`
}

type InvokeResp1 struct {
	Response      CQrcodeResp `json:"response"`
	ErrorResponse []byte      `json:"error_response"`
}

type CQrcodeResp struct {
	QrUrl  string `json:"qr_url"`
	QrCode string `json:"qr_code"`
	QrID   int64  `json:"qr_id"`
}

func (sdk *SDK) Invoke(apiName string, version string, method string, params map[string]string, files map[string]string) (result []byte, err error) {
	var httpUrl = youzanyunAPIURL
	var apiNameList = strings.Split(apiName, ".")
	var serviceList = apiNameList[0 : len(apiNameList)-1]
	var service = strings.Join(serviceList, ".")
	var actionList = apiNameList[len(apiNameList)-1:]
	var action = strings.Join(actionList, ".")

	var paramMap = map[string]string{}
	httpUrl += "/oauthentry"
	paramMap["access_token"], err = sdk.GetAccessToken()
	if err != nil {
		return
	}
	for k, v := range params {
		paramMap[k] = v
	}
	httpUrl += "/"
	httpUrl += service
	httpUrl += "/"
	httpUrl += version
	httpUrl += "/"
	httpUrl += action
	resp := util.SendRequest(httpUrl, method, paramMap, files)
	var invokeResp1 InvokeResp1
	json.Unmarshal(resp, &invokeResp1)
	fmt.Println(invokeResp1)
	var invokeResp InvokeResp
	json.Unmarshal(resp, &invokeResp)
	fmt.Println(invokeResp)

	if invokeResp.ErrorResponse != nil {
		result = invokeResp.ErrorResponse
	} else {
		result = invokeResp.Response
	}
	return
}

//NewSDK 实例化SDK
func NewSDK(context *context.Context) *SDK {
	sdk := new(SDK)
	sdk.Context = context
	return sdk
}
