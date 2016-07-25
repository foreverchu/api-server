package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
)

var (
	ErrPartyNotFound = errors.New("models.Party.FindBy : party not found")
)

const (
	DB_PARTY_FN_NAME            = "name"
	DB_PARTY_FN_REG_START_TIME  = "reg_start_time"
	DB_PARTY_FN_CLOSE_TIME      = "close_time"
	DB_PARTY_FN_CITY            = "city"
	DB_PARTY_FN_PROVINCE        = "province"
	DB_PARTY_FN_COUNTRY         = "country"
	DB_PARTY_FN_ADDRESS         = "addr"
	DB_PARTY_FN_LOCLONG         = "loc_long"
	DB_PARTY_FN_LOCLAT          = "loc_lat"
	DB_PARTY_FN_REG_END_TIME    = "reg_end_time"
	DB_PARTY_FN_START_TIME      = "start_time"
	DB_PARTY_FN_END_TIME        = "end_time"
	DB_PARTY_FN_LIMITATION      = "limitation"
	DB_PARTY_FN_LIMITATION_TYPE = "limitation_type"
	DB_PARTY_FN_VALID_STATE     = "valid_state"
	DB_PARTY_FN_USER_ID         = "user_id"
)

type Party struct {
	Id           uint32
	UserId       uint32
	Name         string `orm:"size(60)"`
	Country      string `orm:"size(32)"`
	Province     string `orm:"size(32)"`
	City         string `orm:"size(32)"`
	Addr         string `orm:"size(90)"`
	LocLong      float32
	LocLat       float32
	RegStartTime time.Time `orm:"type(datetime)"`
	RegEndTime   time.Time `orm:"type(datetime)"`
	StartTime    time.Time `orm:"type(datetime)"`
	EndTime      time.Time `orm:"type(datetime)"`
	CloseTime    time.Time `orm:"type(datetime)"`
	ValidState   uint8

	Limitation     uint32
	LimitationType uint8
}

type PartyAndDetail struct {
	PartyId      uint32 `orm:"pk"`
	Slogan       string `orm:"size(128)"`
	Like         uint32
	Website      string `orm:"size(60)"`
	Type         string `orm:"size(32)"`
	Price        string `orm:"size(60)"`
	Introduction string `orm:"type(text)"`
	Schedule     string `orm:"type(text)"`
	Score        float32
	SignupMale   uint32
	SignupFemale uint32

	// From party table.
	Name         string `orm:"size(60)"`
	Country      string `orm:"size(32)"`
	Province     string `orm:"size(32)"`
	City         string `orm:"size(32)"`
	Addr         string `orm:"size(90)"`
	LocLong      float32
	LocLat       float32
	RegStartTime time.Time `orm:"type(datetime)"`
	RegEndTime   time.Time `orm:"type(datetime)"`
	StartTime    time.Time `orm:"type(datetime)"`
	EndTime      time.Time `orm:"type(datetime)"`
	CloseTime    time.Time `orm:"type(datetime)"`

	Limitation     uint32
	LimitationType uint8

	IfClose uint8

	PhotoUrl string
	Tag      string
}

func GetPartyByID(party_id string) (*Party, error) {
	party_id_int, err := strconv.Atoi(party_id)
	if err != nil {
		return nil, err
	}

	party := &Party{
		Id: uint32(party_id_int),
	}
	err = party.Get()
	return party, err
}

func (p *Party) FindBy(conditions map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Debug("models.Order.FindBy : error : %s, condtions : %v", err.Error(), conditions)
		}
	}()
	qs := Orm.QueryTable(p)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err = qs.One(p)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrPartyNotFound
	}
	return nil
}

// IsClosed 表示赛事是否已经关闭
func (p *Party) IsClosed() bool {
	if p.CloseTime.IsZero() {
		return false
	}
	return true
}
func (p *Party) Close() error {
	if p.Id == 0 {
		return errors.New("PartyNotExist!")
	}
	p.CloseTime = time.Now().Local()
	return p.Update()
}
func (p *Party) IsChecked() bool {
	if p.ValidState == 0 {
		return false
	}
	return true
}

func GetPartyByName(name string) (*Party, error) {
	party := NewParty()

	qs := Orm.QueryTable("party").Filter("name", name)
	count, err := qs.Count()
	if err != nil {
		return party, err
	}
	if count != 1 {
		return party, errors.New("NoParty")
	}
	err = qs.One(party)
	if err != nil {
		return party, err
	}
	return party, err
}

func NewParty() *Party {
	return new(Party)
}

func (p *Party) Insert() error {
	_, err := Orm.Insert(p)
	return err
}

func (p *Party) Update() error {
	_, err := Orm.Update(p)
	return err
}

func (p *Party) Delete() error {
	_, err := Orm.Delete(p)
	return err
}

func (p *Party) Get() error {
	err := Orm.Read(p)
	return err
}
