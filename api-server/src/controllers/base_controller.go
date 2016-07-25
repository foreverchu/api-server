package controllers

import (
	"err_code"
	"errors"
	"fmt"
	"net/http"

	"services/auth"
	"services/notice"
	"services/user"
	"time"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"github.com/chinarun/utils"
)

const (
	VERSION = "chinarun 1.0"

	PARAM_REQUESTED = true
	PARAM_OPTIONAL  = false
)

type BaseController struct {
	beego.Controller
	CurrentUser *userSrv.User
	js          *simplejson.Json
	cmd         string
	Admin       bool
}

func (c *BaseController) Prepare() {
	/*
		if !utils.IsClientFromSamePrivateNetwork(c.Ctx) {
			c.Abort("403")
		}
	*/
	c.ApplyAccessControlConfig()

	retJson := init_retJson()
	js, cmd, errcode := c.ValidJsonParams()
	if errcode != err_code.OK {
		retJson_edit(retJson, errcode, "")
		c.Data["json"] = retJson
		c.ServeJson()
		return
	}
	c.js = js
	c.cmd = cmd

	c.IsAdmin()
}

func getAdminKey(admin, key, time string) string {
	return utils.GetMd5(utils.GetMd5(fmt.Sprintf("%s%s%s", admin, key, time)))
}

func (c *BaseController) IsAdmin() {
	admin := c.Ctx.Input.Header("admin")
	if admin == "" {
		c.Admin = false
		return
	}

	auth := c.Ctx.Input.Header("auth")
	if auth == "" {
		c.Admin = false
		return
	}

	time := c.Ctx.Input.Header("time")
	if time == "" {
		c.Admin = false
		return
	}
	if getAdminKey(admin, noticeSrv.Key, time) == auth {
		c.Admin = true
		return
	}
	c.Admin = false
	return
}

func (c *BaseController) Finish() {
	if len(c.Ctx.Input.RequestBody) > 0 {
		utils.DevErrPrintDefer(c.Data["json"], string(c.Ctx.Input.RequestBody))
	}

	if c.js == nil {
		return
	}

	client_ip := GetClientIp(c.Ctx.Request)
	if cmd, err := c.js.Get("cmd").String(); err != nil {
		utils.Logger.Info("%s, Method: %s, json: %v.", client_ip, cmd, c.Data["json"])
	}

}

func (c *BaseController) Get() {
	c.Ctx.WriteString("ChinaRun API.")
	return
}

func (c *BaseController) ApplyAccessControlConfig() {
	str, _ := utils.Cfg.GetString("beego", "Access-Control-Allow-Origin")
	c.Ctx.Output.Header("Access-Control-Allow-Origin", str)
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Authorization") //for jwt
}

func (c *BaseController) ValidJsonParams() (js *simplejson.Json, cmd string, errcode int) {
	errcode = err_code.OK
	js, err := simplejson.NewJson(c.Ctx.Input.RequestBody)
	if err != nil {
		return nil, "", err_code.ErrJson
	}
	cmd, err = js.Get("cmd").String()
	if err != nil || cmd == "" {
		return nil, "", err_code.NoCmd
	}

	return
}

func (c *BaseController) IsLogin() (err error) {
	token := authSrv.NewHeaderTokenParser(c.Ctx.Request)

	if err = token.Parse(); err != nil {
		return err
	}

	tokenString := token.Token()

	auth := authSrv.New()
	if ok, err := auth.Verify(tokenString); !ok {
		return err
	}
	uid := auth.Claims()["uid"].(float64)
	userId := fmt.Sprintf("%d", int(uid))

	if c.CurrentUser, err = userSrv.NewUser(userId); err != nil {
		return err
	}

	if c.CurrentUser.User().Token != tokenString {
		return errors.New("token已经过期")
	}

	return nil
}

func (c *BaseController) GetCurrentUser() *userSrv.User {
	return c.CurrentUser
}

const (
	RETURN_JSON_RESULT = "result"
	RETURN_JSON_MSG    = "msg"
)

func GetClientIp(request *http.Request) string {
	return request.RemoteAddr
}

func Get_file_name() string {
	t := time.Now()
	filename := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%03d%06d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), int(t.Nanosecond()/1000000), utils.Rander.Intn(1000000))
	return filename

}

func init_retJson() map[string]interface{} {
	var retJson map[string]interface{}
	retJson = make(map[string]interface{})
	retJson[RETURN_JSON_RESULT] = err_code.OK
	retJson[RETURN_JSON_MSG] = ""
	return retJson
}

func retJson_edit(retJson map[string]interface{}, result int, msg string) {
	retJson[RETURN_JSON_RESULT] = result
	_msg := err_code.Get_err_msg(result)
	if msg == "" {
		retJson[RETURN_JSON_MSG] = _msg
	} else {
		retJson[RETURN_JSON_MSG] = _msg + "\n" + msg
	}

}
