package gameSrv

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"

	"models"
)

// GameQuery is an object to query a game model
// packed a lot of bussiness logic process
type GameQuery struct {
	game_id       string
	game_name     string
	party_id      string
	game          *models.Game
	party         *models.Party   //表示game所属的赛事
	players       *models.Players // 表示参加此比赛的所有选手的集合
	registrations []*models.Registration
	gameList      []*models.Game //表示game列表
}

func NewGameQuery() *GameQuery {
	return &GameQuery{}
}

// FindByID 根据game_id来查找game model
// 获取一个orm对象统一使用Findxxx
func (gq *GameQuery) FindByID(game_id string) (*GameQuery, error) {
	game_model, err := models.GetGameByID(game_id)
	if err != nil {
		return nil, err
	}
	gq.game_id = game_id
	gq.game = game_model

	return gq, nil
}

// FindByName 根据name来查找game model
func (gq *GameQuery) FindByName(party_id string, name string) (*GameQuery, error) {
	game_model, err := models.GetGameByName(party_id, name)
	if err != nil {
		return gq, err
	}
	gq.party_id = party_id
	gq.game_name = name
	gq.game = game_model
	return gq, nil
}

// Game 用于获取GameQuery的orm对象
func (gq *GameQuery) Game() (*models.Game, error) {
	if gq.game == nil {
		return nil, errors.New("services/game: no game exists")
	}
	return gq.game, nil
}

// Party 用于获取game所于的party的orm对象
func (gq *GameQuery) Party() (*models.Party, error) {
	if gq.party == nil {
		party_id := fmt.Sprintf("%d", gq.game.Party.Id)
		party, err := models.GetPartyByID(party_id)
		if err != nil {
			return nil, err
		}
		gq.party_id = party_id
		gq.party = party
	}

	return gq.party, nil
}

// Players 用于获取此比赛的报名的models.Players
func (gq *GameQuery) Players() (*models.Players, error) {
	if gq.players == nil {
		players, err := models.GetPlayersByGameId(gq.game_id)
		if err != nil {
			return nil, err
		}
		gq.players = players
	}
	return gq.players, nil
}

// RegisterInfo 用于获取一个比赛的报名选手信息
// btw: 直接通过sql完成统计, 而不需要查出对象然后再loop
func (gq *GameQuery) RegisterInfo() (*GameRegInfo, error) {
	regInfo := &GameRegInfo{}

	game_id := fmt.Sprintf("%d", gq.game.Id)
	registrations, err := models.GetRegistrationByGameID(game_id)
	if err != nil {
		return nil, err
	}

	for _, reg := range registrations {
		if reg.IsWaiting() {
			regInfo.applyCount += 1
			continue
		}
		if reg.IsPayed() {
			regInfo.payedCount += 1
			continue
		}
	}

	return regInfo, nil
}

// FindByID 根据game_id来查找game model
// 获取一个orm对象统一使用Findxxx
func (gq *GameQuery) FindListByPID(party_id string) (*GameQuery, error) {
	gameList, _, err := models.GetGameListByPID(party_id)
	if err != nil {
		return nil, err
	}
	gq.gameList = gameList

	return gq, nil
}
func (gq *GameQuery) CloseGameList() error {
	var game_ids []uint32
	for _, game := range gq.gameList {
		game_ids = append(game_ids, game.Id)
	}
	_, err := models.Orm.QueryTable("game").Filter("id__in", game_ids).Filter("close_time__isnull").Update(orm.Params{
		"CloseTime": time.Now().Local(),
	})
	return err
}

// Game 用于获取GameQuery的orm对象
func (gq *GameQuery) GameList() ([]*models.Game, error) {
	if gq.gameList == nil {
		return nil, errors.New("gameSrv.GameList: no game exists")
	}
	return gq.gameList, nil
}
