package context

import (
	"net/http"
	"sync"
)

// Context struct
type Context struct {
	AppID     string
	AppSecret string
	KdtID     int
	Token     Token

	Writer  http.ResponseWriter
	Request *http.Request

	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	////jsAPITicket 读写锁 同一个AppID一个
	//jsAPITicketLock *sync.RWMutex
}

type Token struct {
	Token        string
	TokenTimeOut int64
}

// Query returns the keyed url query value if it exists
func (ctx *Context) Query(key string) string {
	value, _ := ctx.GetQuery(key)
	return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (ctx *Context) GetQuery(key string) (string, bool) {
	req := ctx.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

//// SetJsAPITicketLock 设置jsAPITicket的lock
//func (ctx *Context) SetJsAPITicketLock(lock *sync.RWMutex) {
//	ctx.jsAPITicketLock = lock
//}
//
//// GetJsAPITicketLock 获取jsAPITicket 的lock
//func (ctx *Context) GetJsAPITicketLock() *sync.RWMutex {
//	return ctx.jsAPITicketLock
//}
