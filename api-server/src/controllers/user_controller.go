package controllers

import (
	"err_code"
	"errors"

	"services/notice"
	"services/parse_params"
	"services/register"
	"services/signin"
	"services/user"

	"github.com/chinarun/utils"
)

const (
	HP_SIGNIN_BY            = "signin_by"
	HP_PASSWORD             = "password"
	HP_OLD_PASSWORD         = "old_password"
	HP_NEW_PASSWORD         = "new_password"
	HP_NEW_PASSWORD_CONFIRM = "new_password_confirm"
	HP_AUTH_SIGNIN          = "auto_signin"
	HP_PHONE                = "phone"
	HP_EMAIL                = "email"
	HP_PASSWORDCONFIRM      = "password_confirm"
	HP_SMSCODE              = "sms_code"
	HP_CONFIRMHOST          = "confirm_host"
	HP_COMEFROM             = "come_from"
	HP_WECHAT_AUTH_CODE     = "code"
)

type UserController struct {
	BaseController
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

func (c *UserController) Finish() {
	defer c.BaseController.Finish()
}

func (c *UserController) Post() {
	retJson := init_retJson()

	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "UserSignin":
			c.Data["json"] = c.Signin()
		case "UserRegister":
			c.Data["json"] = c.Register()
		case "UserInfoQuery":
			c.Data["json"] = c.Info()
		case "UserWechatSignin":
			c.Data["json"] = c.WechatSignin()
		case "UserResetPassword":
			c.Data["json"] = c.ResetPassword()
		case "UserForgetPassword":
			c.Data["json"] = c.ForgetPassword()
		case "UserEmailConfirm":
			c.Data["json"] = c.EmailConfirm()
		case "UserRegisterSMSCode":
			c.Data["json"] = c.SMSCode()
		case "EditPersonalInfo":
			c.Data["json"] = c.EditPersonalInfo()
		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}
	c.ServeJson()
}

var ErrRegister = errors.New("非法注册")

func (c *UserController) EditPersonalInfo() map[string]interface{} {
	retJson := init_retJson()
	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.EditPersonalInfo : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.EditPersonalInfo : %#v", c.js)
		}
	}()

	if err = c.IsLogin(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	param_values, err := utils.ParseParam(c.js, parseParamsSrv.GetEditPersonalInfoMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	err = c.CurrentUser.Update(param_values)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	return retJson
}

func (c *UserController) Signin() map[string]interface{} {
	var err error
	retJson := init_retJson()
	by, _ := c.js.Get(HP_SIGNIN_BY).String()
	password, _ := c.js.Get(HP_PASSWORD).String()
	authSignin, _ := c.js.Get(HP_AUTH_SIGNIN).Bool()

	signin := signinSrv.New(by, password, authSignin)
	err = signin.Do()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	token, err := signin.Token()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	retJson["token"] = token
	retJson["user"], _ = signin.User()
	return retJson
}

func (c *UserController) Register() map[string]interface{} {
	retJson := init_retJson()
	var reg registerSrv.Registerable

	password, _ := c.js.Get(HP_PASSWORD).String()
	passwordConfirm, _ := c.js.Get(HP_PASSWORDCONFIRM).String()

	comeFrom, _ := c.js.Get(HP_COMEFROM).String()
	switch comeFrom {
	case registerSrv.ComeFromEmail:
		var emailReg registerSrv.EmailRegisterable
		emailReg = registerSrv.NewEmailRegister()
		email, _ := c.js.Get(HP_EMAIL).String()
		emailReg.SetEmail(email)
		confirmHost, _ := c.js.Get(HP_CONFIRMHOST).String()
		emailReg.SetConfirmHost(confirmHost)
		emailReg.SetPassword(password)
		emailReg.SetPasswordConfirm(passwordConfirm)
		reg = emailReg.(registerSrv.Registerable)

	case registerSrv.ComeFromPhone:
		var phoneReg registerSrv.PhoneRegisterable
		phoneReg = registerSrv.NewPhoneRegister()
		phone, _ := c.js.Get(HP_PHONE).String()
		phoneReg.SetPhone(phone)
		smsCode, _ := c.js.Get(HP_SMSCODE).String()
		phoneReg.SetSMSCode(smsCode)
		phoneReg.SetPassword(password)
		phoneReg.SetPasswordConfirm(passwordConfirm)
		reg = phoneReg.(registerSrv.Registerable)

	default:
		retJson_edit(retJson, err_code.InvalidData, ErrRegister.Error())
		return retJson
	}

	if err := reg.Create(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	if comeFrom == registerSrv.ComeFromPhone {
		user, err := reg.UserInfo()
		utils.Logger.Debug("controllers.UserController.Register : user = %v", user)
		if err != nil {
			retJson_edit(retJson, err_code.InvalidData, err.Error())
			return retJson
		}

		token, err := signinSrv.AutoSignin(user.Id, user.ComeFrom)
		if err != nil {
			retJson_edit(retJson, err_code.InvalidData, err.Error())
			return retJson
		}
		retJson["token"] = token
		retJson["user"] = user
		return retJson
	}

	return retJson
}

func (c *UserController) Info() map[string]interface{} {
	retJson := init_retJson()
	userId, _ := c.js.Get(HP_USER_ID).String()
	user, err := userSrv.NewUser(userId)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	returnUser := make(map[string]interface{})
	returnUser["uesrId"] = userId
	returnUser["user"] = user.User()
	returnUser["profile"] = user.Profile()
	retJson["user"] = returnUser
	return retJson
}

func (c *UserController) WechatSignin() map[string]interface{} {
	retJson := init_retJson()
	code, _ := c.js.Get(HP_WECHAT_AUTH_CODE).String()
	wechatAPI := registerSrv.NewWechatAPI()
	wechatAPI.SetCode(code)
	err := wechatAPI.Signin()
	if err != nil {
		utils.Logger.Debug("%s", err.Error())
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	user, err := wechatAPI.UserInfo()
	if err != nil {
		utils.Logger.Debug("%s", err.Error())
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	token, err := signinSrv.AutoSignin(user.Id, registerSrv.ComeFromWechatInt)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	retJson["token"] = token
	retJson["user"] = user

	return retJson
}

func (c *UserController) ResetPassword() map[string]interface{} {
	retJson := init_retJson()

	if err := c.IsLogin(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	oldPassword, _ := c.js.Get(HP_OLD_PASSWORD).String()
	newPassword, _ := c.js.Get(HP_NEW_PASSWORD).String()
	newPasswordConfirm, _ := c.js.Get(HP_NEW_PASSWORD_CONFIRM).String()
	account := userSrv.NewPasswordReset(c.CurrentUser.User(), oldPassword, newPassword, newPasswordConfirm)
	if err := account.Do(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	token, err := signinSrv.AutoSignin(c.CurrentUser.User().Id, c.CurrentUser.User().ComeFrom)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	retJson["token"] = token
	retJson["user"] = c.CurrentUser
	return retJson

	return retJson
}

func (c *UserController) ForgetPassword() map[string]interface{} {
	retJson := init_retJson()
	return retJson
}

// SMSCode 用于获取短信验证码
func (c *UserController) SMSCode() map[string]interface{} {
	retJson := init_retJson()
	phoneNumber, _ := c.js.Get("phone").String()
	phone := registerSrv.NewPhone(phoneNumber)
	smser := noticeSrv.NewSms()
	err := phone.SendSMS(smser)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	return retJson
}

func (c *UserController) EmailConfirm() map[string]interface{} {
	retJson := init_retJson()
	confirmToken, _ := c.js.Get("token").String()
	ecSrv := registerSrv.NewEmailConfirm()
	err := ecSrv.Confirm(confirmToken)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	activedUser := ecSrv.ActivedUserInfo()
	token, err := signinSrv.AutoSignin(activedUser.Id, registerSrv.ComeFromEmailInt)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	retJson["token"] = token
	retJson["user"] = activedUser
	return retJson
}
