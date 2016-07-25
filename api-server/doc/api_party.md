party（赛事）操作接口列表
---

接口名称|接口描述|开发情况
---|---|---
PartyQueryRegOrders|[查询赛事报名订单](#PartyQueryRegOrders)|[YES]
PartyQueryRegResult|[获取赛事报名统计资料](#PartyQueryRegResult)|[YES]
PartyCreate|[赛事创建](#PartyCreate)|[YES]
PartyEdit|[赛事编辑](#PartyEdit)|[YES]
PartyList|[赛事列表](#PartyList)|[YES]
PartyQuery|[赛事查询（可以查自己创建到赛事,或者全部赛事列表)](#PartyQuery)|[YES]
PartyClose|[赛事关闭](#PartyClose)|[YES]
PartyQueryGames|[查询赛事中所有比赛](#PartyQueryGames)|[ YES ]
PartyStateUpdate|[赛事审核](#PartyStateUpdate)|[ YES ]



---

<div id="PartyQueryRegOrders"></div>
查询赛事报名订单 (PartyQueryRegOrders)
请求样例:
```json
{"cmd":"PartyQueryRegOrders"
,"party_id":"12345"
,"order_no":""
,"certificate_type": x
,"certificate_no": "xxx"
,"pay_status": x
,"page_no"：1
,"page_size":10
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyQueryRegOrders
party_id|string|是|赛事id
order_no|string|否|要查询的订单id, 如果为空则按下面证件信息查询
certificate_type|uint8|否|身份证件类型, 0:身份证, 1:护照, 2:军官证, 3:台胞证, 4:港澳通行证, 5:回乡证
certificate_no|string|否|证件号码， 若为空，且订单号为空，则查询所有指定赛事订单
page_no|int|否|分页, 默认为0
page_size|int|否|每页几条订单信息, 默认为20
pay_status|uint8|否|付款状态(0表示等待支付,1表示已支付,2表示已取消,3表示已退款), 无此参数表示查询所有支付状态的订单, 这个参数只有在order_no和证件信息都为空的情况下才有效

返回样例:

```json
{"result":0
,"msg":""
,"orders":{[
    "game_id":"xxxx"
    ,"game_name":""
    ,"order_no":""
    ,"submit_time":""
    ,"pay_time":""
    ,"refund_time":""
    ,"cancel_time":""
    ,"price":xxx
    ,"currency_type":x
    ,"pay_method":xxxx
    ,"pay_account":""
    ,"user_id":""
    ,"players":[
      {
        "id":"xxxx"
        ,"name":""
        ,"gender":x
        ,"certificate_type": x
        ,"certificate_no": "xxx"
      }
    ]
]}
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-1,-10,-11
msg|string|否|如果错误,可以将具体错误信息返回
game_id|string|是|比赛id
game_name|string|是|比赛名称
order_no|string|是|订单号
submit_time|string|是|订单提交时间
pay_time|string|是|订单支付时间，为"0001-01-01T00:00:00Z"表示没有支付
refund_time|string|是|订单退款时间，为"0001-01-01T00:00:00Z"表示没有退款
cancel_time|string|是|订单取消时间，为"0001-01-01T00:00:00Z"表示没有取消
price|float|是|订单价格
currency_type|uint8|否|付款币种， 1:人民币, 2：美元
pay_method|uint8|否|支付方法， 1：支付宝，2：微信支付，3：Paypal，4：其他
pay_account|string|否|支付账号
user_id|string|是|创建订单的Wordpress用户id
id|string|是|选手id
name|string|是|选手姓名
gender|uint8|是|性别，1：男，2：女
certificate_type|uint8|是|身份证件类型, 0:身份证, 1:护照, 2:军官证, 3:台胞证, 4:港澳通行证, 5:回乡证
certificate_no|string|是|证件号码， 若为空，且订单号为空，则查询所有指定赛事订单


---

<div id="PartyQueryRegResult"></div>
获取赛事报名统计资料 (PartyQueryRegResult)
请求样例:
```json
{"cmd":"PartyQueryRegResult"
,"party_id":"12345"
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyQueryRegResult
party_id|string|是|赛事id



返回样例:
```json
{"result":0
,"msg":""
,"games":{[
    ,"game_id":"xxxx"
    ,"name":""
    ,"payed_player_count":xxxx 
    ,"wait_pay_player_count":xxx
]}
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-1,-10,-11
msg|string|否|如果错误,可以将具体错误信息返回
payed_player_count|int|是|比赛付费人数
wait_pay_player_count|int|是|等待付费人数
game_id|string|是|比赛id
name|string|是|比赛名称

---


<div id="PartyCreate"></div>
创建赛事 (PartyCreate)
请求样例:
```json
{"cmd":"PartyCreate"
   ,"party_name":"xx"
   ,"limitation":1
   ,"limitation_type":1
   ,"country":""
   ,"province":""
   ,"city":""
   ,"addr":""
   ,"loc_long":xx.xx
   ,"loc_lat":xx.xx
   ,"reg_start_time":""
   ,"reg_end_time":""
   ,"start_time":""
   ,"end_time":""
   ,"slogan":""
   ,"like":256
   ,"type":""
   ,"introduction":""
   ,"website":""
   ,"schedule":""
   ,"score":9.5
   ,"signup_male":356
   ,"signup_female":35
   ,"price":""
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyCreate
party_name|string|是|赛事名称（唯一）
limitation|int|是|赛事名额限制
limitation_type|uint8|是|赛事名额限制类型
country|string|是|国家
province|string|否|省份
city|string|否|城市
addr|string|是|地址
loc_long|float|否|精度
loc_lat|float|否|维度
extra_info_json|string|否|扩展信息（json样式到字符串）
reg_start_time|string|是|报名开始时间
reg_end_time|string|是|报名截止时间
start_time|string|是|赛事开始时间
end_time|string|是|赛事结束时间
detail|int|否|是否有赛事详情信息
slogan|string|否|赛事详情精彩短评
like|int|否|想跑
type|string|否|赛别
introduction|string|否|介绍
website|string|否|官网
schedule|string|否|日程
score|string|否|评分
signup_male|int|否|已报名（男）
signup_female|int|否|已报名（女）
price|string|否|价格


返回样例:
```json
{"result":0
,"party_id":""
,"msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-1,-7,-8
msg|string|否|如果错误,可以将具体错误信息返回
party_id|string|是|生成到赛事编号



---



<div id="PartyEdit"></div>
编辑赛事 (PartyEdit)
请求样例:
```json
{"cmd":"PartyEdit"
   ,"party_id":"xx"
   ,"party_name":"xx"
   ,"limitation": 1
   ,"limitation_type": 1
   ,"country":""
   ,"province":""
   ,"city":""
   ,"addr":""
   ,"loc_long":xx.xx
   ,"loc_lat":xx.xx
   ,"website":""
   ,"extra_info_json":""
   ,"reg_start_time":""
   ,"reg_end_time":""
   ,"start_time":""
   ,"end_time":""

   ,"slogan":""
   ,"like":256
   ,"type":""
   ,"introduction":""
   ,"website":""
   ,"schedule":""
   ,"score":9.5
   ,"signup_male":356
   ,"signup_female":35
   ,"price":""
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyEdit
party_id|string|是|赛事id
party_name|string|否|赛事名称（唯一）
limitation|int|否|赛事名额限制
limitation_type|uint8|否|限制类型:每个game的报名人数限制,总名额限制,衣服型号类型限制,或以上3种组合
country|string|否|国家
province|string|否|省份
city|string|否|城市
addr|string|否|地址
loc_long|float|否|精度
loc_lat|float|否|维度
extra_info_json|string|否|扩展信息（json样式到字符串）
reg_start_time|string|否|报名开始时间
reg_end_time|string|否|报名截止时间
start_time|string|否|赛事开始时间
end_time|string|否|赛


返回样例:
```json
{"result":0
,"msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-1,-9
msg|string|否|如果错误,可以将具体错误信息返回


---

<div id="PartyList"></div>
赛事列表 (PartyList)
请求样例:
```json
{"cmd":"PartyList"
,"page_no"：1
,"page_size":10
,"order_by":""
,"include_close":0
,"user_id":"123",
,"type":0
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyList
valid_state|int| 否| 0表示待审核，1表示已审核通过，2表示审核不通过
month|int|否|表示查询的月份
country|string| 否| 国家
province|string| 否| 省份
city|string| 否| 城市
tags_string|string|否|标签
page_no|int|否|分页, 默认为0
page_size|int|否|每页几条赛事信息, 默认为10
order_by|string|否|排序字段，默认为start_time，还可选reg_start_time, 为倒序
include_close|int|否|是否包含已关闭赛事，如果为1表示包含，0为不包含，默认不包含
user_id|string|否|与用户相关的赛事。user_id为空，查询所有用户；user_id非空，只查该用户的赛事。默认为空。
type|int|否|赛事列表类型。0 - 普通赛事列表（所有类型），1 － user_id创建的赛事，2 － user_id报名参加的赛事，3 － user_id关注的赛事。   


返回样例:
```json
{
  "result":0
  "msg": "",
  "count":xxx
  "partylist": [
    {
      "Id": 2,
      "Name": "test",
      "Country": "cn",
      "Province": "sh",
      "City": "sh",
      "Addr": "wenhui st",
      "LocLong": 11.1,
      "LocLat": 11.1,
      "RegStartTime": "2015-01-01T01:01:01+08:00",
      "RegEndTime": "2015-01-01T01:01:01+08:00",
      "StartTime": "2015-01-01T01:01:01+08:00",
      "EndTime": "2015-01-01T01:01:01+08:00",
      "CloseTime": "0001-01-01T00:00:00Z",
      "Limitation": 1000,
      "LimitationType": 1,
      "ExtraInfoJson": "",
      "UserId": xxx,
    }
  ],
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-1,-10,-11
msg|string|否|如果错误,可以将具体错误信息返回
count|int|否|返回赛事个数
partylist|json|否|所有赛事列表,如果没有则为空


---




<div id="PartyQuery"></div>
赛事查询 (PartyQuery)
请求样例:
```json
{"cmd":"PartyQuery"
,"party_id":""
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyQuery
party_id|string|是|赛事id


返回样例:
```json
{
  "result": 0,
  "msg": "",
  "party":
  {
    "Id": 4,
    "Name": "axiba",
    "Country": "cn",
    "Province": "sh",
    "City": "sh",
    "Addr": "wenhui st",
    "LocLong": 11.1,
    "LocLat": 11.1,
    "RegStartTime": "2015-01-01T01:01:01+08:00",
    "RegEndTime": "2015-01-01T01:01:01+08:00",
    "StartTime": "2015-01-01T01:01:01+08:00",
    "EndTime": "2015-01-01T01:01:01+08:00",
    "CloseTime": "0001-01-01T00:00:00Z",
    "Limitation": 1000,
    "LimitationType": 1,
    "ExtraInfoJson": "",
    "UserId": 1
  }
}
```
参数说明: 

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-1,-11,-9
msg|string|否|如果错误,可以将具体错误信息返回
party|json|是|赛事信息


---

<div id="PartyClose"></div>
赛事关闭 (PartyClose)
请求样例:
```json
{"cmd":"PartyClose"
,"party_id": ""
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|--- 
cmd|string| 是| PartyClose
party_id|string|是|需要关闭的赛事id

返回样例:
```json
{"result":0
,"msg":""
} 
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确,赛事关闭成功; 其它[错误码](#ErrorCode):-1,-10,-11
msg|string|否|如果错误,可以将具体错误信息返回



---

<div id="PartyQueryGames"></div>
查询赛事中的所有比赛
请求样例:
```json
{"cmd":"PartyQueryGames"
,"party_id":""}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| PartyQueryGames
party_id|string|是| party_id



返回样例:
```json
{"result":0
  ,"msg":""
  ,"games":[
    {
        "game_id":""
        ,"game_name":"xxxx"
        ,"rmb_price":xxxx
        ,"usd_price":xxxx
    }
  ,"count": 1
]}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -19
msg|string|否|如果错误,可以将具体错误信息返回
games|json|该项赛事下的所有game数据
game_id|string|是| 比赛id
game_name|string|是| 比赛名称
rmb_price|float|是| 人民币价格
usd_price|float|是| 美元价格
count|int|该赛事下games的数量(games的json里有多少条数据)


---

<div id="PartyStateUpdate"></div>
赛事审核 (PartyStateUpdate)
请求样例:
```json
{"cmd":"PartyClose"
,"state": 1
,"party_id": ""
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|--- 
cmd|string| 是| PartyStateUpdate
party_id|string|是|需要审核的赛事id
state|int|是|赛事审核结果，1表示审核通过，2表示不通过

返回样例:
```json
{"result":0
,"msg":""
} 
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确,赛事关闭成功; 
msg|string|否|如果错误,可以将具体错误信息返回


