package controllers

import (
	"err_code"
	"fmt"
	"services/message"
	"services/party"
	"services/user"
	"strconv"

	"github.com/chinarun/utils"
)

type MessageController struct {
	BaseController
}

func (c *MessageController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MessageController) Finish() {
	defer c.BaseController.Finish()
}

func (c *MessageController) Post() {
	retJson := init_retJson()

	if err := c.IsLogin(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		c.Data["json"] = retJson
		c.ServeJson()
		return
	}

	var user = c.CurrentUser.User()
	c.Data["user"] = user

	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "MsgList":
			c.Data["json"] = c.List()
		case "ShowMsg":
			c.Data["json"] = c.Show()
		case "DeleteMsg":
			c.Data["json"] = c.Destroy()
		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}

	c.ServeJson()
}

func (c *MessageController) List() map[string]interface{} {
	retJson := init_retJson()
	user := c.CurrentUser.User()
	msgs, err := messageSrv.List(user.Id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "")
		return retJson
	}
	retJson_edit(retJson, err_code.OK, "")
	retJson["msgs"] = msgs
	return retJson
}

func (c *MessageController) Show() map[string]interface{} {
	retJson := init_retJson()

	msgIdStr, _ := c.js.Get("Id").String()
	Id, err := strconv.ParseUint(msgIdStr, 10, 64)
	if err != nil {
		utils.Logger.Error("got incorrect msgid type, msgIdStr: %s", msgIdStr)
		retJson_edit(retJson, err_code.InvalidData, "信息不存在")
		return retJson
	}
	user := c.CurrentUser.User()

	msg, err := messageSrv.GetMsg(Id, user.Id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "信息不存在")
		return retJson
	}

	var toWhom string
	switch msg.MsgType {
	case utils.MSG_ORDER_PAID:
		toWhom = strconv.FormatUint(uint64(msg.FromTableId), 64)
	case utils.MSG_ORDER_REFUND:
		toWhom = strconv.FormatUint(uint64(msg.FromTableId), 64)
	case utils.MSG_PARTY_REG:
		toWhom = partySrv.FindPartyNameById(msg.FromTableId)
	case utils.MSG_PARTY_START:
		toWhom = partySrv.FindPartyNameById(msg.FromTableId)
	case utils.MSG_PARTY_PHOTO:
		toWhom = partySrv.FindPartyNameById(msg.FromTableId)
	case utils.MSG_USER_FOLLOW:
		toWhom = userSrv.FindNameById(msg.FromTableId)
	default:
		fmt.Println("未知的msg Type")
	}

	retJson_edit(retJson, err_code.OK, "")
	retJson["message"] = msg
	retJson["toWhom"] = toWhom
	return retJson
}

func (c *MessageController) Destroy() map[string]interface{} {
	retJson := init_retJson()

	user := c.CurrentUser.User()
	msgIdInt, _ := c.js.Get("msgId").Int()
	err := messageSrv.Delete(uint64(msgIdInt), user.Id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "")
		return retJson
	}
	retJson_edit(retJson, err_code.OK, "")
	return retJson
}
