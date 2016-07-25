package registerSrv

import (
	"errors"
	"services/notice"
	"strings"
	"time"

	"github.com/chinarun/utils"
)

var (
	ErrEmailExists  = errors.New("邮箱已经已经被注册")
	ErrInvalidEmail = errors.New("邮箱格式不正确")
)

type EmailRegisterable interface {
	Registerable
	SetEmail(string)
	SetConfirmHost(string)
}

type EmailRegister struct {
	email       *Email
	confirmHost string
	*BaseRegister
}

func NewEmailRegister() *EmailRegister {
	return &EmailRegister{
		BaseRegister: NewBaseRegister(),
	}
}

func (r *EmailRegister) SetEmail(email string) {
	r.comeFromInt = ComeFromEmailInt
	r.email = NewEmail(email)
}

func (r *EmailRegister) SetConfirmHost(host string) {
	r.confirmHost = strings.TrimRight(host, "/") + "/"
}

func (r *EmailRegister) getDefaultName() string {
	return strings.Split(r.email.EmailAddress(), "@")[0]
}

func (r *EmailRegister) Valid() error {
	if valid := r.email.IsValid(); !valid {
		return ErrInvalidEmail
	}

	if err := r.validPassword(); err != nil {
		return err
	}

	if exists := r.email.IsExists(); exists {
		return ErrEmailExists
	}

	return nil
}

func (r *EmailRegister) setUser() {
	r.BaseRegister.setUser()

	r.user.Name = r.getDefaultName()
	r.user.Email = r.email.EmailAddress()
	r.user.ComeFrom = r.comeFromInt
}

func (r *EmailRegister) Create() (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("registerSrv.EmailRegister.Create : %s", err.Error())
		} else {
			utils.Logger.Debug("registerSrv.EmailRegister.Create : create success")
		}
	}()

	if err = r.Valid(); err != nil {
		return
	}

	r.setUser()
	if err = r.save(); err != nil {
		return
	}

	if err = r.sendActiveEmail(); err != nil {
		return
	}
	return
}

func (r *EmailRegister) sendActiveEmail() error {

	emailConfirm := NewEmailConfirm()
	if err := emailConfirm.Create(r.user.Id); err != nil {
		return err
	}
	confirmToken := emailConfirm.Token()
	confirmUrl := `<a href="` + r.confirmHost + confirmToken + `" target=_blank>` + r.confirmHost + confirmToken + `</a>`

	data := &noticeSrv.EmailData{
		To: r.email.EmailAddress(),
		Data: map[string]interface{}{
			"%url%":  confirmUrl,
			"%date%": time.Now().Format(time.RFC3339),
		},
	}
	templateName := "email_active"
	emailSrv := noticeSrv.NewEmail()
	emailData := []*noticeSrv.EmailData{data}
	err := emailSrv.SendByTemplate(templateName, emailData)
	if err != nil {
		return err
	}
	return nil
}

var _ EmailRegisterable = (*EmailRegister)(nil)
