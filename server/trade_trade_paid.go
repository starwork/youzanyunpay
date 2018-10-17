package server

import (
	"encoding/json"
	"fmt"
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
	Orders     Orders     `json:"orders"`      //
	SourceInfo SourceInfo `json:"source_info"` //
	OrderInfo  OrderInfo  `json:"order_info"`  //
}

type BuyerInfo struct {
	BuyerPhone string `json:"buyer_phone"` //
}

type Orders struct {
	BuyerPhone    string `json:"buyer_phone"`    //
	GoodsUrl      string `json:"goods_url"`      //
	PicPath       string `json:"pic_path"`       //
	Oid           string `json:"oid"`            //
	Title         string `json:"title"`          //
	BuyerMessages string `json:"buyer_messages"` //
	Price         string `json:"price"`          //
	TotalFee      string `json:"total_fee"`      //
	Payment       string `json:"payment"`        //
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
	IsPayed string `json:"is_payed"` //
}

func (srv *Server) GetTradeTradePaid(msg ReqMsg) (m ReqMsgMsg, err error) {
	// 6. 处理消息体 —> 解码 msg ，反序列化消息结构体
	fmt.Println("msg.Msg", msg.Msg)
	json.Unmarshal([]byte(msg.Msg), &m)

	fmt.Println("m.QrInfo.QrID:", m.QrInfo.QrID)
	fmt.Println("m.QrInfo.QrName:", m.QrInfo.QrName)

	fmt.Println("m.FullOrderInfo.BuyerInfo.BuyerPhone:", m.FullOrderInfo.BuyerInfo.BuyerPhone)

	fmt.Println("m.FullOrderInfo.OrderInfo.Created:", m.FullOrderInfo.OrderInfo.Created)
	fmt.Println("m.FullOrderInfo.OrderInfo.ExpiredTime:", m.FullOrderInfo.OrderInfo.ExpiredTime)
	fmt.Println("m.FullOrderInfo.OrderInfo.OrderExtra:", m.FullOrderInfo.OrderInfo.OrderExtra)
	fmt.Println("m.FullOrderInfo.OrderInfo.OrderTags:", m.FullOrderInfo.OrderInfo.OrderTags)
	fmt.Println("m.FullOrderInfo.OrderInfo.PayTime:", m.FullOrderInfo.OrderInfo.PayTime)
	fmt.Println("m.FullOrderInfo.OrderInfo.PayType:", m.FullOrderInfo.OrderInfo.PayType)
	fmt.Println("m.FullOrderInfo.OrderInfo.Status:", m.FullOrderInfo.OrderInfo.Status)
	fmt.Println("m.FullOrderInfo.OrderInfo.StatusStr:", m.FullOrderInfo.OrderInfo.StatusStr)
	fmt.Println("m.FullOrderInfo.OrderInfo.Tid:", m.FullOrderInfo.OrderInfo.Tid)

	fmt.Println("m.FullOrderInfo.Orders.BuyerPhone:", m.FullOrderInfo.Orders.BuyerPhone)
	fmt.Println("m.FullOrderInfo.Orders.BuyerMessages:", m.FullOrderInfo.Orders.BuyerMessages)
	fmt.Println("m.FullOrderInfo.Orders.GoodsUrl:", m.FullOrderInfo.Orders.GoodsUrl)
	fmt.Println("m.FullOrderInfo.Orders.Oid:", m.FullOrderInfo.Orders.Oid)
	fmt.Println("m.FullOrderInfo.Orders.Payment:", m.FullOrderInfo.Orders.Payment)
	fmt.Println("m.FullOrderInfo.Orders.PicPath:", m.FullOrderInfo.Orders.PicPath)
	fmt.Println("m.FullOrderInfo.Orders.Price:", m.FullOrderInfo.Orders.Price)
	fmt.Println("m.FullOrderInfo.Orders.Title:", m.FullOrderInfo.Orders.Title)
	fmt.Println("m.FullOrderInfo.Orders.TotalFee:", m.FullOrderInfo.Orders.TotalFee)

	fmt.Println("m.FullOrderInfo.SourceInfo.Source.Platform:", m.FullOrderInfo.SourceInfo.Source.Platform)
	fmt.Println("m.FullOrderInfo.SourceInfo.Source.Platform:", m.FullOrderInfo.SourceInfo.Source.WxEntrance)

	return
}
