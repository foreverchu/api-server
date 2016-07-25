package partyOperationSrv

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	"models"
	"services/party"
	"services/photo"
	"services/registration"
	"services/user"
)

type PartyOperation struct {
	gameList       []*models.Game
	user           *userSrv.User
	srvParty       *partySrv.SrvParty
	srvPartyDetail *partySrv.SrvPartyDetail
	photos         *photoSrv.SrvPhoto
}

func (p *PartyOperation) GetPhoto() *photoSrv.SrvPhoto {
	return p.photos
}

func (p *PartyOperation) GetUser() *userSrv.User {
	return p.user
}
func (p *PartyOperation) GetSrvParty() *partySrv.SrvParty {
	return p.srvParty
}

func (p *PartyOperation) GetSrvPartyDetail() *partySrv.SrvPartyDetail {
	return p.srvPartyDetail
}

func (p *PartyOperation) GetGameList() []*models.Game {
	return p.gameList
}

func NewPartyOperationForCreate(user_id uint32) *PartyOperation {
	p := &PartyOperation{
		srvParty:       partySrv.NewSrvParty(),
		srvPartyDetail: partySrv.NewSrvPartyDetail(),
	}

	p.srvParty.GetModelsParty().UserId = user_id
	return p
}

func (p *PartyOperation) SetParamsForCreate(params map[string]interface{}) *PartyOperation {
	p.srvParty = p.srvParty.SetParamsForInsert(params)
	p.srvPartyDetail = p.srvPartyDetail.SetParams(params)
	return p
}

func (p *PartyOperation) SetParamsForUpdate(params map[string]interface{}) *PartyOperation {
	if p.srvParty.GetModelsParty().IsClosed() {
		return nil
	}
	p.srvParty = p.srvParty.SetParamsForUpdate(params)
	p.srvPartyDetail = p.srvPartyDetail.SetParams(params)
	return p
}

func (p *PartyOperation) Create() (partyId int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return partyId, err
	}

	partyId, err = o.Insert(p.srvParty.GetModelsParty())
	if err != nil {
		err1 := o.Rollback()
		if err1 != nil {
			return partyId, err1
		}
		return partyId, err
	}

	p.srvPartyDetail.GetModelsPartyDetail().PartyId = uint32(partyId)
	_, err = o.Insert(p.srvPartyDetail.GetModelsPartyDetail())
	if err != nil {
		err1 := o.Rollback()
		if err != nil {
			return partyId, err1
		}
		return partyId, err
	} else {
		err = o.Commit()
		return partyId, err
	}

}
func (p *PartyOperation) IsUserCreateParty() bool {
	if p.user.User().Id != p.srvParty.GetModelsParty().UserId {
		return false
	}
	return true
}

func NewPartyOperationForUpdate(partyId string, user *userSrv.User) (p *PartyOperation, err error) {
	p = &PartyOperation{
		user:           user,
		srvParty:       partySrv.NewSrvParty(),
		srvPartyDetail: partySrv.NewSrvPartyDetail(),
		photos:         new(photoSrv.SrvPhoto),
	}
	party, err := models.GetPartyByID(partyId)
	if err != nil {
		return
	}

	p.srvParty.SetModelsParty(party)
	partyIdInt, err := strconv.Atoi(partyId)
	if err != nil {
		return
	}
	partyDetail, err := models.GetPartyDetailByPartyId(uint32(partyIdInt))
	if err != nil {
		return
	}
	p.srvPartyDetail.SetModelsPartyDetail(partyDetail)
	return
}

func (p *PartyOperation) Edit() (err error) {
	if !p.IsUserCreateParty() {
		return errors.New("User can't update the party")
	}

	if p.srvParty.GetModelsParty().IsClosed() {
		return errors.New("PartyClosed!")
	}
	p.srvParty.GetModelsParty().ValidState = 0
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}

	_, err1 := o.Update(p.srvParty.GetModelsParty())

	_, err2 := o.Update(p.srvPartyDetail.GetModelsPartyDetail())
	if err1 != nil || err2 != nil {
		err := o.Rollback()
		if err != nil {
			return err
		}
		return errors.New("PartyEdit err")
	} else {
		err = o.Commit()
		return err
	}

}

func (p *PartyOperation) UpdateState(state uint8) (err error) {
	if state != 1 && state != 2 {
		return errors.New("Invalid state")
	}
	if p.srvParty.GetModelsParty().IsClosed() {
		return errors.New("Party Closed!")
	}
	if p.srvParty.GetModelsParty().IsChecked() {
		return errors.New("Party  Checked!")
	}

	p.srvParty.GetModelsParty().ValidState = state

	_, err = models.Orm.Update(p.srvParty.GetModelsParty())
	return err
}

func (p *PartyOperation) SetGameList() error {
	gameList, _, err := models.GetGameListByPID(strconv.Itoa(int(p.srvParty.GetModelsParty().Id)))
	if err != nil {
		return err
	}
	p.gameList = gameList

	return err

}

func (p *PartyOperation) Close() (err error) {
	if !p.IsUserCreateParty() {
		return errors.New("User  can't close the party")
	}
	if p.srvParty.GetModelsParty().IsClosed() {
		return errors.New("Party Closed!")
	}
	if !p.srvParty.GetModelsParty().IsChecked() {
		return errors.New("Party NO Checked!")
	}
	err = p.SetGameList()
	if err != nil {
		return err
	}
	var game_ids []uint32
	for _, game := range p.gameList {
		game_ids = append(game_ids, game.Id)
	}
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}

	_, err1 := o.QueryTable("party").Filter("id", p.GetSrvParty().GetModelsParty().Id).Update(orm.Params{
		"close_time": time.Now().Local(),
	})
	var err2 error
	if len(game_ids) != 0 {
		_, err2 = o.QueryTable("game").Filter("id__in", game_ids).Filter("close_time__isnull", true).Update(orm.Params{
			"close_time": time.Now().Local(),
		})
	}
	if err1 != nil || err2 != nil {
		err := o.Rollback()
		if err != nil {
			return err
		}
		return errors.New("PartyClose err")
	} else {
		err = o.Commit()
		return err
	}

}

func (p *PartyOperation) QueryRegResult() (games_reg_info []map[string]interface{}, err error) {

	for _, game := range p.gameList {
		game_reg_info := make(map[string]interface{})
		game_reg_info["game_id"] = game.Id
		game_reg_info["name"] = game.Name

		queryselector := models.Orm.QueryTable(models.DB_TABLE_REGISTRATION).Filter("game_id", game.Id)
		game_reg_info["wait_pay_player_count"], err = queryselector.Filter("pay_state", registrationSrv.ORDER_PAY_STATE_WAIT_PAY).Count()
		if err != nil {
			return games_reg_info, err
		}

		game_reg_info["payed_player_count"], err = queryselector.Filter("pay_state", registrationSrv.ORDER_PAY_STATE_PAYED).Count()
		if err != nil {
			return games_reg_info, err
		}

		games_reg_info = append(games_reg_info, game_reg_info)

	}
	return games_reg_info, err
}

func (p *PartyOperation) GetPhotoList() error {
	return p.photos.GetSrvPhotoByRelId(p.srvParty.GetModelsParty().Id)

}

func (p *PartyOperation) Detail() map[string]interface{} {
	partyDetail := make(map[string]interface{})
	partyDetail["Id"] = p.srvParty.GetModelsParty().Id
	partyDetail["UserId"] = p.srvParty.GetModelsParty().UserId

	partyDetail["Name"] = p.srvParty.GetModelsParty().Name

	partyDetail["Country"] = p.srvParty.GetModelsParty().Country

	partyDetail["Province"] = p.srvParty.GetModelsParty().Province

	partyDetail["City"] = p.srvParty.GetModelsParty().City

	partyDetail["Addr"] = p.srvParty.GetModelsParty().Addr

	partyDetail["LocLong"] = p.srvParty.GetModelsParty().LocLong
	partyDetail["LocLat"] = p.srvParty.GetModelsParty().LocLat

	partyDetail["RegStartTime"] = p.srvParty.GetModelsParty().RegStartTime

	partyDetail["RegEndTime"] = p.srvParty.GetModelsParty().RegEndTime

	partyDetail["StartTime"] = p.srvParty.GetModelsParty().StartTime

	partyDetail["EndTime"] = p.srvParty.GetModelsParty().EndTime

	partyDetail["CloseTime"] = p.srvParty.GetModelsParty().CloseTime

	partyDetail["Limitation"] = p.srvParty.GetModelsParty().Limitation

	partyDetail["LimitationType"] = p.srvParty.GetModelsParty().LimitationType

	partyDetail["Slogan"] = p.srvPartyDetail.GetModelsPartyDetail().Slogan
	partyDetail["Like"] = p.srvPartyDetail.GetModelsPartyDetail().Like
	partyDetail["Website"] = p.srvPartyDetail.GetModelsPartyDetail().Website

	partyDetail["Type"] = p.srvPartyDetail.GetModelsPartyDetail().Type

	partyDetail["Price"] = p.srvPartyDetail.GetModelsPartyDetail().Price
	partyDetail["Introduction"] = p.srvPartyDetail.GetModelsPartyDetail().Introduction
	partyDetail["Schedule"] = p.srvPartyDetail.GetModelsPartyDetail().Schedule

	partyDetail["Score"] = p.srvPartyDetail.GetModelsPartyDetail().Score
	partyDetail["SignupMale"] = p.srvPartyDetail.GetModelsPartyDetail().SignupMale
	partyDetail["SignupFemale"] = p.srvPartyDetail.GetModelsPartyDetail().SignupFemale

	return partyDetail
}

func (p *PartyOperation) PartyQueryGames() []map[string]interface{} {
	var gamesPrice []map[string]interface{}
	for _, game := range p.gameList {
		gamePrice := make(map[string]interface{})
		gamePrice["game_id"] = game.Id
		gamePrice["game_name"] = game.Name
		gamePrice["rmb_price"] = game.RmbPrice
		gamePrice["usd_price"] = game.UsdPrice

		gamesPrice = append(gamesPrice, gamePrice)
	}
	return gamesPrice
}
