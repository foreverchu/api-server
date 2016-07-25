package registerSrv

import (
	"errors"
	"models"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/chinarun/utils"
)

var (
	ErrPasswordEmpty      = errors.New("密码不能为空")
	ErrorCreateUserFailed = errors.New("创建用户失败")
	ErrorInvalidUser      = errors.New("无法获取用户信息")
	ErrPasswordNotSame    = errors.New("密码不一致")
)

const (
	ComeFromEmail     = "email"
	ComeFromPhone     = "phone"
	ComeFromWechat    = "wecaht"
	ComeFromEmailInt  = 1
	ComeFromPhoneInt  = 2
	ComeFromWechatInt = 3
)

type Registerable interface {
	SetPassword(string)
	SetPasswordConfirm(string)
	Create() error
	UserInfo() (*models.User, error)
}

type BaseRegister struct {
	password        *Password
	passwordConfirm string
	comeFrom        string
	comeFromInt     uint8
	valid           *validation.Validation
	user            *models.User
}

func NewBaseRegister() *BaseRegister {
	return &BaseRegister{
		valid: &validation.Validation{},
		user:  &models.User{},
	}
}

func (r *BaseRegister) SetPassword(pwd string) {
	r.password = NewPassword(pwd)
}

func (r *BaseRegister) SetPasswordConfirm(pwd string) {
	r.passwordConfirm = pwd
}

func (r *BaseRegister) Valid() error {
	defer func() {
		utils.Logger.Critical("registerSrv.BaseRegister.Valid: cant no call Valid func")
	}()

	panic("cant not call Valid")
}

func (r *BaseRegister) validPassword() error {
	if !r.password.IsConfirmSame(r.passwordConfirm) {
		return ErrPasswordNotSame
	}
	return r.password.Valid()
}

func (r *BaseRegister) setUser() {
	r.user.Salt = r.password.Salt()
	r.user.Password = r.password.GenPwd()
	r.user.CreatedAt = time.Now()
}

func (r *BaseRegister) save() (err error) {
	var actualErr error

	defer func() {
		if err != nil {
			utils.Logger.Error("registerSrv.BaseRegister.save : %s", actualErr.Error())
		} else {
			utils.Logger.Debug("registerSrv.BaseRegister.save :  success : UserId = %d", r.user.Id)
		}
	}()

	err = r.user.Insert()
	if err != nil {
		actualErr = err
		err = ErrorCreateUserFailed
		return
	}
	return
}

func (r *BaseRegister) UserInfo() (*models.User, error) {
	if r.user.Id != 0 {
		return r.user, nil
	} else {
		return nil, ErrorInvalidUser
	}
}
