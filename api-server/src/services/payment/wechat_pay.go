package paymentSrv

import (
	"bytes"
	"errors"
	"io"
	"models"
	"net"
	"services/order"
	"strconv"
	"strings"
	"time"

	"github.com/chanxuehong/util"
	"github.com/chanxuehong/wechat/mch"
	"github.com/chinarun/utils"
)

var (
	ErrSign           = errors.New("签名失败")
	ErrParams         = errors.New("参数格式校验错误")
	ErrOrderCannotPay = errors.New("订单无法支付")
)

type WechatPay struct {
	notifyUrl    string
	order        OrderInfo
	params       map[string]string
	codeUrl      string //请求返回的支付地址, 用于生成二维码图片
	noticeParams map[string]string
}

const (
	appid      string = "wxb5636d699fc1b88b"
	mchid      string = "1227429802"
	apiKey     string = "F5483873309762D60EBE8F6E2F987AD9"
	appSecret  string = "1fea0e0ce8ecada2c9f4e319cdbd8d14"
	tradeType  string = "NATIVE"
	orderUrl   string = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	payTimeout string = "30m"
)

func NewWechatPay() (wp *WechatPay, err error) {

	notifyUrl, err := utils.Cfg.GetString("wechat_pay", "notify_url")
	if err != nil {
		return nil, err
	}

	wp = &WechatPay{
		notifyUrl:    notifyUrl,
		params:       make(map[string]string),
		noticeParams: make(map[string]string),
	}
	return
}

func (p *WechatPay) setParams() (err error) {
	defer func() {
		utils.Logger.Debug("%s", utils.Sdump(p.params))
	}()
	p.params["appid"] = appid
	p.params["mch_id"] = mchid
	p.params["nonce_str"] = utils.SubStr(utils.GetMd5Sha1(strconv.Itoa(int(time.Now().Unix()))), 0, 32)
	p.params["body"] = p.order.GetDesc()
	p.params["detail"] = p.order.GetDetail()
	p.params["out_trade_no"] = p.order.GetNo()
	p.params["total_fee"] = strconv.Itoa(p.order.GetPrice())
	p.params["spbill_create_ip"] = get_local_ip()
	p.params["notify_url"] = p.notifyUrl
	p.params["trade_type"] = tradeType
	p.params["product_id"] = p.order.GetProductId()

	timeoutDuration, _ := time.ParseDuration(payTimeout)
	time_expire := time.Now().Local().Add(timeoutDuration)
	p.params["time_expire"] = time_expire.Format("20060102150405")
	p.params["sign"] = mch.Sign(p.params, apiKey, nil)

	return err
}

func (p *WechatPay) validOrder(orderNo string) (err error) {
	order, err := orderSrv.NewOrder(orderNo)
	if err != nil {
		return
	}

	if err = order.Valid(); err != nil {
		return
	}

	p.order = order
	return
}

func (p *WechatPay) Pay(orderNo string) (codeUrl string, err error) {
	if err = p.validOrder(orderNo); err != nil {
		return
	}

	if err = p.setParams(); err != nil {
		return
	}

	if err = p.getCodeUrl(); err != nil {
		return
	}

	codeUrl = p.codeUrl
	return
}

func get_local_ip() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func (p *WechatPay) getCodeUrl() (err error) {
	proxy := mch.NewProxy(appid, mchid, apiKey, nil)

	resp, err := proxy.PostXML(orderUrl, p.params)
	if err != nil {
		utils.Logger.Error("paymentSrv.WechatPay.getCodeUrl : error : %s", err)
		return
	}

	if resp["return_code"] != "SUCCESS" || resp["result_code"] != "SUCCESS" {
		utils.Logger.Error("paymentSrv.WechatPay.getCodeUrl : %s .  result_code: %s ", resp["return_code"], resp["result_code"])
		return errors.New("获取code_url失败")
	}
	p.codeUrl = resp["code_url"]
	return
}

func (p *WechatPay) validNotice() (err error) {

	if p.noticeParams["return_code"] != "SUCCESS" || p.noticeParams["result_code"] != "SUCCESS" {
		return ErrParams
	}

	if p.noticeParams["appid"] != appid || p.noticeParams["mch_id"] != mchid {
		return ErrParams
	}

	if mch.Sign(p.noticeParams, apiKey, nil) != p.noticeParams["sign"] {
		return ErrSign
	}

	return
}

func (p *WechatPay) HandleNotify(reader io.Reader) (respXMLString string) {
	wechatNoticeParams, err := util.ParseXMLToMap(reader)
	utils.Logger.Debug("paymentSrv.WechatPay.HandleNotify : wechatNoticeParams : ", wechatNoticeParams)
	if err != nil {
		return
	}
	p.noticeParams = wechatNoticeParams

	bodyBuf := new(bytes.Buffer)
	retXml := make(map[string]string)

	retXml["return_code"] = "SUCCESS"

	if err := p.validNotice(); err != nil {
		utils.Logger.Debug("paymentSrv.WechatPay.HandleNotify : validNotice falied: %s", err.Error())
		retXml["return_msg"] = err.Error()
		util.FormatMapToXML(bodyBuf, retXml)
		return bodyBuf.String()
	}

	if err := p.validOrder(p.noticeParams["out_trade_no"]); err != nil {
		utils.Logger.Debug("paymentSrv.WechatPay.HandleNotify :validOrder failed, order_no : %s", p.noticeParams["out_trade_no"])
	}

	utils.Logger.Debug("paymentSrv.WechatPay.HandleNotify :order %s", utils.Sdump(p.order))
	updateCondition := map[string]interface{}{
		models.DB_ORDER_PAY_TIME:    time.Now().Local(),
		models.DB_ORDER_PAY_METHOD:  orderSrv.PAY_METHOD_WECHAT,
		models.DB_ORDER_PAY_ACCOUNT: p.noticeParams["open_id"],
	}

	if err := p.order.Update(updateCondition); err != nil {
		utils.Logger.Debug("paymentSrv.WechatPay.HandleNotify : p.order.Update failed conditions: %v", updateCondition)
	}

	retXml["return_msg"] = ""
	util.FormatMapToXML(bodyBuf, retXml)

	utils.Logger.Debug("paymentSrv.WechatPay.HandleNotify : bodyBuf: %v", bodyBuf)
	return bodyBuf.String()
}
