package sdk

import "github.com/goinggo/mapstructure"

type Trade struct {
	FullOrderInfo  interface{} `json:"full_order_info"` // 交易基础信息结构体
	RefundOrder    interface{} `json:"refund_order"`    // 订单退款信息结构体
	DeliveryOrder  interface{} `json:"delivery_order"`  // 订单发货详情结构体
	OrderPromotion interface{} `json:"order_promotion"` // 订单优惠详情结构体
}

func (sdk *SDK) GetTrade(tid string) (trade Trade, err error) {
	params := map[string]string{
		"tid": tid,
	}
	responseMap, err := sdk.Invoke("youzan.trade.get", "4.0.0", "GET", params, map[string]string{})
	if err != nil {
		return
	}

	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(responseMap, &trade); err != nil {
		return
	}
	return
}
