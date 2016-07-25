package partySrv

import (
	"models"
)

type SrvPartyDetail struct {
	modelsPartyDetail *models.PartyDetail
}

func NewSrvPartyDetail() *SrvPartyDetail {
	return &SrvPartyDetail{
		modelsPartyDetail: new(models.PartyDetail),
	}
}

func (p *SrvPartyDetail) SetParams(params map[string]interface{}) *SrvPartyDetail {
	for k, v := range params {
		switch k {
		case models.DB_PARTY_DETAIL_FN_SLOGAN:
			p.modelsPartyDetail.Slogan = v.(string)
		case models.DB_PARTY_DETAIL_FN_LIKE:
			p.modelsPartyDetail.Like = v.(uint32)
		case models.DB_PARTY_DETAIL_FN_WEBSITE:
			p.modelsPartyDetail.Website = v.(string)
		case models.DB_PARTY_DETAIL_FN_TYPE:
			p.modelsPartyDetail.Type = v.(string)
		case models.DB_PARTY_DETAIL_FN_PRICE:
			p.modelsPartyDetail.Price = v.(string)
		case models.DB_PARTY_DETAIL_FN_INTRODUCTION:
			p.modelsPartyDetail.Introduction = v.(string)
		case models.DB_PARTY_DETAIL_FN_SCHEDULE:
			p.modelsPartyDetail.Schedule = v.(string)
		case models.DB_PARTY_DETAIL_FN_SCORE:
			p.modelsPartyDetail.Score = v.(float32)
		case models.DB_PARTY_DETAIL_FN_SIGNUP_MALE:
			p.modelsPartyDetail.SignupMale = v.(uint32)
		case models.DB_PARTY_DETAIL_FN_SIGNUP_FEMALE:
			p.modelsPartyDetail.SignupFemale = v.(uint32)

		default:
			continue
		}
	}
	return p
}

func (p *SrvPartyDetail) GetModelsPartyDetail() *models.PartyDetail {
	return p.modelsPartyDetail
}
func (p *SrvPartyDetail) SetModelsPartyDetail(partyDetail *models.PartyDetail) {
	p.modelsPartyDetail = partyDetail
}
func (p *SrvPartyDetail) GetPartyDetailByPartyId(partyId uint32) {

}
