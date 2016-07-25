team（队伍）操作接口列表：
---

接口名称|接口描述|开发情况
---|---|---
TeamCreate|[组队](#TeamCreate)|[ NO ]
TeamQueryPlayer|[列出选手在某一赛事中所属队伍](#TeamQueryPlayer)|[ NO ]
TeamQueryInfo|[查询队伍资料](#TeamQueryInfo)|[ NO ]


---


<div id="TeamCreate"></div>
组队
请求样例:
```json
{"cmd":"TeamCreate"
,"team_name":"xxxx"
,"players_id_no": [
    "xxxx"
    ,"xxxx"....
]}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| TeamCreate
team_name|string|是| 组队名称
players_id_no|数组|是| 组队选手的证件id数组



返回样例:
```json
{"result":0
,"team_id":""
,"msg":"xxxxx"}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -21, -22
team_id|string| 创建成功时返回team_id
msg|string|当result为-21时,给出更详细错误信息


---



<div id="TeamQueryPlayer"></div>
列出选手在某一赛事中所属队伍
请求样例:
```json
{"cmd":"TeamQueryPlayer"
,"party_id":"xxxx"
,"player_id":"xxxx"
,"player_id_no":"xxxx"}
```

###<font color="red"> 参数说明:</font> 

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| TeamQueryPlayer
party_id|string|是| 赛事id
player_id|string|否| 必须和player_id_no有一个,当两个都存在时,已player_id为准
player_id_no|string|否| 选手证件号码



返回样例:
```json
{"result":0
  ,"msg":""
  ,"team_id":"xxxx"
  ,"team_name":"xxxxx"}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -21, -23
msg|string|否|如果错误,可以将具体错误信息返回
team_id|string| team_id
team_name|string| 队伍名称


---


<div id="TeamQueryInfo"></div>
查询队伍资料
请求样例:
```json
{"cmd":"TeamQueryInfo"
,"team_id":"xxxx"
,"party_id"
,"team_name":"xxxx"}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| TeamQueryInfo
team_id|string|否| 队伍id,当有队伍id时下面两个字段不需要
party_id|string|否| 赛事id,没有队伍id时,必须要有
team_name|string|否| 队伍名称,没有队伍id时,必须要有



返回样例: 
```json
{"result":0
  ,"msg":""
  ,"team_id":"xxxxx"
  ,"team_name":"xxxxx"
  ,"players":[
    {
        "player_id":"xxxx"
        "player_name":"xxxx"
    }
]
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -9, -24, -25
msg|string|否|如果错误,可以将具体错误信息返回
team_id|string| team_id
team_name|string| 队伍名称
players|数组| 选手数据


---