package youzanyunpay

import (
	"github.com/yuyan2077/youzanyunpay/cache"
	"github.com/yuyan2077/youzanyunpay/context"
	"github.com/yuyan2077/youzanyunpay/sdk"
	"sync"
)

var YZYPay *Youzanyunpay

type Youzanyunpay struct {
	Context *context.Context
}

type Config struct {
	AppID     string
	AppSecret string
	Cache     cache.Cache
}

//func NewYouzanyunpay(cfg *Config) *Youzanyunpay {
//	ctx := new(context.Context)
//	copyConfigToContext(cfg, ctx)
//	return &Youzanyunpay{ctx}
//}

func NewYouzanyunpay(cfg *Config) {
	ctx := new(context.Context)
	copyConfigToContext(cfg, ctx)
	YZYPay = &Youzanyunpay{ctx}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

//GetAccessToken 获取access_token
func (pay *Youzanyunpay) GetAccessToken() (string, error) {
	return pay.Context.GetAccessToken()
}

//// GetSDK
//func (pay *Youzanyunpay) GetSDK() *sdk.SDK {
//	return sdk.NewSDK(pay.Context)
//}

// GetSDK
func (pay *Youzanyunpay) GetSDK() {
	YZYPay.Context.SDK = sdk.NewSDK(pay.Context)
}
