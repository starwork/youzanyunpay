package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/yuyan2077/youzanyunpay/context"
	"github.com/yuyan2077/youzanyunpay/util"
	"strings"
)

const YouzanyunAPIURL = "https://open.youzan.com/api"

//SDK
type SDK struct {
	*context.Context
}

type ResultInterface struct {
	Response      interface{} `json:"response"`
	ResponseError interface{} `json:"response_error"`
}

func (sdk *SDK) Invoke(apiName string, version string, method string, params map[string]string, files map[string]string) (responseMap map[string]interface{}, err error) {
	var httpUrl = YouzanyunAPIURL
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
		fmt.Println("k", k)
		fmt.Println("v", v)
		paramMap[k] = v
	}
	httpUrl += "/"
	httpUrl += service
	httpUrl += "/"
	httpUrl += version
	httpUrl += "/"
	httpUrl += action
	resp := util.SendRequest(httpUrl, method, paramMap, files)

	var resultInterface ResultInterface
	json.Unmarshal(resp, &resultInterface)
	responseErrorMap := resultInterface.ResponseError.(map[string]interface{})

	var commonError util.CommonError
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(responseErrorMap, &commonError); err != nil {
		return
	}

	if commonError.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", resultInterface.ResponseError.(util.CommonError).ErrCode, resultInterface.ResponseError.(util.CommonError).ErrMsg)
		return
	}

	responseMap = resultInterface.Response.(map[string]interface{})
	return
}

//NewSDK 实例化SDK
func NewSDK(context *context.Context) *SDK {
	sdk := new(SDK)
	sdk.Context = context
	return sdk
}
