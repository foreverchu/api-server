package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	DB_THRID_PARTY_REGISTER_FROM_ID = "from_id"
)

type ThirdPartyRegister struct {
	Id           uint
	UserId       uint32
	Type         uint8  // 1 =>  wechat  2=> weibo 3=> ?
	FromId       string // maybe opoenid
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time //将第三方的过期时间统一处理为time.Time
	UpdatedAt    time.Time `orm:"type(datetime)"`
	CreatedAt    time.Time `orm:"type(datetime)"`
}

func (r *ThirdPartyRegister) FindBy(conditions map[string]interface{}) error {
	qs := Orm.QueryTable(r)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err := qs.One(r)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrUserNotFound
	}
	return nil
}
