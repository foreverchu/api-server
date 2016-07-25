package userSrv

import (
	"errors"
	"models"
	"services/register"

	"github.com/chinarun/utils"
)

var (
	ErrOldPassword       = errors.New("原密码不正确")
	ErrUpdateNewPassword = errors.New("更新新密码失败")
)

// PasswordReset 用于重新设置用户密码
type PasswordReset struct {
	user          *models.User
	oldPwd        *registerSrv.Password
	newPwd        *registerSrv.Password
	newPwdConfirm string
}

func NewPasswordReset(currentUser *models.User, oldPassword, newPassword, newPasswordConfirm string) *PasswordReset {
	return &PasswordReset{
		user:          currentUser,
		oldPwd:        registerSrv.NewPassword(oldPassword),
		newPwd:        registerSrv.NewPassword(newPassword),
		newPwdConfirm: newPasswordConfirm,
	}
}

func (pr *PasswordReset) oldPasswordValid() bool {
	pr.oldPwd.SetSalt(pr.user.Salt)
	return pr.oldPwd.IsEncryptedSame(pr.user.Password)
}

func (pr *PasswordReset) newPasswordConfirm() bool {
	return pr.newPwd.IsConfirmSame(pr.newPwdConfirm)
}

func (pr *PasswordReset) newPasswordValid() error {
	return pr.newPwd.Valid()
}

func (pr *PasswordReset) updatePassword() error {
	pr.user.Salt = pr.newPwd.Salt()
	pr.user.Password = pr.newPwd.Encryped()
	affectedRowsNum, err := models.Orm.Update(pr.user, models.DB_USER_PASSWORD, models.DB_USER_SALT)
	if affectedRowsNum < 1 || err != nil {
		return ErrUpdateNewPassword
	}
	return nil
}

func (pr *PasswordReset) Do() (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("accountSrv.PasswordReset.Do : %s", err.Error())
		} else {
			utils.Logger.Debug("accountSrv.PasswordReset.Do : pr %v", pr)
		}
	}()
	if !pr.newPasswordConfirm() {
		return registerSrv.ErrPasswordNotSame
	}

	if err = pr.newPasswordValid(); err != nil {
		return err
	}

	if !pr.oldPasswordValid() {
		return ErrOldPassword
	}

	if err = pr.updatePassword(); err != nil {
		return err
	}
	return nil
}
