package server

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type ReqMsgMsg struct {
	QrInfo        QrInfo        `json:"qr_info"`         //
	FullOrderInfo FullOrderInfo `json:"full_order_info"` //
}

type QrInfo struct {
	QrID   int    `json:"qr_id"`   //
	QrName string `json:"qr_name"` //
}

type FullOrderInfo struct {
	BuyerInfo  BuyerInfo  `json:"buyer_info"`  //
	PayInfo    PayInfo    `json:"pay_info"`    //
	SourceInfo SourceInfo `json:"source_info"` //
	OrderInfo  OrderInfo  `json:"order_info"`  //
}

type BuyerInfo struct {
	BuyerPhone string `json:"buyer_phone"` //
}

type PayInfo struct {
	Payment string `json:"payment"` //
}

type SourceInfo struct {
	Source Source `json:"source"` //
}

type Source struct {
	Platform   string `json:"platform"`    //
	WxEntrance string `json:"wx_entrance"` //
}

type OrderInfo struct {
	OrderExtra  OrderExtra `json:"order_extra"`  //
	Created     string     `json:"created"`      //
	StatusStr   string     `json:"status_str"`   //
	ExpiredTime string     `json:"expired_time"` //
	Tid         string     `json:"tid"`          //
	PayTime     string     `json:"pay_time"`     //
	PayType     int        `json:"pay_type"`     //
	Status      string     `json:"status"`       //
	OrderTags   OrderTags  `json:"order_tags"`   //
}

type OrderExtra struct {
	BuyerName string `json:"buyer_name"` //
}

type OrderTags struct {
	IsPayed bool `json:"is_payed"` //
}

func (srv *Server) GetTradeTradePaid(msg string) (msgMsg ReqMsgMsg, err error) {
	// 6. 处理消息体 —> 解码 msg ，反序列化消息结构体
	m, _ := url.QueryUnescape(msg)
	fmt.Println(m)
	json.Unmarshal([]byte(m), &msgMsg)

	fmt.Println("m.QrInfo.QrID:", msgMsg.QrInfo.QrID)
	fmt.Println("m.QrInfo.QrName:", msgMsg.QrInfo.QrName)

	fmt.Println("m.FullOrderInfo.BuyerInfo.BuyerPhone:", msgMsg.FullOrderInfo.BuyerInfo.BuyerPhone)

	fmt.Println("m.FullOrderInfo.OrderInfo.Created:", msgMsg.FullOrderInfo.OrderInfo.Created)
	fmt.Println("m.FullOrderInfo.OrderInfo.ExpiredTime:", msgMsg.FullOrderInfo.OrderInfo.ExpiredTime)
	fmt.Println("m.FullOrderInfo.OrderInfo.OrderExtra.BuyerName:", msgMsg.FullOrderInfo.OrderInfo.OrderExtra.BuyerName)
	fmt.Println("m.FullOrderInfo.OrderInfo.OrderTags.IsPayed:", msgMsg.FullOrderInfo.OrderInfo.OrderTags.IsPayed)
	fmt.Println("m.FullOrderInfo.OrderInfo.PayTime:", msgMsg.FullOrderInfo.OrderInfo.PayTime)
	fmt.Println("m.FullOrderInfo.OrderInfo.PayType:", msgMsg.FullOrderInfo.OrderInfo.PayType)
	fmt.Println("m.FullOrderInfo.OrderInfo.Status:", msgMsg.FullOrderInfo.OrderInfo.Status)
	fmt.Println("m.FullOrderInfo.OrderInfo.StatusStr:", msgMsg.FullOrderInfo.OrderInfo.StatusStr)
	fmt.Println("m.FullOrderInfo.OrderInfo.Tid:", msgMsg.FullOrderInfo.OrderInfo.Tid)

	fmt.Println("m.FullOrderInfo.PayInfo.Payment:", msgMsg.FullOrderInfo.PayInfo.Payment)

	fmt.Println("m.FullOrderInfo.SourceInfo.Source.Platform:", msgMsg.FullOrderInfo.SourceInfo.Source.Platform)
	fmt.Println("m.FullOrderInfo.SourceInfo.Source.WxEntrance:", msgMsg.FullOrderInfo.SourceInfo.Source.WxEntrance)
	return
}
