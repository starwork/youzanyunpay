package youzanyunpay

import (
	"github.com/yuyan2077/youzanyunpay/context"
	"github.com/yuyan2077/youzanyunpay/sdk"
	"github.com/yuyan2077/youzanyunpay/server"
	"net/http"
	"sync"
)

type Youzanyunpay struct {
	Context *context.Context
}

type Config struct {
	AppID     string
	AppSecret string
	KdtID     int
}

// GetServer 消息管理
func (pay *Youzanyunpay) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	pay.Context.Request = req
	pay.Context.Writer = writer
	return server.NewServer(pay.Context)
}

func NewYouzanyunpay(cfg *Config) *Youzanyunpay {
	ctx := new(context.Context)
	copyConfigToContext(cfg, ctx)
	return &Youzanyunpay{ctx}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.KdtID = cfg.KdtID
	context.SetAccessTokenLock(new(sync.RWMutex))
	//context.SetJsAPITicketLock(new(sync.RWMutex))
}

//GetAccessToken 获取access_token
func (pay *Youzanyunpay) GetAccessToken() (string, error) {
	return pay.Context.GetAccessToken()
}

// GetSDK
func (pay *Youzanyunpay) GetSDK() *sdk.SDK {
	return sdk.NewSDK(pay.Context)
}
