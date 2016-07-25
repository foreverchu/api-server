register meta(报名扩展数据)
---

接口名称|接口描述|开发情况
---|---|---
RegMetaGet|[获取注册扩展数据](#RegMetaGetAll)|[ NO ]
RegMetaGetAll|[获取所有注册扩展数据](#RegMetaGetAll)|[ NO ]
RegMetaAdd|[添加注册扩展数据](#RegMetaAdd)|[ NO ]
RegMetaEdit|[编辑注册扩展数据](#RegMetaEdit)|[ NO ]
RegMetaDel|[删除注册扩展数据](#RegMetaEdit)|[ NO ]
RegMetaBindParty|[给赛事增加扩展数据绑定](#RegMetaBindParty)|[ NO ]
RegMetaUnBindParty|[给赛事去除扩展数据绑定](#RegMetaBindParty)|[ NO ]
RegMetaPartyGetAllMetas|[获得赛事所有扩展数据](#RegMetaPartyGetAllMetas)|[ NO ]
RegMetaSetRegMetaData|[设置报名的扩展数据值](#RegMetaSetRegMetaData)|[ NO ]
RegMetaGetRegMetaData|[获取报名的扩展数据值](#RegMetaSetRegMetaData)|[ NO ]



---

<div id="RegMetaGet"></div>
获取注册扩展数据
请求样例:
```json
{"cmd":"RegMetaGet"
,"meta_id":xx}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaGet
meta_id|int|是| meta id



返回样例:
```json
{"result":0
  ,"msg":""
  ,"meta":{
      ,"Id":xxx
      ,"Name":""
      ,"Type":xx
      ,"ExData":""
  }
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -42
msg|string|否|如果错误,可以将具体错误信息返回
Id|int|是|meta id
Name|string|是|扩展数据名称
Type|string|是|数据类型, 请参看[注册扩展数据类型说明](#RegMetaType)
ExData|string|是|扩展数据，根据Type的值会有不同解释


---


<div id="RegMetaGetAll"></div>
获取所有注册扩展数据
请求样例:
```json
{"cmd":"RegMetaGetAll"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaGetAll



返回样例:
```json
{"result":0
  ,"msg":""
  ,"metas":[
    {
      ,"Id":xxx
      ,"Name":""
      ,"Type":xx
      ,"ExData":""
    }
  ]
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|否|如果错误,可以将具体错误信息返回
Id|int|是|meta id
Name|string|是|扩展数据名称
Type|string|是|数据类型, 请参看[注册扩展数据类型说明](#RegMetaType)
ExData|string|是|扩展数据，根据Type的值会有不同解释


---


<div id="RegMetaAdd"></div>
添加注册扩展数据
请求样例:
```json
{"cmd":"RegMetaAdd"
  ,"Name":""
  ,"Type":xx
  ,"ExData":""
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaAdd
Name|string|是|扩展数据名称
Type|string|是|数据类型, 请参看[注册扩展数据类型说明](#RegMetaType)
ExData|string|是|扩展数据，根据Type的值会有不同解释



返回样例:
```json
{"result":0
  ,"msg":""
  ,"meta_id":xxx
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|否|如果错误,可以将具体错误信息返回
meta_id|int|否|添加成功返回meta id


---


<div id="RegMetaEdit"></div>
编辑注册扩展数据

请求样例: 
```json
{"cmd":"RegMetaEdit"
  ,"meta_id":xxx
  ,"Name":""
  ,"Type":xx
  ,"ExData":""
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaEdit
meta_id|int|是|meta id
Name|string|否|扩展数据名称
Type|string|是|数据类型, 请参看[注册扩展数据类型说明](#RegMetaType)
ExData|string|是|扩展数据，根据Type的值会有不同解释



返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -42
msg|string|否|如果错误,可以将具体错误信息返回



---


<div id="RegMetaDel"></div>
删除注册扩展数据

请求样例:
```json
{"cmd":"RegMetaDel"
  ,"meta_id":xxx
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaDel
meta_id|int|是|meta id



返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -42
msg|string|否|如果错误,可以将具体错误信息返回


---


<div id="RegMetaBindParty"></div>
给赛事增加扩展数据绑定
请求样例:
```json
{"cmd":"RegMetaBindParty"
  ,"party_id":""
  ,"meta_id":xxx
  ,"flags":xxx
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaBindParty
party_id|string|是|party id
meta_id|int|是|meta id
flags|int|是|标志，与0x01与为非0，表示此字段为必填



返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -9
msg|string|否|如果错误,可以将具体错误信息返回


---


<div id="RegMetaUnBindParty"></div>
给赛事去除扩展数据绑定

请求样例:
```json
{"cmd":"RegMetaUnBindParty"
  ,"party_id":""
  ,"meta_id":xxx
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaUnBindParty
party_id|string|是|party id
meta_id|int|是|meta id



返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -9
msg|string|否|如果错误,可以将具体错误信息返回


---


<div id="RegMetaPartyGetAllMetas"></div>
获得赛事所有扩展数据

请求样例:
```json
{"cmd":"RegMetaPartyGetAllMetas"
  ,"party_id":""
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaPartyGetAllMetas
party_id|string|是|party id



返回样例:
```json
{"result":0
  ,"msg":""
  ,"metas":[
    {
      ,"Id":xxx
      ,"Name":""
      ,"Type":xx
      ,"ExData":""
    }
  ]
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -9
msg|string|否|如果错误,可以将具体错误信息返回
Id|int|是|meta id
Name|string|是|扩展数据名称
Type|string|是|数据类型, 请参看[注册扩展数据类型说明](#RegMetaType)
ExData|string|是|扩展数据，根据Type的值会有不同解释



---

<div id="RegMetaSetRegMetaData"></div>
给赛事去除扩展数据绑定

请求样例:
```json
{"cmd":"RegMetaSetRegMetaData"
  ,"party_id":""
  ,"meta_id":xxx
  ,"value":""
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaSetRegMetaData
party_id|string|是|party id
meta_id|int|是|meta id
value|string|是|值



返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1, -9, -13
msg|string|否|如果错误,可以将具体错误信息返回


---

<div id="RegMetaGetRegMetaData"></div>
给赛事去除扩展数据绑定

请求样例:
```json
{"cmd":"RegMetaGetRegMetaData"
  ,"party_id":""
  ,"meta_id":xxx
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RegMetaGetRegMetaData
party_id|string|是|party id
meta_id|int|是|meta id
value|string|是|值



返回样例:
```json
{"result":0
  ,"msg":""
  ,"value":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|否|如果错误,可以将具体错误信息返回
value|string|否|值， 函数调用正确才会返回



---
