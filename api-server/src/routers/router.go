package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/api/user", &controllers.UserController{})
	beego.Router("/api/party", &controllers.PartyController{})
	beego.Router("/api/game", &controllers.GameController{})
	beego.Router("/api/reg", &controllers.RegController{})
	beego.Router("/api/msg", &controllers.MessageController{})
	beego.Router("/api/photo", &controllers.PhotoController{})
	beego.Router("/api/payment", &controllers.PaymentController{})

	beego.Router("/api/qrcode", &controllers.QrcodeController{})

	//payment
	beego.Router("/wechatpay/notify", &controllers.WechatPayNotify{})
	beego.Router("/alipay/notify", &controllers.AlipayNotifyController{})
}
