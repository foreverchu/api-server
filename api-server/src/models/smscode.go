package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
)

var ErrSMSCodeNotFound = errors.New("models.SMScode.FindBy : not found")

var (
	DB_SMSCODE_PHONE      = "phone"
	DB_SMSCODE_CODE       = "code"
	DB_SMSCODE_USED_AT    = "used_at"
	DB_SMSCODE_CREATED_AT = "created_at"
)

type Smscode struct {
	Id        uint32
	Phone     string
	Code      string
	UsedAt    time.Time `orm:"type(datetime)"`
	CreatedAt time.Time `orm:"type(datetime)"`
}

func (s *Smscode) FindBy(conditions map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Debug("models.Smscode.FindBy : error : %s", err.Error())
		}
	}()
	qs := Orm.QueryTable(s)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err = qs.One(s)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrSMSCodeNotFound
	}
	return nil
}

func (s *Smscode) IsUsed() bool {
	return !s.UsedAt.IsZero()
}

func (s *Smscode) Update() error {
	affectedRowNum, err := Orm.Update(s)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}

func (s *Smscode) Insert() error {
	lastInsertId, err := Orm.Insert(s)
	if err != nil {
		return err
	}
	s.Id = uint32(lastInsertId)

	return nil
}

func (s *Smscode) IsRequestedCode() bool {
	err := Orm.QueryTable(DB_TALBE_SMSCODE).Filter(DB_SMSCODE_PHONE, s.Phone).Filter(DB_SMSCODE_USED_AT+"__isnull", true).OrderBy("-" + DB_SMSCODE_CREATED_AT).Limit(1).One(s)
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
