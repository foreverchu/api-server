package controllers

import (
	"encoding/base64"
	"services/qrcode"
)

type QrcodeController struct {
	BaseController
}

func (c *QrcodeController) Prepare() {
	//c.BaseController.Prepare()
}

func (c *QrcodeController) Finish() {
	defer c.BaseController.Finish()
}

func (c *QrcodeController) Post() {
	content := c.GetString("content")
	size, _ := c.GetInt("size")
	qrcode := qrcodeSrv.NewQrcode(content, size)
	var png []byte
	png, err := qrcode.Generate()
	if err != nil {
		c.Data["json"] = struct {
			ErrMsg string `json:"errMsg"`
		}{
			ErrMsg: err.Error(),
		}
		c.ServeJson()
	}
	c.Data["png"] = base64.StdEncoding.EncodeToString(png)
	c.TplNames = "qrcode.tpl"
}
