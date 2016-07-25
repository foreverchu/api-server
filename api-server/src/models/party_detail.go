package models

const (
	//partt_detail
	DB_PARTY_DETAIL_FN_PARTY_ID      = "party_id"
	DB_PARTY_DETAIL_FN_SLOGAN        = "slogan"
	DB_PARTY_DETAIL_FN_LIKE          = "like"
	DB_PARTY_DETAIL_FN_WEBSITE       = "website"
	DB_PARTY_DETAIL_FN_TYPE          = "type"
	DB_PARTY_DETAIL_FN_PRICE         = "price"
	DB_PARTY_DETAIL_FN_INTRODUCTION  = "introduction"
	DB_PARTY_DETAIL_FN_SCHEDULE      = "schedule"
	DB_PARTY_DETAIL_FN_SCORE         = "score"
	DB_PARTY_DETAIL_FN_SIGNUP_MALE   = "signup_male"
	DB_PARTY_DETAIL_FN_SIGNUP_FEMALE = "signup_female"
)

type PartyDetail struct {
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
}

func NewPartyDetail() *PartyDetail {
	return new(PartyDetail)
}

func (p *PartyDetail) Insert() error {
	_, err := Orm.Insert(p)
	return err
}

func (p *PartyDetail) Update() error {
	_, err := Orm.Update(p)
	return err
}

func (p *PartyDetail) Delete() error {
	_, err := Orm.Delete(p)
	return err
}

func (p *PartyDetail) Get() error {
	err := Orm.Read(p)
	return err
}

func GetPartyDetailByPartyId(partyId uint32) (*PartyDetail, error) {
	partyDetail := &PartyDetail{
		PartyId: partyId,
	}
	err := partyDetail.Get()

	return partyDetail, err
}
