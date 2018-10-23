package sdk

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
	resultInterface, err := sdk.Invoke("youzan.pay.qrcode.create", "3.0.0", "POST", params, map[string]string{})
	if err != nil {
		return
	}

	createQrcodeResp = resultInterface.Response.(CreateQrcodeResp)
	return
}
