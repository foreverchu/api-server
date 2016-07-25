package controllers

import "err_code"

type RegController struct {
	BaseController
}

func (c *RegController) Prepare() {
	c.BaseController.Prepare()
}

func (c *RegController) Finish() {
	defer c.BaseController.Finish()
}

func (c *RegController) Post() {
	retJson := init_retJson()

	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "RegGetQueueToken":
			c.Data["json"] = RegGetQueueToken(c.js)
		case "RegQueryQueueTokenState":
			c.Data["json"] = RegQueryQueueTokenState(c.js)
			/*
				case "RegSetGameLimitation":
					c.Data["json"] = RegSetGameLimitation(c.js, c.Ctx)
				case "RegGetGameLimitationInfo":
					c.Data["json"] = RegGetGameLimitationInfo(c.js, c.Ctx)
				case "RegGetAllGamesBalanceInfo":
					c.Data["json"] = RegGetAllGamesBalanceInfo(c.js, c.Ctx)
				case "RegSetAllGamesLimitation":
					c.Data["json"] = RegSetAllGamesLimitation(c.js, c.Ctx)
			*/

		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}

	c.ServeJson()
}
