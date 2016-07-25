game（比赛）操作接口列表：
---

接口名称|接口描述|开发情况
---|---|---
GameQueryRegInfo|[查询比赛报名统计资料](#GameQueryRegInfo)|[ NO ]
GameInputScore|[输入比赛选手成绩](#GameInputScore)|[ NO ]
GameCreate|[创建赛事中比赛](#GameCreate)|[ YES ]
GameEdit|[编辑赛事中比赛](#GameEdit)|[ YES ]
GameClose|[关闭比赛](#GameClose)|[ YES ]
GameQuery|[查询比赛](#GameQuery)|[ YES ]
GameList|[查询比赛列表](#GameList)|[ YES ]


---

<div id="GameQueryRegInfo"></div>
查询比赛报名统计资料
请求样例: 
```
{"cmd":"GameQueryRegInfo"
,"game_id":"xxxx"}
```

参数说明: 

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameQueryRegInfo
game_id|string|是| 比赛id


返回样例: 
```json
{"result":0
  ,"apply":xxxx
  ,"payed":xxxxx
  ,"msg":""}
```
参数说明: 

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -10, -11
apply|string| 报名人数
payed|string| 已付款人数
msg|string|否|如果错误,可以将具体错误信息返回


---
<div id="GameInputScore"></div>
输入比赛选手成绩
请求样例:
```json
{"cmd":"GameInputScore"
,"game_id":"xxxx"
,"player_id:"xxxx"
,"result":3600}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameInputScore
game_id|string|是| 比赛id
player_id|string|是| 选手id
result|int|是| 选手成绩(单位为秒, 指完成跑步所用时间)

返回样例:
```json
{"result":0
  ,"msg":""}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -10, -13(无效的成绩数据）, -23
msg|string|否|如果错误,可以将具体错误信息返回



---

<div id="GameCreate"></div>
创建赛事中比赛
请求样例:
```json
{"cmd":"GameCreate"
,"party_id":"xxxx"
,"name":"xxxx"
,"limitation":xxxx
,"start_time":""
,"end_time":""
,"rmb_price":xxxx
,"usd_price":xxxx
,"gender_req":xxxx
,"min_age_req":xxxx
,"max_age_req":xxxx
,"user_id":""
}
```
  
  
参数说明: 
  
字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameCreate
party_id|string|是| 赛事id 
name|string|是| 比赛名称 
limitation|int|是| 比赛名额限制 
start_time|string|是|比赛开始时间 
end_time|string|是|比赛结束时间 
rmb_price|float|是| 人民币价格 
usd_price|float|是| 美元价格 
gender_req|uint8|否| 比赛性别限制：0 无限制，1 只允许男性参加，2 只允许女性参加 
min_age_req|uint8|否| 比赛性别限制：最小年龄限制, 为0表示无限制 
max_age_req|uint8|否| 比赛性别限制：最大年龄限制, 为0表示无限制 
user_id|string|是|wordpress用户id 

  
返回样例: 
```json
{"result":0
  ,"msg":""
  ,"game_id":""}
```
参数说明: 
 
字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -11, -13, -26
msg|string|否|如果错误,可以将具体错误信息返回
game_id|string|若创建成功,result=0时，返回刚创建成功的game id

  
  
---

<div id="GameEdit"></div>
编辑比赛
请求样例:
```json
{"cmd":"GameEdit"
,"game_id":"xxxx"
,"name":"xxxx"
,"limitation":xxxx
,"start_time":""
,"end_time":""
,"rmb_price":xxxx
,"usd_price":xxxx
,"gender_req":xxxx
,"min_age_req":xxxx
,"max_age_req":xxxx
,"user_id":""
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameEdit
game_id|string|是| 比赛id
name|string|否| 比赛名称
limitation|int|否| 比赛名额限制
start_time|string|否|比赛开始时间
end_time|string|否|比赛结束时间
rmb_price|float|否| 人民币价格
usd_price|float|否| 美元价格
gender_req|uint8|否| 比赛性别限制：0 无限制，1 只允许男性参加，2 只允许女性参加
min_age_req|uint8|否| 比赛性别限制：最小年龄限制, 为0表示无限制
max_age_req|uint8|否| 比赛性别限制：最大年龄限制, 为0表示无限制
user_id|string|否|wordpress用户id


返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明: 

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -10, -11, -13, -26
msg|string|否|如果错误,可以将具体错误信息返回



---

<div id="GameClose"></div>
关闭比赛
请求样例: 
```json
{"cmd":"GameClose"
,"game_id":""
,"user_id":""
}
```

参数说明: 

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameClose
game_id|string|是| 比赛id
user_id|string|是|wordpress用户id



返回样例: 
```json
{"result":0
  ,"msg":""}
```
参数说明: 

字段名称|类型|描述
---|---|---
result|int|0,正确,比赛已经删除; 其它[错误码](#ErrorCode): -1, -10, -11
msg|string|否|如果错误,可以将具体错误信息返回


---

<div id="GameQuery"></div>
查询比赛
请求样例:
```json
{"cmd":"GameQuery"
,"game_id":"xxxx"
}
```


参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameQuery
game_id|string|是| 比赛id


返回样例:
```json
{"result":0
  ,"msg":""
  ,"party_id":""
  ,"game": {
    "Id": "1",
    "Name": "半程马拉松",
    "Limitation": 20000,
    "RmbPrice": 19900,
    "UsdPrice": 8000,
    "GenderReq": 0,
    "MinAgeReq": 0,
    "MaxAgeReq": 0,
    "StartTime": "2034-08-24T00:00:00+08:00",
    "EndTime": "2034-08-26T23:59:59+08:00",
    "CloseTime": "0001-01-01T00:00:00Z"
  }
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -11, -13, -26
msg|string|否|如果错误,可以将具体错误信息返回
party_id|string|比赛所从属的赛事id
Id|string|是| 比赛id
Name|string|是| 比赛名称
Limitation|int|是| 比赛名额限制
StartTime|string|是|比赛开始时间
EndTime|string|是|比赛结束时间
CloseTime|string|是|比赛关闭时间，若为"0001-01-01T00:00:00Z"表示比赛没有关闭
RmbPrice|float|是| 人民币价格
UsdPrice|float|是| 美元价格
GenderReq|uint8|否| 比赛性别限制：0 无限制，1 只允许男性参加，2 只允许女性参加
MinAgeReq|uint8|否| 比赛最小年龄限制, 为0表示无限制
MaxAgeReq|uint8|否| 比赛最大年龄限制, 为0表示无限制


---


<div id="GameList"></div>
查询比赛列表
请求样例:
```json
{"cmd":"GameList"
,"party_id":"xxxx"
}
```


参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| GameList
party_id|string|是| 赛事id


返回样例:
```json
{"result":0
  ,"msg":""
  ,"count":1
  ,"gameList": [{
    "Id": "1",
    "Name": "半程马拉松",
    "Limitation": 20000,
    "RmbPrice": 19900,
    "UsdPrice": 8000,
    "GenderReq": 0,
    "MinAgeReq": 0,
    "MaxAgeReq": 0,
    "StartTime": "2034-08-24T00:00:00+08:00",
    "EndTime": "2034-08-26T23:59:59+08:00",
    "CloseTime": "0001-01-01T00:00:00Z"
  }]
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -11, -13, -26
msg|string|否|如果错误,可以将具体错误信息返回
count|int|比赛列表数目
Id|string|是| 比赛id
Name|string|是| 比赛名称
Limitation|int|是| 比赛名额限制
StartTime|string|是|比赛开始时间
EndTime|string|是|比赛结束时间
CloseTime|string|是|比赛关闭时间，若为"0001-01-01T00:00:00Z"表示比赛没有关闭
RmbPrice|float|是| 人民币价格
UsdPrice|float|是| 美元价格
GenderReq|uint8|否| 比赛性别限制：0 无限制，1 只允许男性参加，2 只允许女性参加
MinAgeReq|uint8|否| 比赛最小年龄限制, 为0表示无限制
MaxAgeReq|uint8|否| 比赛最大年龄限制, 为0表示无限制


---



