package server

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/message"
	"github.com/yuyan2077/youzanyunpay/context"
	"github.com/yuyan2077/youzanyunpay/util"
	"net/http"
)

//Server struct
type Server struct {
	*context.Context

	debug bool

	openID string

	messageHandler func(message.MixMessage) *message.Reply

	requestRawXMLMsg  []byte
	requestMsg        message.MixMessage
	responseRawXMLMsg []byte
	responseMsg       interface{}

	random    []byte
	nonce     string
	timestamp int64
}

type ReqMsg struct {
	ClientID  string `json:"client_id"` // 对应开发者后台的client_id
	ID        string `json:"id"`        // 业务消息的标识: 如 订单消息为订单编号,会员卡消息为会员卡id标识
	KdtID     int    `json:"kdt_id"`    // 店铺ID
	KdtName   string `json:"kdt_name"`  // 店铺名
	Mode      int    `json:"mode"`      // 1-自用型/工具型/平台型消息；0-签名模式消息
	Msg       string `json:"msg"`       // 经过UrlEncode（UTF-8）编码的消息对象，具体参数请看本页中各业务消息结构文档
	MsgID     string `json:"msg_id"`    // 消息id
	SendCount int    `json:"sendCount"` // 重发的次数
	Sign      string `json:"sign"`      // 防伪签名 ：MD5(client_id+msg+client_secrect) ; MD5 方法可通过搜索引擎搜索得到或者可参考 MD5
	Status    string `json:"status"`    // 消息状态，对应消息业务类型。如TRADE_ORDER_STATE-订单状态事件，对应有等待买家付款（WAIT_BUYER_PAY）、等待卖家发货（WAIT_SELLER_SEND_GOODS）等多种状态，详细可参考 消息结构里的描述
	Test      bool   `json:"test"`      // false-非测试消息，true- 测试消息 ；PUSH服务会定期通过发送测试消息检查开发者服务是否正常
	Type      string `json:"type"`      // 消息业务类型：TRADE_ORDER_STATE-订单状态事件，TRADE_ORDER_REFUND-退款事件，TRADE_ORDER_EXPRESS-物流事件，ITEM_STATE-商品状态事件，ITEM_INFO-商品基础信息事件，POINTS-积分，SCRM_CARD-会员卡（商家侧），SCRM_CUSTOMER_CARD-会员卡（用户侧），TRADE-交易V1，ITEM-商品V1
	Version   string `json:"version"`   // 消息版本号，为了解决顺序性的问题 ，高版本覆盖低版本

}

//{"code":0,"msg":"success"}
type RespMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//NewServer init
func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

// SetDebug set debug field
func (srv *Server) SetDebug(debug bool) {
	srv.debug = debug
}

//Serve 处理有赞云的请求消息
func (srv *Server) Serve() error {
	_, err := srv.handleRequest()
	if err != nil {
		return err
	}

	return nil
}

//HandleRequest 处理有赞云的请求
func (srv *Server) handleRequest() (m map[string]interface{}, err error) {

	var msg ReqMsg
	msg, err = srv.getMessage()
	if err != nil {
		return
	}

	// 1. 判断消息是否测试  —> 解析 test
	if msg.Test == true {
		srv.SendResponseMsg()
		return
	}

	// 2. 判断消息推送的模式 —> 解析 mode
	if msg.Mode == 0 {
		srv.SendResponseMsg()
		return
	}

	// 3. 判断消息是否伪造 —> 解析 sign
	msgSignatureGen := util.Signature(msg.ClientID, msg.Msg, srv.AppSecret)
	if msg.Sign != msgSignatureGen {
		return nil, fmt.Errorf("消息不合法，验证签名失败")
	}

	// 4. 判断消息版本  —> 解析 version
	//fmt.Println(msg.Version)

	// 5. 判断消息的业务 —> 解析 type
	if msg.Type != "TRADE_ORDER_STATE" {
		srv.SendResponseMsg()
		return
	}

	// 6. 处理消息体 —> 解码 msg ，反序列化消息结构体
	fmt.Println("msg.Msg", msg.Msg)
	//var m map[string]interface{}
	json.Unmarshal([]byte(msg.Msg), &m)
	fmt.Println("m", m)

	// 7. 返回接收成功标识 {"code":0,"msg":"success"}
	srv.SendResponseMsg()

	return
}

//getMessage 解析有赞云推送的消息
func (srv *Server) getMessage() (msg ReqMsg, err error) {
	if err = json.NewDecoder(srv.Request.Body).Decode(&msg); err != nil {
		return
	}
	return
}

func (srv *Server) SendResponseMsg() {
	respMsg := RespMsg{0, "success"}
	js, err := json.Marshal(respMsg)
	if err != nil {
		http.Error(srv.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	srv.Writer.Header().Set("Content-Type", "application/json")
	srv.Writer.Write(js)
}

//SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(message.MixMessage) *message.Reply) {
	srv.messageHandler = handler
}
