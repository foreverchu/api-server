#!/bin/sh
#######1.修改连接数据库配置文件
API_SERVER_PATH="/root/api-server/src"
cd  $API_SERVER_PATH/consoles/scripts
ruby update_apiserver_conf.rb
#######2.停止运行服务
pgrep  chinarun_api | xargs kill -9
#######3.重新构建表结构
cd $API_SERVER_PATH
mysql -hrds2g2rdxxt711sworh8.mysql.rds.aliyuncs.com -uchinarun -pchEtHe7h  chinarun < $API_SERVER_PATH/db/chinarun_api.sql
####初始化数据
cd $API_SERVER_PATH/consoles/dbseeds
chmod +x dbseeds
./dbseeds
#####4.api文档构建
cd  $API_SERVER_PATH/consoles/apigenerator
chmod +x apigenerator
./apigenerator
#######5.重新命名可运行文件并加入可运行权限
mv $API_SERVER_PATH/chinarun_api_`date +%Y-%m-%d`  $API_SERVER_PATH/chinarun_api
cd $API_SERVER_PATH
chmod +x chinarun_api