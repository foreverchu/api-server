server（服务器）操作接口列表：
---

接口名称|接口描述|开发情况
---|---|---
AddWhiteListIP|[添加许可白名单ip](#AddWhiteListIP)|[ NO ]
RemoveWhiteListIP|[删除许可白名单IP](#RemoveWhiteListIP)|[ NO ]


---


<div id="AddWhiteListIP"></div>
添加许可白名单ip
请求样例:
```json
{"cmd":"AddWhiteListIP"
,"ip":"1.1.1.1"}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| AddWhiteListIP
ip|string|否| 若没有该参数，则将请求方ip放入信任列表


返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|否|如果错误,可以将具体错误信息返回


---


<div id="RemoveWhiteListIP"></div>
删除许可白名单ip
请求样例:
```json
{"cmd":"RemoveWhiteListIP"
,"ip":"1.1.1.1"}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| RemoveWhiteListIP
ip|string|否| 若没有该参数，则将请求方ip从信任列表删除



返回样例:
```json
{"result":0
  ,"msg":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|否|如果错误,可以将具体错误信息返回


---
