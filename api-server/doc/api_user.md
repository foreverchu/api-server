user操作接口列表：
----

接口列表：

接口名称|接口描述|开发情况
---|---|---
UserRegister|[用户注册](#UserRegister)|[YES]
UserWechatSignin|[微信登录或注册](#UserWechatSignin)|[YES]
UserRegisterSMSCode|[短信验证码](#UserRegisterSMSCode)|(YES)
UserEmailConfirm|[邮箱激活](#UserEmailConfirm)|[YES]
UserSignin|[用户登录](#UserSignin)|[YES]
UserResetPassword|[重置密码](#UserResetPassword)|[YES]
UserForgetPassword|[忘记密码](#UserForgetPassword)|[NO]
UserInfoQuery|[用户信息查询](#UserInfoQuery)|[ YES ]
EditPersonalInfo|[用户信息编辑](#EditPersonalInfo)|[ YES ]

---

<div id="UserRegister"></div>

用户注册(UserRegister)

请求样例:
```json
email:
{
        "cmd":"UserRegister",
        "come_from":"email"
        "email":"",
        "password":"",
        "password_confirm":"",
        "confirm_host":"http://www.wanbisai.com/email/confirm/token/"
}
phone:
{
        "cmd":"UserRegister",
        "come_from":"phone"
        "phone":"",
        "password":"",
        "password_confirm":"",
        "sms_code":"",
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserRegister
email|string|是|邮箱
password|string|是|登录密码
password_confirm|string|是|登录密码
come_from|string|是|"email"
confirm_host|string|是|用户邮箱里收到的地址


字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserRegister
phone|string|是|手机号码
password|string|是|登录密码
password_confirm|string|是|登录密码
sms_code|string|是|短信验证码
come_from|string|是|"phone"


返回样例:
```json
{
    "result":0,
    "msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-15
msg|string|否|如果错误,可以将具体错误信息返回


---
<div id="UserWechatSignin"></div>

微信用户注册 (UserWechatSignin)

请求样例:


---
<div id="UserRegisterSMSCode"></div>

短信验证码 (UserRegisterSMSCode)

请求样例:
{
    "cmd":"UserRegisterSMSCode",
    "phone":"18101635160"
}

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserRegisterSMSCode
phone |string|是|手机号码


返回样例:
```json
{
    "result":0,
    "msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-15
msg|string|否|如果错误,可以将具体错误信息返回

---


<div id="UserSignin"></div>

用户登录 (UserSignin)

请求样例:</font> 
```json
{"cmd":"UserSign"
   ,"signin_by":""
   ,"password":""
   ,"auto_signin":true
}

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserSiginin
signin_by|string|是|用户名或者手机
password|string|是|登录密码
auto_signin |string|否| 是否自动登录,可以传非0非空字符串


返回样例:
```json
{
  "msg": "",
  "result": 0,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NDU1NzU4MzAsInNyYyI6MSwidWlkIjoxfQ.lBScS9YhCUMgmXZb-4YM7uy9kjeQHo8-XUy9z24t_aQ",
  "user": {
    "Id": 1,
    "Name": "Adriana Crona",
    "Avatar": "",
    "ComeFrom": 1,
    "Active": 1,
    "LastSigninAt": "0001-01-01T00:00:00Z",
    "CreatedAt": "2015-09-17T16:07:03+08:00"
  }
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-13
msg|string|否|如果错误,可以将具体错误信息返回


---

<div id="EditPersonalInfo"></div>

用户修改个人资料 (EditPersonalInfo)

请求样例:</font> 
```json
{"cmd":"EditPersonalInfo"
   ,"name":""
   ,"avatar":""
   ,"address":""
   ,"gender":1
   ,"constellation":1
   ,"profession":"da"
   ,"about":""
   }

```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是|EditPersonalInfo
name|string|否|用户昵称
avatar|string|否|用户头像
address|string|否|地址
gender|int|否|性别。1表示男，2表示女 
constellation|int|否|星座
profession|string|否| 专业
about |string|否| 介绍

返回样例:
```json
{
  "msg": "",
  "result": 0,
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-13
msg|string|否|如果错误,可以将具体错误信息返回


---

<div id="UserEmailConfirm"></div>

邮箱注册验证 (UserEmailConfirm)

请求样例:</font> 
```json
{
    "cmd":"UserEmailConfirm",
    "token":"JehggoPt2pXbc2tMMwddqZvigJkaN1RUBYKMampvHiGtLm5vXIwWXW9cIDIFIJDx"
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserEmailConfirm
token |string|是|邮箱验证token


返回样例:
```json
{
    "result":0,
    "msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-15
msg|string|否|如果错误,可以将具体错误信息返回

---
<div id="UserSignin"></div>
用户登录 (UserSignin)


请求样例:</font> 
```json
{
    "cmd": "UserSignin",
    "signin_by": "qiangaoming@chinarun.com",
    "password": "123456",
    "auto_signin": true
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserSignin
signin_by |string|是|user name
password |string|是|user password
auto_signin |bool|否|true表示自动登录


返回样例:
```json
{
    "result":0,
    "msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-13
msg|string|否|如果错误,可以将具体错误信息返回

---

<div id="UserResetPassword"></div>
用户重置密码 (UserResetPassword)

请求样例:</font> 
```json
{
    "cmd": "UserResetPassword",
    "old_password": "123456",
    "new_password": "654321",
    "new_password_confirm": "654321"
}
```

参数说明:

字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserResetPassword
old_password |string|是|旧密码
new_password |string|是|新密码
new_password_confirm |string|是|确认新密码


返回样例:
```json
{
    "result":0,
    "msg":""
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-13
msg|string|否|如果错误,可以将具体错误信息返回

---


<div id="UserInfoQuery"></div>
用户信息查询 (UserInfoQuery)


请求样例:</font> 
```json
{
    "cmd": "UserInfoQuery",
    "user_id": "1"
}
```

参数说明:


字段名称|类型|必须|描述
---|---|---|---
cmd|string| 是| UserInfoQuery
user_id |string|是|用户ID



返回样例:
```json
{
    "result":0,
    "msg":"",
    "user":{
      "profile": {
         "Id": 0,
         "UserId": 0,
         "Address": "",
         "Gender": 0,
         "Constellation": 0,
         "Profession": "",
         "About": "",
         "CreatedAt": "0001-01-01T00:00:00Z",
         "UpdatedAt": "0001-01-01T00:00:00Z"
         },
      "userId": "1",
      "user":{
         "Id":1,           
         "Name":"xx",
         "Avatar":"xxxx",
         "ComeFrom":1,
         "Active":1,
         "LastSigninAt":"2014-11-11 12-12-12", 
         "CreatedAt":"2015-10-20 13:29:24"    
      }
    }
}
```
参数说明:

字段名称|类型|必须|描述
---|---|---|---
result|int|是|0,正确; 其它[错误码](#ErrorCode):-13
msg|string|否|如果错误,可以将具体错误信息返回
userId|int|否|如果result为0， 返回userid
Id|int|否|如果result为0， 返回Id
Name|string|否|如果result为0， 返回Id
Avatar|string|否|如果result为0， 返回avatar
ComeFrom|int|否|如果result为0， 返回comefrom
Active|int|否|如果result为0， 返回active
LastSigninAt|string|否|如果result为0， 返回LastSigninAt
CreatedAt|string|否|如果result为0， 返回CreateAt
profile|string|是|用户拓展信息

---

