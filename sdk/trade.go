package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/yuyan2077/youzanyunpay/util"
)

type Trade struct {
	util.CommonError

	FullOrderInfo  interface{} `json:"full_order_info"` // 交易基础信息结构体
	RefundOrder    interface{} `json:"refund_order"`    // 订单退款信息结构体
	DeliveryOrder  interface{} `json:"delivery_order"`  // 订单发货详情结构体
	OrderPromotion interface{} `json:"order_promotion"` // 订单优惠详情结构体
}

func (sdk *SDK) GetTrade(tid string) (trade Trade, err error) {
	params := map[string]string{
		"tid": tid,
	}
	resp, err := sdk.Invoke("youzan.trade.get", "4.0.0", "GET", params, map[string]string{})
	if err != nil {
		return
	}

	var resultInterface map[string]Trade
	json.Unmarshal(resp, &resultInterface)
	if trade.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", trade.ErrCode, trade.ErrMsg)
		return
	}
	trade = resultInterface["response"]
	return
}
