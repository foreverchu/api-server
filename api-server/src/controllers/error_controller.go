package controllers

import "github.com/chinarun/utils"

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	utils.Logger.Error("controllers.ErrorController.Error404 : error")
	c.Abort("404")
}

func (c *ErrorController) Error500() {
	utils.Logger.Error("controllers.ErrorController.Error500 : error")
	c.Ctx.WriteString("<h1> hello world! </h1>")
	c.Abort("500")
}
