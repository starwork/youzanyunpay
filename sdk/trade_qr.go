package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/yuyan2077/youzanyunpay/util"
)

type CreateQrcodeResp struct {
	util.CommonError

	QrUrl  string `json:"qr_url"`
	QrCode string `json:"qr_code"`
	QrID   string `json:"qr_id"`
}

func (sdk *SDK) CreateQrcode(qrName, qrPrice, qrType string) (createQrcodeResp CreateQrcodeResp, err error) {
	params := map[string]string{
		"qr_name":  qrName,
		"qr_price": qrPrice,
		"qr_type":  qrType,
	}
	resp, err := sdk.Invoke("youzan.pay.qrcode.create", "3.0.0", "POST", params, map[string]string{})
	if err != nil {
		return
	}

	var resultInterface map[string]CreateQrcodeResp
	json.Unmarshal(resp, &resultInterface)
	if createQrcodeResp.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", createQrcodeResp.ErrCode, createQrcodeResp.ErrMsg)
		return
	}
	createQrcodeResp = resultInterface["response"]
	if createQrcodeResp.QrCode == "" {
		return createQrcodeResp, fmt.Errorf("GetUserAccessToken error")
	}
	return
}
