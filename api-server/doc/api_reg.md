register（报名）操作接口列表：
---

接口名称|接口描述|开发情况
---|---|---
RegGetQueueToken|[获取报名排队token](#RegGetQueueToken)|[ NO ]
RegQueryQueueTokenState|[查询排队token状态](#RegQueryQueueTokenState)|[ NO ]
RegRegistration|[报名](#RegRegistration)|[ NO ] 没有文档说明
RegUpdate|[更新报名](#RegUpdate)|[ NO ]
RegAskRefund|[申请退款](#RegAskRefund)|[ NO ]
RegQueryUserOrders|[查询用户报名订单](#RegQueryUserOrders)|[ NO ]
RegQueryOrder|[查询订单下所有报名](#RegQueryOrder)|[ NO ]
RegSetGameLimitation|[设置赛事名额限制](#RegSetGameLimitation)|[ YES ]
RegSetAllGamesLimitation|[设置所有赛事名额限制](#RegSetAllGamesLimitation)|[ YES ]
RegGetGameLimitationInfo|[获取赛事名额限制及已报名人数](#RegGetGameLimitationInfo)|[ YES ]
RegGetAllGamesBalanceInfo|[获取所有正在报名比赛余下的名额](#RegGetAllGamesBalanceInfo)|[ YES ]
QueryRegStateByCert|[根据证件号码查询赛事报名情况](#QueryRegStateByCert)|[ NO ]
RegQueryPartyRegState|[查询赛事报名情况](#RegQueryPartyRegState)|[ NO ]没有文档说明


---
<div id="QueryRegStateByCert"></div>
根据证件号码查询赛事报名情况(QueryRegStateByCert)
请求样例:
```json
{"cmd":"QueryRegStateByCert"
,"party_id":"xxx"
,"certificate_type": x
,"certificate_no": "xxx"
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| QueryRegStateByCert
party_id|string|是|party id
certificate_type|uint8|是|证件类型
certificate_no|string|是|证件号码

返回样例: 
```json
{"result":0
,"msg":""
,"order_info":{
"order_no": "xxx"
,"game_name": "xxx"
,"pay_status": x
,"currency_type": x
,"price": float
}
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode)
msg|string|否|如果错误,可以将具体错误信息返回
order_info|json|是|包含付款比赛信息的json
order_no|string|否|订单号
game_name|string|否|比赛名称
pay_status|uint8|否|付款状态(0表示等待支付,1表示已支付,2表示已取消,3表示已退款)
currency_type|uint8|否|付款币种， 1:人民币, 2：美元
price|float|否|付款金额

---

<div id="RegGetAllGamesBalanceInfo"></div>
获取所有正在报名比赛余下的名额
求样例:
```json
{"cmd":"RegGetAllGamesBalanceInfo"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegGetAllGamesBalanceInfo



返回样例:
```json
{"result":0
  ,"msg":""
  ,"balances":[
    {
      "game_id":""
      ,"limitation":xxxx
      ,"balance":xxxx
    }
  ]
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确,比赛已经删除; 其它[错误码](#ErrorCode): -1, -10, -11
msg|string|否|如果错误,可以将具体错误信息返回
game_id|string|是|game id
limitation|int|比赛名额限制
balance|int|比赛剩余名额



---


<div id="RegGetQueueToken"></div>
填写选手信息获取token排队(RegGetQueueToken)
请求样例:
```json
{"cmd":"RegGetQueueToken"
,"game_id":""
,"user_id":""
,"players":[
    {
        "name":"xxxx"
        , "certificate_type": x
        , "certificate_no": "xx"
        , "mobile":"xxxx"
        , "email":"xxx"
        , "country":"xxx"
        , "province":"xxx"
        , "city":"xxx"
        , "address1":"xxx"
        , "address2":"xxx"
        , "zip":"xxx"
        , "gender":x
        , "birth_date":"xxx"
        , "emergency_contact_name":"xxx"
        , "emergency_contact_mobile":"xxx"
        , "t_shirt_size":x
        , "extra_info_json":"xxx"
    }
]}
```

参数说明:</font> 

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegInfoGetToken
game_id|string|是|要报名的比赛id
user_id|string|是|wordpress用户id
players|json|是|报名选手信息数组,具体字段请参加下面的 [选手信息字段说明](#PlayerFields)


---

<div id="RegQueryQueueTokenState"></div>
查询排队token状态(RegQueryQueueTokenState)
请求样例:
```json
{"cmd":"RegQueryQueueTokenState"
,"token": "xxx"}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegQueryQueueTokenState
token|string|是| 排队token id


---


返回样例:
```json
{"result":0
,"msg":""
,"queue_state":x
,"order_no":""
,"count": xxx
,"query_interval": 5}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1,-12
queue_state|int|0：需要继续等待,1：排队成功,请根据给定order_no支付, 2：报名名额已满
order_no|string|报名成功后的订单号，要根据此id支付
count|int|现在排队第几个
query_interval|int|秒数，提示重新刷新时间


---

<div id="RegUpdate"></div>
报名更新(RegUpdate)[暂不支持]
请求样例:
```json
{"cmd":"RegUpdate"
,"reg_id": "xxx"
,"name":"xxxx"
, "certificate_type":x
, "mobile":"xxxx"
, "email":"xxx"
, "country":"xxx"
, "province":"xxx"
, "city":"xxx"
, "address1":"xxx"
, "address2":"xxx"
, "zip":"xxx"
, "gender":"xxx"
, "birth_date":"xxx"
, "emergency_contact_name":"xxx"
, "emergency_contact_mobile":"xxx"
, "t_shirt_size":"xxx"
, "extra_info_json":"xxx"
]}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegGetQueueToken
reg_id|string|是| 选手报名id
name|string| 是| 选手姓名
certificate_type|string|是| 身份证件类型
certificate_no|string|是| 报名token id
mobile|string|是|选手手机号码
email|string|否|选手email地址
country|string|是|选手国家和地区
province|string|否|选手省份
city|string|否|选手城市
address1|string|否|选手街道地址1
address2|string|否|选手街道地址2
zip|string|否|邮编
gender|string|是|性别
birth_date|string|是|生日
emergency_contact_name|string|是|紧急联系人姓名
emergency_contact_mobile|string|是|紧急联系人手机
t_shirt_size|string|是|T恤大小 (XL, M, S...)
extra_info_json|string|否|扩展信息


返回样例:
```json
{"result":0}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,更新成功; 其它[错误码](#ErrorCode): -1,-11,-13



---
<div id="RegAskRefund"></div>
申请退款
求样例:
```json
{"cmd":"RegAskRefund"
,"order_no": "xxx"}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegAskRefund
order_no|string|是| 订单 id



返回样例:
```json
{"result":0}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1,-12,-14,-17


---

<div id="RegQueryUserOrders"></div>
查询用户报名订单
求样样例:
```json
{"cmd":"RegQueryUserOrders"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegAskRefund




返回样例:
```json
{"result":0
,"orders":[
    {
        "order_no":"xxxx"
        ,"party_id":"xxxx"
        ,"party_name":"xxxx"
        ,"apply_time":"2015-07-14T10:08:15+00:00"
        ,"pay_time":"2015-07-14T10:08:15+00:00"
    }
]}
```
参数说明: 

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
orders|数组| 订单数组,具体字段参看[订单字段说明](#OrderFields)


---

<div id="RegQueryOrder"></div>
查询订单下所有报名

请求样例:
```json
{"cmd":"RegQueryOrder"
,"order_no": "xxx"}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegQueryOrder
order_no|string|是| 订单 id



返回样例:
```json
{"result":0
,"registrations":[
    {
        ,"game_id":"xxxx"
        ,"game_name":"xxxx"
        ,"player_id":"xxxx"
        ,"player_name":"xxxx"
        ,""
    }
]
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1,-12,-14,-17
registrations|数组|比赛注册信息数组,具体字段参看[注册字段说明](#RegistrationFields)


---



<div id="RegSetAllGamesLimitation"></div>
设置赛事名额限制(RegSetAllGamesLimitation)
请求样例:
```json
{"cmd":"RegSetAllGamesLimitation"
  ,"balances":[
    {
      "game_id":""
      ,"limitation":xxxx
      ,"balance":xxxx
    }
  ]
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegSetAllGamesLimitation
game_id|string|是|比赛id
limitation|int|是|比赛报名限额
balance|int|否|比赛剩余名额（若无此参数，则等于limitation)


---


返回样例:
```json
{"result":0
,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1,-10, -13
msg|string|错误详细说明




---

<div id="RegSetGameLimitation"></div>
设置赛事名额限制(RegSetGameLimitation)
请求样例:
```json
{"cmd":"RegSetGameLimitation"
,"game_id": "xxx"
,"limitation": xxx
,"balance": xxx
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegSetGameLimitation
game_id|string|是|比赛id
limitation|int|是|比赛报名限额
balance|int|否|比赛剩余名额（若无此参数，则等于limitation)


---


返回样例:
```json
{"result":0
,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1,-10, -13
msg|string|错误详细说明


---


<div id="RegGetGameLimitationInfo"></div>
获取赛事名额限制及已报名人数(RegGetGameLimitationInfo)
请求样例:
```json
{"cmd":"RegGetGameLimitationInfo"
,"game_id": "xxx"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegGetGameLimitationInfo
game_id|string|是|比赛id


---


返回样例:
```json
{"result":0
,"msg":""
,"limitation": xxx
,"balance": xxx
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1,-10
msg|string|错误详细说明
limitation|int|比赛报名限额
balance|int|比赛剩余名额


---




<div id="RegGetQueueToken"></div>
填写选手信息获取token排队(RegGetQueueToken)
请求样例:
```json
{"cmd":"RegGetQueueToken"
,"game_id":""
,"user_id":""
,"players":[
    {
        "name":"xxxx"
        , "certificate_type": x
        , "certificate_no": "xx"
        , "mobile":"xxxx"
        , "email":"xxx"
        , "country":"xxx"
        , "province":"xxx"
        , "city":"xxx"
        , "address1":"xxx"
        , "address2":"xxx"
        , "zip":"xxx"
        , "gender":x
        , "birth_date":"xxx"
        , "emergency_contact_name":"xxx"
        , "emergency_contact_mobile":"xxx"
        , "t_shirt_size":x
        , "extra_info_json":"xxx"
    }
]}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegInfoGetToken
game_id|string|是|要报名的比赛id
user_id|string|是|wordpress用户id
players|json|是|报名选手信息数组,具体字段请参加下面的 [选手信息字段说明](#PlayerFields)


---
