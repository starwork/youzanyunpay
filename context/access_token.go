package context

import (
	"encoding/json"
	"fmt"
	"github.com/yuyan2077/youzanyunpay/util"
	"strconv"
	"sync"
)

const (
	//AccessTokenURL 获取access_token的接口
	AccessTokenURL = "https://open.youzan.com/oauth/token"
)

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

//SetAccessTokenLock 设置读写锁（一个appID一个读写锁）
func (ctx *Context) SetAccessTokenLock(l *sync.RWMutex) {
	ctx.accessTokenLock = l
}

//GetAccessToken 获取access_token
func (ctx *Context) GetAccessToken() (accessToken string, err error) {
	ctx.accessTokenLock.Lock()
	defer ctx.accessTokenLock.Unlock()

	if ctx.Token.TokenTimeOut > util.GetCurrTs() {
		accessToken = ctx.Token.Token
		return
	}

	//从有赞云服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = ctx.GetAccessTokenFromServer()
	if err != nil {
		return
	}

	accessToken = resAccessToken.AccessToken
	return
}

//GetAccessTokenFromServer 强制从有赞云服务器获取token
func (ctx *Context) GetAccessTokenFromServer() (resAccessToken ResAccessToken, err error) {
	httpUrl := AccessTokenURL
	var params = map[string]string{
		"client_id":     ctx.AppID,
		"client_secret": ctx.AppSecret,
		"grant_type":    "silent",
		"kdt_id":        strconv.Itoa(ctx.KdtID),
	}
	resp := util.SendRequest(httpUrl, "POST", params, map[string]string{})

	json.Unmarshal(resp, &resAccessToken)
	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	ctx.Token.TokenTimeOut = expires + util.GetCurrTs()
	ctx.Token.Token = resAccessToken.AccessToken
	return
}
