package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

var ErrUserProfileNotFound = errors.New("models.Profile.FindBy : user profile not found")

const (
	DB_PROFILE_ADDRESS       = "address"
	DB_PROFILE_GENDER        = "gender"
	DB_PROFILE_CONSTELLATION = "constellation"
	DB_PROFILE_PROFESSION    = "profession"
	DB_PROFILE_ABOUT         = "about"
	DB_PROFILE_CREATEDAT     = "created_at"
	DB_PROFILE_UPDATEDAT     = "updated_at"
)

type Profile struct {
	Id            uint
	UserId        uint32
	Address       string
	Gender        uint8
	Constellation uint8 //星座
	Profession    string
	About         string
	CreatedAt     time.Time `orm:"type(datetime)"`
	UpdatedAt     time.Time `orm:"type(datetime)"`
}

func (p *Profile) FindBy(conditions map[string]interface{}) error {

	qs := Orm.QueryTable(p)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err := qs.One(p)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrUserProfileNotFound
	}
	return nil
}

func (p *Profile) UpdateProfile(cond map[string]interface{}) error {
	for column, value := range cond {
		switch column {
		case DB_PROFILE_ADDRESS:
			p.Address = value.(string)
		case DB_PROFILE_ABOUT:
			p.About = value.(string)
		case DB_PROFILE_GENDER:
			p.Gender = value.(uint8)
		case DB_PROFILE_CONSTELLATION:
			p.Constellation = value.(uint8)
		case DB_PROFILE_PROFESSION:
			p.Profession = value.(string)
		case DB_PROFILE_CREATEDAT:
			p.CreatedAt = value.(time.Time)
		case DB_PROFILE_UPDATEDAT:
			p.UpdatedAt = value.(time.Time)

		}
	}
	_, err := Orm.Update(p)
	return err
}
