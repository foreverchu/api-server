package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	DB_USER_NAME           = "name"
	DB_USER_EMAIL          = "email"
	DB_USER_PHONE          = "phone"
	DB_USER_PASSWORD       = "password"
	DB_USER_SALT           = "salt"
	DB_USER_AVATAR         = "avatar"
	DB_USER_ACTIVE         = "active"
	DB_USER_COMEFROM       = "come_from"
	DB_USER_TOKEN          = "token"
	DB_USER_CREATED_AT     = "created_at"
	DB_USER_LAST_SIGNIN_AT = "last_signin_at"
)

var (
	ErrUserNotFound = errors.New("models.User.FindBy : user not found")
)

type User struct {
	Id           uint32
	Name         string
	Phone        string `json:"-"`
	Email        string `json:"-"`
	Password     string `json:"-"`
	Salt         string `json:"-"`
	Avatar       string
	ComeFrom     uint8
	Active       uint8
	Token        string    `json:"-"`
	LastSigninAt time.Time `orm:"type(datetime)"`
	CreatedAt    time.Time `orm:"type(datetime)"`
}

func (u *User) FindBy(conditions map[string]interface{}) error {
	qs := Orm.QueryTable(u)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	//暂时不需要过滤没有激活的用户
	//qs = qs.Filter(DB_USER_ACTIVE, 1)
	err := qs.One(u, DB_ID, DB_USER_NAME, DB_USER_PHONE, DB_USER_EMAIL, DB_USER_AVATAR, DB_USER_COMEFROM, DB_USER_ACTIVE, DB_USER_LAST_SIGNIN_AT, DB_USER_CREATED_AT, DB_USER_PASSWORD, DB_USER_SALT, DB_USER_TOKEN)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrUserNotFound
	}
	return nil
}

// IsActive判断此用户是否激活了
func (u *User) IsActive() bool {
	return u.Active == 1
}

func (u *User) Update(columns ...string) error {
	affectedRowNum, err := Orm.Update(u, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}

func (u *User) Insert() error {
	lastInsertId, err := Orm.Insert(u)
	if err != nil {
		return err
	}
	u.Id = uint32(lastInsertId)
	return nil
}
