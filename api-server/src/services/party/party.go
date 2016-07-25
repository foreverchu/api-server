package partySrv

import (
	"fmt"
	"models"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type SrvParty struct {
	modelsParty *models.Party
}

func (p *SrvParty) Close() (err error) {
	err = p.modelsParty.Close()
	return err
}

func IsPartyExists(party_id string) bool {
	return models.IsValueExists(models.DB_TABLE_PARTY, models.DB_ID, party_id)
}

func LikeParty(party_id string) (err error) {
	_, err = models.Orm.QueryTable(models.DB_TABLE_PARTY_DETAIL).Filter("party_id", party_id).Update(orm.Params{
		"like": orm.ColValue(orm.Col_Add, 1),
	})

	return
}

func DislikeParty(party_id string) (err error) {
	_, err = models.Orm.QueryTable(models.DB_TABLE_PARTY_DETAIL).Filter("party_id", party_id).Filter("like__gt", 0).Update(orm.Params{
		"like": orm.ColValue(orm.Col_Minus, 1),
	})

	return
}
func get_party_state_string(state uint8) string {
	switch state {
	case 0:
		return "等待审核中"
	case 1:
		return "已审核通过"
	case 2:
		return "审核未通过"
	default:
		return "unknow state"
	}

}

func UpdatePartyState(party_id int, state uint8) (err error) {
	var count int64
	count, err = models.Orm.QueryTable(models.DB_TABLE_PARTY).Filter("id", party_id).Filter("valid_state", state).Count()
	if err != nil {
		return
	}
	if count != 0 {
		return fmt.Errorf("party_id %d %s", party_id, get_party_state_string(state))
	}
	_, err = models.Orm.QueryTable(models.DB_TABLE_PARTY).Filter("id", party_id).Update(orm.Params{
		"valid_state": state,
	})

	return
}
func FindPartyNameById(id uint32) (name string) {
	party, err := models.GetPartyByID(strconv.FormatUint(uint64(id), 64))
	if err != nil {
		return ""
	}
	return party.Name
}
func GetPartyByName(name string) (p *SrvParty, err error) {
	p.modelsParty, err = models.GetPartyByName(name)
	return
}

func IsPartyIdExist(id string) bool {
	return models.IsValueExists(models.DB_TABLE_PARTY, models.DB_ID, id)
}

func IsPartyNameExist(name string) bool {
	return models.IsValueExists(models.DB_TABLE_PARTY, models.DB_PARTY_FN_NAME, name)
}

func NewSrvParty() *SrvParty {
	return &SrvParty{
		modelsParty: new(models.Party),
	}
}
func (p *SrvParty) GetModelsParty() *models.Party {
	return p.modelsParty
}
func (p *SrvParty) SetModelsParty(party *models.Party) {
	p.modelsParty = party
}

func (p *SrvParty) SetParamsForUpdate(params map[string]interface{}) *SrvParty {
	for k, v := range params {
		switch k {
		case models.DB_PARTY_FN_NAME:
			p.modelsParty.Name = v.(string)
		case models.DB_PARTY_FN_COUNTRY:
			p.modelsParty.Country = v.(string)
		case models.DB_PARTY_FN_PROVINCE:
			p.modelsParty.Province = v.(string)
		case models.DB_PARTY_FN_CITY:
			p.modelsParty.City = v.(string)
		case models.DB_PARTY_FN_CLOSE_TIME:
			p.modelsParty.CloseTime = v.(time.Time)
		case models.DB_PARTY_FN_START_TIME:
			p.modelsParty.StartTime = v.(time.Time)
		case models.DB_PARTY_FN_ADDRESS:
			p.modelsParty.Addr = v.(string)
		case models.DB_PARTY_FN_LOCLONG:
			p.modelsParty.LocLong = v.(float32)
		case models.DB_PARTY_FN_LOCLAT:
			p.modelsParty.LocLat = v.(float32)
		case models.DB_PARTY_FN_REG_END_TIME:
			p.modelsParty.RegEndTime = v.(time.Time)
		case models.DB_PARTY_FN_REG_START_TIME:
			p.modelsParty.RegStartTime = v.(time.Time)
		case models.DB_PARTY_FN_LIMITATION:
			p.modelsParty.Limitation = v.(uint32)
		case models.DB_PARTY_FN_LIMITATION_TYPE:
			p.modelsParty.LimitationType = v.(uint8)
		case models.DB_PARTY_FN_END_TIME:
			p.modelsParty.EndTime = v.(time.Time)
		case models.DB_PARTY_FN_VALID_STATE:
			p.modelsParty.ValidState = v.(uint8)
		default:
			continue
		}
	}

	return p
}

func (p *SrvParty) SetParamsForInsert(params map[string]interface{}) *SrvParty {
	for k, v := range params {
		switch k {
		case models.DB_PARTY_FN_USER_ID:
			p.modelsParty.UserId = v.(uint32)
		case models.DB_PARTY_FN_NAME:
			p.modelsParty.Name = v.(string)
		case models.DB_PARTY_FN_COUNTRY:
			p.modelsParty.Country = v.(string)
		case models.DB_PARTY_FN_PROVINCE:
			p.modelsParty.Province = v.(string)
		case models.DB_PARTY_FN_CITY:
			p.modelsParty.City = v.(string)
		case models.DB_PARTY_FN_CLOSE_TIME:
			p.modelsParty.CloseTime = v.(time.Time)
		case models.DB_PARTY_FN_START_TIME:
			p.modelsParty.StartTime = v.(time.Time)
		case models.DB_PARTY_FN_ADDRESS:
			p.modelsParty.Addr = v.(string)
		case models.DB_PARTY_FN_LOCLONG:
			p.modelsParty.LocLong = v.(float32)
		case models.DB_PARTY_FN_LOCLAT:
			p.modelsParty.LocLat = v.(float32)
		case models.DB_PARTY_FN_REG_END_TIME:
			p.modelsParty.RegEndTime = v.(time.Time)
		case models.DB_PARTY_FN_REG_START_TIME:
			p.modelsParty.RegStartTime = v.(time.Time)
		case models.DB_PARTY_FN_LIMITATION:
			p.modelsParty.Limitation = v.(uint32)
		case models.DB_PARTY_FN_LIMITATION_TYPE:
			p.modelsParty.LimitationType = v.(uint8)
		case models.DB_PARTY_FN_END_TIME:
			p.modelsParty.EndTime = v.(time.Time)
		default:
			continue
		}
	}
	return p
}
