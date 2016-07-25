图片相关操作接口列表
---

接口名称|接口描述|开发情况 
---|---|---
AddPhotos|[新增图片](#AddPhotos)|[YES]
DelPhotos|[删除图片](#DelPhotos)|[YES]


---

<div id="AddPhotos"></div>
添加图片接口

请求样例:
```json
{"cmd":"AddPhotos"
  ,"rel_id":""
  ,"type": int
  ,"url":["/photo/1.png","http://www.chianrun.com/photo/party/2.png"]
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| AddPhotos
rel_id|string|是|关联id, 结合type使用: type＝1 － rel_id为赛事party_id； type＝2 － rel_id为比赛game_id；type＝3 － rel_id为赛事party_id， 其他待扩展
type|int|是|照片类型: 1 － 赛事party， 2 － 比赛game， 3 － 赛事（party）路线图片， 其他待扩展
url|string数组|是|照片url数组



返回样例:
```json
{"result":0
  ,"msg":""
  ,"num":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|如果错误,可以将具体错误信息返回
num|int|成功添加图片的数目


---
<div id="DelPhotos"></div>
删除图片接口

请求样例:
```json
{"cmd":"DelPhotos"
  ,"rel_id":""
  ,"type": int
  ,"url":["/photo/1.png","http://www.chianrun.com/photo/party/2.png"]
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| DelPhotos
rel_id|string|是|关联id, 结合type使用: type＝1 － rel_id为赛事party_id； type＝2 － rel_id为比赛game_id；type＝3 － rel_id为赛事party_id， 其他待扩展
type|int|是|照片类型: 1 － 赛事party， 2 － 比赛game， 3 － 赛事（party）路线图片， 其他待扩展
url|string 数组|是|照片url数组



返回样例:
```json
{"result":0
  ,"msg":""
  ,"num":""
}
```
参数说明:

字段名称|类型|描述
---|---|---
result|int|0,正确; 其它[错误码](#ErrorCode): -1
msg|string|如果错误,可以将具体错误信息返回
num|int|成功删除图片的数目


---


