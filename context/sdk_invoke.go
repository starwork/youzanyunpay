package context

import (
	"github.com/yuyan2077/youzanyunpay/util"
	"strings"
)

const youzanyunAPIURL = "https://open.youzan.com/api"

func (ctx *Context) Invoke(apiName string, version string, method string, params map[string]string, files map[string]string) (resp []byte, err error) {
	var httpUrl = youzanyunAPIURL
	var apiNameList = strings.Split(apiName, ".")
	var serviceList = apiNameList[0 : len(apiNameList)-1]
	var service = strings.Join(serviceList, ".")
	var actionList = apiNameList[len(apiNameList)-1:]
	var action = strings.Join(actionList, ".")

	var paramMap = map[string]string{}
	httpUrl += "/oauthentry"
	paramMap["access_token"], err = ctx.GetAccessToken()
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
	resp = util.SendRequest(httpUrl, method, paramMap, files)
	return
}
