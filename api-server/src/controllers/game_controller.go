package controllers

import "err_code"

type GameController struct {
	BaseController
}

func (c *GameController) Prepare() {
	c.BaseController.Prepare()
}

func (c *GameController) Finish() {
	defer c.BaseController.Finish()
}

func (c *GameController) Post() {
	retJson := init_retJson()

	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "GameEdit":
			c.Data["json"] = GameEdit(c.js, c.Ctx)
		case "GameCreate":
			c.Data["json"] = GameCreate(c.js, c.Ctx)
		case "GameClose":
			c.Data["json"] = GameClose(c.js, c.Ctx)
		case "GameInputScore":
			c.Data["json"] = GameInputScore(c.js, c.Ctx)
		case "GameQueryRegInfo":
			c.Data["json"] = GameQueryRegInfo(c.js, c.Ctx)
		case "GameQuery":
			c.Data["json"] = GameQuery(c.js, c.Ctx)
		case "GameList":
			c.Data["json"] = GameList(c.js, c.Ctx)

		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}

	c.ServeJson()
}
