package controllers

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/chinarun/utils"

	"services/order"
	"services/payment"
)

type AlipayNotifyController struct {
	beego.Controller
}

func (c *AlipayNotifyController) Get() {
}

func (c *AlipayNotifyController) Post() {
	status, orderId, buyerEmail, _ := c.AlipayNotify()

	if status {
		//支付成功

		order, err := orderSrv.NewOrder(orderId)
		if err != nil {
			c.Ctx.WriteString("failed")
		}

		alipay := paymentSrv.NewAlipay(order, "")
		resp := alipay.HandleNotify(buyerEmail)
		c.Ctx.WriteString(resp)
	}
	c.Ctx.WriteString("failed")
}

func (c *AlipayNotifyController) AlipayNotify() (status bool, orderId, buyerEmail, tradeNo string) {
	utils.Logger.Debug("requestBody----->%v", c.Ctx.Request.Body)
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		status = false
		return
	}

	bodyStr, _ := url.QueryUnescape(string(body))
	postArray := strings.Split(bodyStr, "&")
	params := map[string]string{}
	utils.Logger.Debug("alipay notify params---->%v", params)
	for _, v := range postArray {
		detail := strings.Split(v, "=")
		params[detail[0]] = detail[1]
	}

	if len(params["out_trade_no"]) == 0 {

		return false, params["out_trade_no"], params["buyer_email"], params["trade_no"]

	} else {
		//验证回调请求是否合法
		if params["sign"] == paymentSrv.Sign(params) {
			//交易成功
			if params["trade_status"] == "TRADE_FINISHED" || params["trade_status"] == "TRADE_SUCCESS" {
				return true, params["out_trade_no"], params["buyer_email"], params["trade_no"]
			} else {
				status = false
				return
			}
		} else {
			tradeNo = fmt.Sprintln(params)
			status = false
			return
		}
	}

}
