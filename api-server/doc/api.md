ChinaRun API 服务器接口设计文档

概述:
----
用户程序与平台交互使用WebService, 通过JSON格式来进行数据交互。
下面就相关接口进行描述。

说明:
----

 * 标注为`TODO`的接口是暂时没有实现的接口.
 * 通过`HTTPS`来保障数据交互的安全性.
 * 所有api都可能返回错误码-1000及-1001,可通过msg字段获取更多错误信息(IMOK除外)
 * 运行项目前，必须先设置环境变量CHINARUN_API_SERVER_MODE 为production、testing或者development

 
接口地址:
----
请求接口地址：

*   party相关操作[https://domain/api/party](http://apitest.wanbisai.com/static/api_party.html)
*   game相关操作 [https://domain/api/game](http://apitest.wanbisai.com/static/api_game.html)
*   user相关操作[https://domain/api/user](http://apitest.wanbisai.com/static/api_user.html)    
*   register(报名)相关操作[https://domain/api/reg](http://apitest.wanbisai.com/static/api_reg.html)
*   register meta(报名扩展数据)相关操作[https://domain/api/reg/meta](http://apitest.wanbisai.com/static/api_regmeta.html)
*   team相关操作[https://domain/api/team](http://apitest.wanbisai.com/static/api_team.html)
*   server相关操作[https://domain/api/server](http://apitest.wanbisai.com/static/api_server.html)
*   健康状态检查 https://domain/api/IMOK

