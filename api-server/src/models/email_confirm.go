package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

var ErrEmailConfirmNotFound = errors.New("models.EmailConfirmToken.FindBy: email confirm not found")

const (
	DB_EMAIL_CONFIRM_USERID       = "user_id"
	DB_EMAIL_CONFIRM_TOKEN        = "token"
	DB_EMAIL_CONFIRM_USED         = "used"
	DB_EMAIL_CONFIRM_CONFIRMED_AT = "confirmed_at"
)

type EmailConfirm struct {
	Id          uint
	UserId      uint32
	Token       string
	Used        uint8
	CreatedAt   time.Time `orm:"type(datetime)"`
	ConfirmedAt time.Time `orm:"type(datetime)"`
	ExpiredAt   time.Time `orm:"type(datetime)"`
}

func (e *EmailConfirm) FindBy(conditions map[string]interface{}) error {
	qs := Orm.QueryTable(e)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err := qs.One(e)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrEmailConfirmNotFound
	}
	return nil

}

func (e *EmailConfirm) Update(columns ...string) error {
	affectedRowNum, err := Orm.Update(e, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}

func (e *EmailConfirm) IsUsed() bool {
	return e.Used == 1
}

func (e *EmailConfirm) IsExpired() bool {
	return e.ExpiredAt.Before(time.Now())
}
