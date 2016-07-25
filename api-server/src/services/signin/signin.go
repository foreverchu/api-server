package signinSrv

import (
	"errors"
	"models"
	"services/auth"
	"services/register"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/chinarun/utils"
)

var (
	ErrIllegalSignin  = errors.New("不合法的登录")
	ErrUserNotFound   = errors.New("用户不存在")
	ErrUserNotActived = errors.New("用户还未激活")
	ErrSignin         = errors.New("账号登录失败")
)

type Signin struct {
	signinBy    string
	password    *registerSrv.Password
	autoSignin  bool
	valid       *validation.Validation
	user        *models.User
	tokenString string
}

func New(by, pwd string, auto bool) *Signin {
	return &Signin{
		signinBy:   by,
		password:   registerSrv.NewPassword(pwd),
		autoSignin: auto,
		valid:      &validation.Validation{},
		user:       &models.User{},
	}
}

func (s *Signin) IsSigninByEmail() bool {
	if v := s.valid.Email(s.signinBy, "email"); v.Ok {
		utils.Logger.Debug("SigninSrv.Signin.IsSigninByEmail : signinBy = %s", s.signinBy)
		return true
	}
	return false
}

func (s *Signin) IsSigninByPhone() bool {
	if v := s.valid.Mobile(s.signinBy, "phone"); v.Ok {
		utils.Logger.Debug("SigninSrv.Signin.IsSigninByPhone : signinBy = %s", s.signinBy)
		return true
	}
	return false

}

func (s *Signin) Do() (err error) {
	var actualErr error

	defer func() {
		if err != nil {
			utils.Logger.Error("signinSrv.Signin.Do : %s", actualErr.Error())
		} else {
			utils.Logger.Debug("signinSrv.Signin.Do : signin success, user_id = %d", s.user.Id)
		}
	}()

	conditions := make(map[string]interface{})

	switch {
	case s.IsSigninByEmail():
		conditions[models.DB_USER_EMAIL] = s.signinBy
	case s.IsSigninByPhone():
		conditions[models.DB_USER_PHONE] = s.signinBy
	default:
		err = ErrIllegalSignin
		actualErr = err
		return
	}

	utils.Logger.Debug("signinSrv.Signin.Do : conditions: %v", conditions)

	if err = s.user.FindBy(conditions); err == models.ErrUserNotFound {
		actualErr = err
		err = ErrUserNotFound
		return
	}

	utils.Logger.Debug("signinSrv.Signin.Do : s.user: %v", s.user)

	if !s.validPassword() {
		err = ErrSignin
		actualErr = err
		return
	}

	if !s.user.IsActive() {
		err = ErrUserNotActived
		actualErr = err
		return
	}

	err = s.genToken()
	if err != nil {
		actualErr = err
		return
	}
	//更新token
	err = s.updateUserToken()
	if err != nil {
		actualErr = err
		return
	}

	return nil
}

func (s *Signin) validPassword() bool {
	s.password.SetSalt(s.user.Salt)
	return s.password.IsEncryptedSame(s.user.Password)
}

func genToken(userId uint32, src uint8) (string, error) {
	auth := authSrv.New()
	data := map[string]interface{}{
		"uid": userId,
		"src": src,
	}
	return auth.Generate(data)
}

func (s *Signin) genToken() (err error) {
	s.tokenString, err = genToken(s.user.Id, s.user.ComeFrom)
	return
}

func (s *Signin) updateUserToken() error {
	s.user.Token = s.tokenString
	s.user.LastSigninAt = time.Now()
	return s.user.Update(models.DB_USER_LAST_SIGNIN_AT, models.DB_USER_TOKEN)
}

func (s *Signin) Token() (string, error) {
	if s.tokenString == "" {
		if err := s.genToken(); err != nil {
			return "", err
		}
		return s.tokenString, nil
	}
	return s.tokenString, nil
}

func (s *Signin) User() (*models.User, error) {
	return s.user, nil
}

func AutoSignin(userId uint32, comeFrom uint8) (token string, err error) {
	user := &models.User{
		Id: userId,
	}
	if err = models.Orm.Read(user); err != nil {
		return "", ErrUserNotFound
	}
	if token, err = genToken(userId, comeFrom); err != nil {
		return
	}
	user.Token = token
	user.LastSigninAt = time.Now()
	if err = user.Update(models.DB_USER_LAST_SIGNIN_AT, models.DB_USER_TOKEN); err != nil {
		return "", err
	}
	return
}
