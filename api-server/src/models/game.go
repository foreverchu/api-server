package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
)

var (
	ErrGameNotFound = errors.New("mdoels.Game.FindBy : game not found")
)

const (
	DB_GAME_FN_NAME     = "name"
	DB_GAME_FN_PARTY_ID = "party_id"
)

type Game struct {
	Id         uint32
	Name       string `orm:"size(60)"`
	Limitation uint32

	RmbPrice int32
	UsdPrice int32

	GenderReq uint8
	MinAgeReq uint8
	MaxAgeReq uint8

	StartTime time.Time `orm:"type(datetime)"`
	EndTime   time.Time `orm:"type(datetime)"`
	CloseTime time.Time `orm:"type(datetime)"`

	Party *Party `orm:"rel(fk)" json:"-"`
}

// GetGameByID 通过id获取game model
func GetGameByID(game_id string) (*Game, error) {
	game_id_int, err := strconv.Atoi(game_id)
	if err != nil {
		return nil, err
	}

	game := &Game{
		Id: uint32(game_id_int),
	}

	err = Orm.Read(game)
	if err == orm.ErrNoRows {
		return nil, err
	}
	return game, nil
}

// GetGameByName 通过name获取game model
func GetGameByName(party_id, game_name string) (*Game, error) {
	game := &Game{}
	err := Orm.Raw("SELECT * FROM "+DB_TABLE_GAME+" WHERE player_id = ? and game_name = ?", party_id, game_name).QueryRow(game)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	return game, nil
}

func (g *Game) FindBy(conditions map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Debug("models.Game.FindBy : error : %s, cond = %v", err.Error(), conditions)
		}
	}()
	qs := Orm.QueryTable(g)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err = qs.One(g)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrGameNotFound
	}
	return nil
}

// IsStarted 表示是否已经开始
func (g *Game) IsStarted() bool {
	return g.StartTime.Before(time.Now())
}

// IsHappening 表示是否正在进行
func (g *Game) IsHappening() bool {
	return g.StartTime.Before(time.Now()) && !g.IsEnded() && !g.IsClosed()
}

// IsEnded 表示是否已经结束
func (g *Game) IsEnded() bool {
	return g.EndTime.Before(time.Now())
}

// IsClosed 表示是否已经关闭
func (g *Game) IsClosed() bool {
	//如果字段为null, 表示没有设置, 没有关闭
	if g.CloseTime.IsZero() {
		return false
	}

	return g.CloseTime.Before(time.Now())
}

// 通过party_id获取game list
func GetGameListByPID(party_id string) (gameList []*Game, num int64, err error) {
	num, err = Orm.Raw("SELECT * FROM "+DB_TABLE_GAME+" WHERE party_id = ?", party_id).QueryRows(&gameList)

	return
}
