package context

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/yuyan2077/youzanyunpay/util"
	"strconv"
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

	accessTokenCacheKey := fmt.Sprintf("access_token_%s", ctx.AppID)
	val := ctx.Cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
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
		"grant_type":    ctx.GrantType,
		"kdt_id":        strconv.Itoa(ctx.KdtID),
	}
	resp := util.SendRequest(httpUrl, "POST", params, map[string]string{})

	json.Unmarshal(resp, &resAccessToken)
	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	accessTokenCacheKey := fmt.Sprintf("access_token_%s", ctx.AppID)
	expires := resAccessToken.ExpiresIn - 1500
	err = ctx.Cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	return
}
