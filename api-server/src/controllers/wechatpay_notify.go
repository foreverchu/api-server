package controllers

import (
	"bytes"
	"services/payment"

	"github.com/astaxie/beego"
	"github.com/chinarun/utils"
)

type WechatPayNotify struct {
	beego.Controller
}

// wechat payment notify
func (c *WechatPayNotify) Post() {
	defer func() {
		utils.Logger.Error("controllers.PaymentController.WechatPayNotify : requestBody : %s", string(c.Ctx.Input.RequestBody))
	}()
	xmlReader := bytes.NewReader(c.Ctx.Input.RequestBody)
	wechatPay, err := paymentSrv.NewWechatPay()
	if err != nil {
		utils.Logger.Error("controllers.PaymentController.WechatPayNotify : paymentSrv.NewWechatPay() error : %s", err.Error())
	}
	respStr := wechatPay.HandleNotify(xmlReader)

	utils.Logger.Error("controllers.PaymentController.WechatPayNotify : respStr: %s", respStr)
	c.Ctx.WriteString(respStr)
	return
}
