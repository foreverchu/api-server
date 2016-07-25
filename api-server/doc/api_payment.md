支付操作接口列表：
----

接口列表：

接口名称|接口描述|开发情况
---|---|---
WechatPay|[微信支付](#WechatPay)|[YES]
WechatPayState|[微信订单状态查询](#WechatPayState)|[YES]
Pay|[支付](#Pay)|[YES]


---

<div id="Pay"></div>
支付(Pay)
请求样例:
```json
{
    "cmd":"Pay",
    "order_no":"1510281411335281083827",
    "type": 1,
    "return_url":"http://test.www.wanbisai.com"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| WechatPay
order_no|string|是|定单号
type|int|是|支付类型。1表示支付宝
return_url|否|支付宝支付时必须

返回样例:
```json
{
  "pay_info": {
    "alipay_url":"https://mapi.alipay.com/gateway.do?_input_charset=utf-8&body=%E4%B8%BA-%E5%85%85%E5%80%BC0.01%E5%85%83&notify_url=http%3A%2F%2F127.0.0.1&out_trade_no=4341034398&partner=2088811198755564&payment_type=1&return_url=http%3A%2F%2F127.0.0.1&seller_email=souvenir%40chinarun.com&service=create_direct_pay_by_user&subject=-&total_fee=0.01&sign=968c5bdf51d9afb82170ba16c7d0b874&sign_type=MD5"
    },
  "msg": "",
  "result": 0
}

```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-15
msg|string|否|如果错误,可以将具体错误信息返回
pay_info|string|是|支付信息
alipay_url|string|否|支付地址，支付类型为1（支付宝支付）时，必须有返回，
wechat_pay_url|string|否|支付地址，支付类型为2时必须返回

---

<div id="WechatPay"></div>
用户注册(UserRegister)
请求样例:
```json
{
    "cmd":"WechatPay",
    "order_no":"1510281411335281083827"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| WechatPay
order_no|string|是|定单号
notify_url|string|是|回调地址


返回样例:
```json
{
  "code_url": "weixin://wxpay/bizpayurl?pr=D7TOzCi",
  "msg": "",
  "result": 0
}

```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-15
msg|string|否|如果错误,可以将具体错误信息返回
code_url|string|是|用于生成二维码的支付地址


---
<div id="WechatPayState"></div>
微信支付状态查询 (WechatPayState)

请求样例:


请求样例:
{
    "cmd":"WechatPayState",
    "order_no":"1510281411335281083827"
"
}

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| WechatPayState
order_no|string| 是| 订单号



返回样例:
```json
{
  "msg": "",
  "order_no": "1510281555449126260705",
  "pay_state": 0,
  "pay_state_desc": "等待支付",
  "result": 0
}
```

支付状态说明:
```
const(
    ORDER_PAY_STATE_WAIT_PAY PayState = 0 //等待支付
    ORDER_PAY_STATE_PAYED    PayState = 1 //已支付
    ORDER_PAY_STATE_CANCELED PayState = 2 //已取消
    ORDER_PAY_STATE_REFUNDED PayState = 3 //已退款
)

```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-15
msg|string|否|如果错误,可以将具体错误信息返回
order_no|string|是|订单号
pay_state|int|是|订单状态
pay_state_desc|string|是|订单状态文字描述

