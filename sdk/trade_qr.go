package sdk

import "github.com/goinggo/mapstructure"

type CreateQrcodeResp struct {
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
	responseMap, err := sdk.Invoke("youzan.pay.qrcode.create", "3.0.0", "POST", params, map[string]string{})
	if err != nil {
		return
	}
	//将 map 转换为指定的结构体
	if err = mapstructure.Decode(responseMap, &createQrcodeResp); err != nil {
		return
	}
	return
}
