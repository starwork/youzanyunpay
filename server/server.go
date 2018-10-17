package server

import (
	"encoding/json"
	"fmt"
	"github.com/yuyan2077/youzanyunpay/context"
	"github.com/yuyan2077/youzanyunpay/util"
	"net/http"
)

//Server struct
type Server struct {
	*context.Context

	debug bool

	openID string

	//requestRawXMLMsg  []byte
	//requestMsg        message.MixMessage
	//responseRawXMLMsg []byte
	//responseMsg       interface{}

	random    []byte
	nonce     string
	timestamp int64
}

//{"code":0,"msg":"success"}
type RespMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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
	Version   int    `json:"version"`   // 消息版本号，为了解决顺序性的问题 ，高版本覆盖低版本

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
func (srv *Server) Serve() (m ReqMsgMsg, err error) {
	m, err = srv.handleRequest()
	if err != nil {
		return
	}
	return
}

//HandleRequest 处理有赞云的请求
//{
//"msg": "%7B%22full_order_info%22%3A%7B%22remark_info%22%3A%7B%22buyer_message%22%3A%22%22%7D%2C%22pay_info%22%3A%7B%22outer_transactions%22%3A%5B%224200000168201810166190128957%22%5D%2C%22post_fee%22%3A%220.00%22%2C%22total_fee%22%3A%220.01%22%2C%22payment%22%3A%220.01%22%2C%22transaction%22%3A%5B%22181016214103000041%22%5D%7D%2C%22source_info%22%3A%7B%22is_offline_order%22%3Afalse%2C%22source%22%3A%7B%22platform%22%3A%22wx%22%2C%22wx_entrance%22%3A%22direct_buy%22%7D%7D%2C%22order_info%22%3A%7B%22consign_time%22%3A%7B%7D%2C%22order_extra%22%3A%7B%22is_from_cart%22%3A%22false%22%7D%2C%22created%22%3A%222018-10-16+21%3A41%3A01%22%2C%22status_str%22%3A%22%E5%B7%B2%E6%94%AF%E4%BB%98%22%2C%22expired_time%22%3A%222018-10-16+22%3A11%3A01%22%2C%22success_time%22%3A%7B%7D%2C%22type%22%3A6%2C%22tid%22%3A%22E20181016214101073500004%22%2C%22confirm_time%22%3A%7B%7D%2C%22pay_time%22%3A%222018-10-16+21%3A41%3A10%22%2C%22update_time%22%3A%222018-10-16+21%3A41%3A11%22%2C%22is_retail_order%22%3Afalse%2C%22pay_type%22%3A10%2C%22team_type%22%3A1%2C%22refund_state%22%3A0%2C%22close_type%22%3A0%2C%22status%22%3A%22TRADE_PAID%22%2C%22express_type%22%3A9%2C%22order_tags%22%3A%7B%22is_payed%22%3Atrue%2C%22is_secured_transactions%22%3Atrue%7D%7D%7D%7D",
//"kdt_name": "自由工作室",
//"test": false,
//"sign": "060c5501ca10ac01170525fc8e07da67",
//"sendCount": 3,
//"type": "trade_TradePaid",
//"version": 1539697271,
//"client_id": "d3ccc77bb8b5bf4e08",
//"mode": 1,
//"kdt_id": 41566520,
//"id": "E20181016214101073500004",
//"msg_id": "59b2499c-59d7-4f94-8406-8b013318fac1",
//"status": "PAID"
//}
func (srv *Server) handleRequest() (m ReqMsgMsg, err error) {

	var msg ReqMsg
	msg, err = srv.getMessage()

	if err != nil {
		//srv.SendResponseMsg()
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
		return m, fmt.Errorf("消息不合法，验证签名失败" + "msg.Sign:" + msg.Sign + "|" + "msgSignatureGen:" + msgSignatureGen)
	}

	// 4. 判断消息版本  —> 解析 version
	//fmt.Println(msg.Version)

	// 5. 判断消息的业务 —> 解析 type
	if msg.Type != "trade_TradePaid" {
		srv.SendResponseMsg()
		return m, fmt.Errorf("消息不合法，不是想要的类型" + msg.Type)
	}

	// 6. 处理消息体 —> 解码 msg ，反序列化消息结构体
	switch msg.Type {
	case "trade_TradePaid":
		srv.GetTradeTradePaid(msg)
	}

	// 7. 返回接收成功标识 {"code":0,"msg":"success"}
	srv.SendResponseMsg()

	return
}

//getMessage 解析有赞云推送的消息
func (srv *Server) getMessage() (msg ReqMsg, err error) {
	if err = json.NewDecoder(srv.Request.Body).Decode(&msg); err != nil {
		return msg, fmt.Errorf("从body中解析json失败,err=%v", err)
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

////SetMessageHandler 设置用户自定义的回调方法
//func (srv *Server) SetMessageHandler(handler func(message.MixMessage) *message.Reply) {
//	srv.messageHandler = handler
//}
