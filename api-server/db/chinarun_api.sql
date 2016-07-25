-- MySQL dump 10.13  Distrib 5.6.24, for osx10.10 (x86_64)
--
-- Host: localhost    Database: chinarun
-- ------------------------------------------------------
-- Server version	5.6.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL,
  `phone` int(11) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` char(32) DEFAULT NULL,
  `salt` char(6) DEFAULT '',
  `avatar` varchar(256) DEFAULT NULL,
  `come_from` tinyint(4) DEFAULT NULL,
  `active` tinyint(4) DEFAULT '0',
  `last_signin_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `api_server`
--

DROP TABLE IF EXISTS `api_server`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `api_server` (
  `ip` varchar(46) NOT NULL DEFAULT '' COMMENT 'api server ip',
  `reg_time` datetime NOT NULL COMMENT '注册时间',
  PRIMARY KEY (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='api server列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `email_confirm`
--

DROP TABLE IF EXISTS `email_confirm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `email_confirm` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `token` varchar(64) NOT NULL DEFAULT '',
  `used` tinyint(4) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `confirmed_at` timestamp NULL DEFAULT NULL,
  `expired_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token_unq` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `game`
--

DROP TABLE IF EXISTS `game`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `game` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `party_id` int(11) unsigned NOT NULL,
  `name` varchar(60) NOT NULL,
  `limitation` int(11) unsigned DEFAULT NULL,
  `rmb_price` int(11) unsigned DEFAULT NULL COMMENT '人民币价格 * 100\\n',
  `usd_price` int(11) unsigned DEFAULT NULL COMMENT '美元价格*100',
  `gender_req` tinyint(3) unsigned DEFAULT '0' COMMENT '比赛性别限制：0 无限制，1 只允许男性参加，2 只允许女性参加',
  `min_age_req` tinyint(3) unsigned DEFAULT '0' COMMENT '最小年龄限制, 为0表示无限制',
  `max_age_req` tinyint(3) unsigned DEFAULT '0' COMMENT '最大年龄限制, 为0表示无限制',
  `start_time` datetime NOT NULL COMMENT '比赛开始时间',
  `end_time` datetime NOT NULL COMMENT '比赛结束时间',
  `close_time` datetime DEFAULT NULL COMMENT '比赛关闭时间，为空表示未关闭',
  PRIMARY KEY (`id`),
  KEY `start_time` (`start_time`),
  KEY `party_id` (`party_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='party（赛事）中的具体比赛';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户消息通知列表',
  `user_id` int(11) unsigned NOT NULL,
  `message_id` char(40) NOT NULL,
  `from_table` varchar(32) NOT NULL DEFAULT '' COMMENT '来自于哪个table',
  `from_table_id` int(11) unsigned NOT NULL COMMENT '来自于哪个table哪个id',
  `msg_type` tinyint(4) NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0表示未读, 1 已读 2 删除',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `message_id` (`message_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `order_no` varchar(32) NOT NULL COMMENT '订单号',
  `game_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL COMMENT 'wordpress user id',
  `submit_time` datetime NOT NULL COMMENT '订单提交时间',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间，为空表示未付费',
  `refund_time` datetime DEFAULT NULL COMMENT '退款时间，为空表示未退款',
  `cancel_time` datetime DEFAULT NULL COMMENT '取消时间，为空表示未取消',
  `price` int(11) unsigned NOT NULL COMMENT '价格*100',
  `currency_type` tinyint(3) unsigned NOT NULL COMMENT '货币类型',
  `pay_method` tinyint(3) unsigned DEFAULT NULL COMMENT '''支付方法\n支付宝，微信，银行，paypal....''',
  `pay_account` varchar(32) DEFAULT NULL COMMENT '支付账号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_no` (`order_no`),
  KEY `submit_time` (`submit_time`),
  KEY `game_id` (`game_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `party`
--

DROP TABLE IF EXISTS `party`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `user_id` int(11) unsigned NOT NULL COMMENT '创建赛事者用户id',
  `limitation` int(11) DEFAULT NULL COMMENT '赛事总报名名额限制',
  `limitation_type` tinyint(4) DEFAULT NULL COMMENT '限制类型\n每个game的报名人数限制\n总名额限制\n衣服型号类型限制\n或以上3种组合',
  `country` varchar(32) DEFAULT NULL,
  `province` varchar(32) DEFAULT NULL,
  `city` varchar(32) DEFAULT NULL,
  `addr` varchar(90) DEFAULT NULL,
  `loc_long` float DEFAULT NULL COMMENT '经度',
  `loc_lat` float DEFAULT NULL COMMENT '纬度',
  `reg_start_time` datetime NOT NULL COMMENT '报名开始时间',
  `reg_end_time` datetime NOT NULL COMMENT '报名截止时间',
  `start_time` datetime NOT NULL COMMENT '赛事开始时间',
  `end_time` datetime NOT NULL COMMENT '赛事结束时间',
  `close_time` datetime DEFAULT NULL COMMENT '赛事关闭时间，为空表示未关闭',
  `valid_state` tinyint(4) DEFAULT '0' COMMENT '审核状态',
  PRIMARY KEY (`id`),
  KEY `start_time` (`start_time`),
  KEY `uesr_id` (`user_id`),
  KEY `reg_start_time` (`reg_start_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='赛事';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `party_detail`
--

DROP TABLE IF EXISTS `party_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party_detail` (
  `party_id` int(11) unsigned NOT NULL COMMENT 'party id',
  `slogan` varchar(256) DEFAULT NULL COMMENT '精彩短评',
  `like` int(11) DEFAULT NULL COMMENT '想跑',
  `website` varchar(60) DEFAULT NULL COMMENT '官网',
  `type` varchar(32) DEFAULT NULL COMMENT '赛别',
  `price` varchar(60) DEFAULT NULL COMMENT '赛事价格',
  `introduction` text COMMENT '介绍',
  `schedule` text COMMENT '日程',
  `score` float DEFAULT NULL COMMENT '用户打分',
  `signup_male` int(11) DEFAULT NULL COMMENT '已报名（男）',
  `signup_female` int(11) DEFAULT NULL COMMENT '已报名（女）',
  PRIMARY KEY (`party_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='赛事详情';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `party_reg_meta`
--

DROP TABLE IF EXISTS `party_reg_meta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party_reg_meta` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '本来是不需要的，但因为beego的model定义没发现如何支持复合主键，只好再加个id',
  `meta_id` int(11) unsigned NOT NULL COMMENT '外键，指向reg_meta->id',
  `party_id` int(11) unsigned NOT NULL COMMENT '外键，指向party->id',
  `flags` int(11) unsigned DEFAULT '0' COMMENT '标志位，& 0x01为真表示这是必填字段，否则为选填',
  PRIMARY KEY (`id`),
  UNIQUE KEY `party_meta` (`party_id`,`meta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `party_tag_map`
--

DROP TABLE IF EXISTS `party_tag_map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party_tag_map` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `party_id` int(11) unsigned NOT NULL,
  `tag_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `party_tag_unique` (`party_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='party_tag关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `photo`
--

DROP TABLE IF EXISTS `photo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `photo` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rel_id` int(11) unsigned NOT NULL COMMENT '关联id, 结合type使用: type＝1 － rel_id为赛事party_id； type＝2 － rel_id为比赛game_id；type＝3 － rel_id为赛事party_id， 其他待扩展',
  `url` varchar(128) NOT NULL,
  `type` tinyint(4) DEFAULT '0' COMMENT '照片类型: 1 － 赛事party， 2 － 比赛game， 3 － 赛事（party）路线图片， 其他待扩展',
  PRIMARY KEY (`id`),
  KEY `rel_type` (`rel_id`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='照片';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `name` varchar(32) NOT NULL,
  `certificate_type` smallint(6) NOT NULL,
  `certificate_no` varchar(20) NOT NULL,
  `mobile` varchar(18) DEFAULT NULL,
  `email` varchar(32) DEFAULT NULL,
  `country` varchar(32) DEFAULT NULL,
  `province` varchar(32) DEFAULT NULL,
  `city` varchar(32) DEFAULT NULL,
  `address1` varchar(32) DEFAULT NULL,
  `address2` varchar(32) DEFAULT NULL,
  `zip` varchar(6) DEFAULT NULL,
  `gender` tinyint(4) unsigned DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `emergency_contact_name` varchar(32) DEFAULT NULL,
  `emergency_contact_mobile` varchar(18) DEFAULT NULL,
  `t_shirt_size` tinyint(4) DEFAULT NULL,
  `last_udpate_time` timestamp NULL DEFAULT NULL,
  `extra_info_json` text COMMENT '扩展信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `certificate_no` (`certificate_no`,`certificate_type`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='information about player';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `player_score`
--

DROP TABLE IF EXISTS `player_score`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_score` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `player_id` int(11) unsigned NOT NULL,
  `game_id` int(11) unsigned NOT NULL,
  `result` int(5) unsigned NOT NULL COMMENT '比赛所用时长,单位为秒, 最长为5位数',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `player_game_unique` (`player_id`,`game_id`),
  KEY `game_id` (`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='比赛选手成绩表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `address` varchar(128) DEFAULT NULL,
  `gender` tinyint(4) DEFAULT NULL,
  `constellation` tinyint(4) DEFAULT NULL,
  `profession` varchar(32) DEFAULT NULL,
  `about` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `reg_meta`
--

DROP TABLE IF EXISTS `reg_meta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reg_meta` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '显示名称',
  `type` tinyint(4) unsigned DEFAULT '0' COMMENT '数据类型，0：单行正文，1：多行正文，2：整数，3：浮点数........',
  `ex_data` text COMMENT '根据type来解释，type是复选框时，为选择值，用\\n分隔',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `reg_meta_data`
--

DROP TABLE IF EXISTS `reg_meta_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reg_meta_data` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '本来是不需要的，但因为beego的model定义没发现如何支持复合主键，只好再加个id',
  `party_id` int(11) unsigned NOT NULL COMMENT '外键，指向party->id',
  `reg_id` int(11) unsigned NOT NULL COMMENT '外键，指向registration->id',
  `meta_id` int(11) unsigned NOT NULL COMMENT '外键，指向reg_meta->id',
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `party_reg_meta` (`party_id`,`reg_id`,`meta_id`),
  KEY `fk_meta` (`meta_id`),
  KEY `fk_reg` (`reg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `registration`
--

DROP TABLE IF EXISTS `registration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `registration` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `player_id` int(11) unsigned NOT NULL,
  `order_id` int(11) unsigned NOT NULL,
  `game_id` int(11) unsigned NOT NULL,
  `pay_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '支付状态, 为0表示等待支付, 为1表示已支付，为2表示取消订单，为3表示订单被退款',
  PRIMARY KEY (`id`),
  UNIQUE KEY `player_game` (`player_id`,`game_id`),
  KEY `game` (`game_id`,`pay_state`),
  KEY `order` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='注册表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `smscode`
--

DROP TABLE IF EXISTS `smscode`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `smscode` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '手机号码',
  `code` varchar(8) NOT NULL DEFAULT '' COMMENT '短信验证码',
  `used_at` timestamp NULL DEFAULT NULL COMMENT '使用时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
  PRIMARY KEY (`id`),
  KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tag表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `third_party_register`
--

DROP TABLE IF EXISTS `third_party_register`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `third_party_register` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL COMMENT '1=> wechat 2=>weibo',
  `from_id` varchar(128) DEFAULT NULL COMMENT 'wechat 是openid',
  `access_token` varchar(256) DEFAULT NULL,
  `refresh_token` varchar(256) DEFAULT NULL,
  `expired_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL,
  `phone` varchar(16) DEFAULT '',
  `email` varchar(100) DEFAULT NULL,
  `password` char(32) DEFAULT NULL,
  `salt` char(6) DEFAULT '',
  `avatar` varchar(256) DEFAULT NULL,
  `come_from` tinyint(4) DEFAULT NULL,
  `active` tinyint(4) NOT NULL DEFAULT '0',
  `token` varchar(256) DEFAULT NULL,
  `last_signin_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_party`
--

DROP TABLE IF EXISTS `user_party`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_party` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户关注赛事列表',
  `user_id` int(11) unsigned DEFAULT NULL,
  `party_id` int(11) unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_party` (`user_id`,`party_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_signin`
--

DROP TABLE IF EXISTS `user_signin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_signin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT NULL,
  `source` tinyint(4) DEFAULT NULL,
  `token` varchar(256) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-10-28 17:03:48
