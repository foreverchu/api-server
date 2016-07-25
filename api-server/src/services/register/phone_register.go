package registerSrv

import (
	"errors"

	"github.com/chinarun/utils"
)

var (
	ErrInvalidPhone = errors.New("手机号码格式不正确")
	ErrPhoneExists  = errors.New("手机号码已经被注册")
	ErrSMSCode      = errors.New("短信验证码错误")
)

type PhoneRegisterable interface {
	Registerable
	SetPhone(string)
	SetSMSCode(string)
}

type PhoneRegister struct {
	phone   *Phone
	smsCode string
	*BaseRegister
}

func NewPhoneRegister() *PhoneRegister {
	return &PhoneRegister{
		BaseRegister: NewBaseRegister(),
	}
}

func (r *PhoneRegister) SetPhone(phone string) {
	r.comeFromInt = ComeFromPhoneInt
	r.phone = NewPhone(phone)
}

func (r *PhoneRegister) SetSMSCode(code string) {
	r.smsCode = code
}

func (r *PhoneRegister) getDefaultName() string {
	return r.phone.PhoneNumber()
}

// 根据struct的tag来获取更加方便
func (r *PhoneRegister) Valid() error {
	if valid := r.phone.IsValid(); !valid {
		return ErrInvalidPhone
	}

	if err := r.phone.ValidSMSCode(r.smsCode); err != nil {
		return err
	}

	if err := r.validPassword(); err != nil {
		return err
	}

	if exists := r.phone.IsExists(); exists {
		return ErrPhoneExists
	}

	return nil
}

func (r *PhoneRegister) setUser() {
	r.BaseRegister.setUser()

	r.user.Name = r.getDefaultName()
	r.user.Phone = r.phone.PhoneNumber()
	r.user.ComeFrom = r.comeFromInt
	// 如果是手机注册, 则是激活的
	r.user.Active = 1
}

func (r *PhoneRegister) Create() (err error) {

	defer func() {
		if err != nil {
			utils.Logger.Error("registerSrv.PhoneRegister.Create : %s", err.Error())
		} else {
			utils.Logger.Debug("registerSrv.PhoneRegister.Create : success")
		}
	}()

	if err = r.Valid(); err != nil {
		return
	}

	r.setUser()
	if err = r.save(); err != nil {
		return
	}

	if err = r.phone.UseCode(); err != nil {
		return
	}
	return nil
}

var _ PhoneRegisterable = (*PhoneRegister)(nil)
