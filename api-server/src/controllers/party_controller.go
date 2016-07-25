package controllers

import "err_code"

type PartyController struct {
	BaseController
}

func (c *PartyController) Prepare() {
	c.BaseController.Prepare()
}

func (c *PartyController) Finish() {
	defer c.BaseController.Finish()
}

func (c *PartyController) Post() {
	retJson := init_retJson()

	ctx := c.Ctx
	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "PartyEdit":
			c.Data["json"] = c.PartyEdit()
		case "PartyClose":
			c.Data["json"] = c.PartyClose()
		case "PartyCreate": //OK
			c.Data["json"] = c.PartyCreate()
		case "PartyQueryRegResult":
			c.Data["json"] = c.PartyQueryRegResult()
		case "PartyQuery":
			c.Data["json"] = c.PartyQuery()
		case "PartyQueryGames":
			c.Data["json"] = c.PartyQueryGames()
		case "PartyList":
			c.Data["json"] = c.PartyList()
		case "PartyStateUpdate":
			c.Data["json"] = c.PartyStateUpdate()

		case "PartyQueryRegOrders":
			c.Data["json"] = PartyQueryRegOrders(c.js, ctx)

		case "QueryRegStateByCert":
			c.Data["json"] = QueryRegStateByCert(c.js)
		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}

	c.ServeJson()
}
