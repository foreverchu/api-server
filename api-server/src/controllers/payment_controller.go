package controllers

import (
	"err_code"
	"fmt"
	"services/order"
	"services/payment"
)

type PaymentController struct {
	BaseController
}

func (c *PaymentController) Prepare() {
	fmt.Print("hit1\n")
	c.BaseController.Prepare()
}

func (c *PaymentController) Finish() {
	defer c.BaseController.Finish()
}
func (c *PaymentController) Post() {
	retJson := init_retJson()
	fmt.Println(c.cmd)
	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "Pay":
			c.Data["json"] = c.Pay()
		case "WechatPay":
			c.Data["json"] = c.WechatPay()
		case "WechatPayState":
			c.Data["json"] = c.WechatPayState()
		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}
	c.ServeJson()
}

func (c *PaymentController) Pay() map[string]interface{} {
	retJson := init_retJson()
	orderNo, _ := c.js.Get("order_no").String()
	payType, _ := c.js.Get("type").Int()
	returnUrl, _ := c.js.Get("return_url").String()

	order, err := orderSrv.NewOrder(orderNo)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	var pay paymentSrv.Pay
	switch payType {
	case 1:
		pay = paymentSrv.NewAlipay(order, returnUrl)
	default:
		retJson_edit(retJson, err_code.InvalidData, "未知支付类型")
		return retJson
	}
	if pay.IsOrderValid() != true {
		retJson_edit(retJson, err_code.InvalidOrderNo, "订单失效")
		return retJson
	}
	retJson["pay_info"] = pay.Pay()
	return retJson
}

func (c *PaymentController) WechatPay() map[string]interface{} {
	retJson := init_retJson()
	orderNo, _ := c.js.Get("order_no").String()

	wechatPay, err := paymentSrv.NewWechatPay()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	var codeUrl string
	if codeUrl, err = wechatPay.Pay(orderNo); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	retJson["code_url"] = codeUrl
	return retJson
}

func (c *PaymentController) WechatPayState() map[string]interface{} {
	retJson := init_retJson()
	orderNo, _ := c.js.Get("order_no").String()

	os, err := orderSrv.NewOrderStateWithOrderNo(orderNo)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	ps := os.State()
	retJson["order_no"] = orderNo
	retJson["pay_state"] = ps
	retJson["pay_state_desc"] = fmt.Sprintf("%s", ps)
	return retJson
}
