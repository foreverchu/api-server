package registerSrv

import (
	"errors"
	"models"
	"time"

	"github.com/chinarun/utils"
)

var ErrUserNotFound = errors.New("用户不存在")
var ErrCreateEmailConfirm = errors.New("创建邮箱验证失败")
var ErrInvalidToken = errors.New("不合法的请求")
var ErrUsedToken = errors.New("请求已经被处理")
var ErrExpiredToken = errors.New("请求已经过期")
var ErrDBUpdateError = errors.New("更新出错")

const (
	ExpiredDays           = 30
	EmailConfirmTokenSize = 64
)

type EmailConfirm struct {
	token string
	user  *models.User
	ec    *models.EmailConfirm
}

func NewEmailConfirm() *EmailConfirm {
	return &EmailConfirm{
		ec:   &models.EmailConfirm{},
		user: &models.User{},
	}
}

// CreateEmailConfirm 创建一个EmailConfirm
func (c *EmailConfirm) Create(userId uint32) (err error) {
	var actualErr error

	defer func() {
		if err != nil {
			utils.Logger.Error("registerSrv.EmailConfirm.Create : %s", actualErr.Error())
		} else {
			utils.Logger.Debug("registerSrv.EmailConfirm.Create : success, user_id : %d", userId)
		}
	}()

	c.genToken()
	if err = c.createEmailConfirm(userId); err != nil {
		return
	}

	return nil
}

func (c *EmailConfirm) createEmailConfirm(userId uint32) (err error) {
	c.ec.UserId = userId
	c.ec.Token = c.Token()
	c.ec.CreatedAt = time.Now()
	c.ec.ExpiredAt = time.Now().AddDate(0, 0, ExpiredDays)

	if _, err = models.Orm.Insert(c.ec); err != nil {
		return ErrCreateEmailConfirm
	}
	return nil
}

func (c *EmailConfirm) genToken() {
	c.token = generateRandomString(EmailConfirmTokenSize)
}

func (c *EmailConfirm) Token() string {
	if c.token == "" {
		c.genToken()
	}
	return c.token
}

func (c *EmailConfirm) validToken() error {
	if len(c.token) != EmailConfirmTokenSize {
		return ErrInvalidToken
	}
	return nil
}

func (c *EmailConfirm) Confirm(token string) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Error("registerSrv.EmailConfirm.Confirm : %s", err.Error())
		} else {
			utils.Logger.Debug("registerSrv.EmailConfirm.Confirm : success: %s", token)
		}
	}()

	c.token = token

	if err = c.validToken(); err != nil {
		return
	}

	if err = c.findRecord(); err != nil {
		return
	}

	if err = c.ActiveEmailConfirm(); err != nil {
		return
	}

	return nil
}

func (c *EmailConfirm) findRecord() (err error) {
	cond := map[string]interface{}{
		models.DB_EMAIL_CONFIRM_TOKEN: c.token,
	}
	err = c.ec.FindBy(cond)
	if err == models.ErrEmailConfirmNotFound {
		return ErrInvalidToken
	}

	if c.ec.IsUsed() {
		return ErrUsedToken
	}

	if c.ec.IsExpired() {
		return ErrExpiredToken
	}
	return nil
}

func (c *EmailConfirm) ActiveEmailConfirm() (err error) {
	models.Orm.Begin()

	if err = c.UseToken(); err != nil {
		models.Orm.Rollback()
		err = ErrDBUpdateError
		return
	}

	if err = c.ActiveUser(c.ec.UserId); err != nil {
		models.Orm.Rollback()
		err = ErrDBUpdateError
		return
	}

	models.Orm.Commit()

	return nil
}

func (c *EmailConfirm) UseToken() error {
	c.ec.Used = 1
	c.ec.ConfirmedAt = time.Now()
	return c.ec.Update(models.DB_EMAIL_CONFIRM_USED, models.DB_EMAIL_CONFIRM_CONFIRMED_AT)
}

func (c *EmailConfirm) ActiveUser(userId uint32) error {
	cond := map[string]interface{}{
		models.DB_ID: userId,
	}
	if err := c.user.FindBy(cond); err == models.ErrUserNotFound {
		return ErrUserNotFound
	}
	c.user.Active = 1
	return c.user.Update(models.DB_USER_ACTIVE)
}

func (c *EmailConfirm) ActivedUserInfo() *models.User {
	return c.user
}
