package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	DB_PS_PLAYER_ID = "player_id"
	DB_PS_GAME_ID   = "game_id"
)

type PlayerScore struct {
	Id uint32

	Game       *Game   `orm:"rel(fk);rel(one)"`
	Player     *Player `orm:"rel(fk);rel(one)"`
	Result     int
	CreateTime time.Time
	UpdateTime time.Time
}

// GetPlayerScore 询某一选手在某一比赛, 是否已经登记过跑步成绩
// beego的orm只支持根据主键查询, 无法根据其它字段查询, 因此用rawsql查询对象
func GetPlayerScore(player *Player, game *Game) (*PlayerScore, error) {
	ps := PlayerScore{
		Game:   game,
		Player: player,
	}

	err := Orm.Raw("SELECT id, player_id, game_id, result, create_time, update_time FROM "+DB_TABLE_PLAYER_SCORE+" WHERE player_id = ? and game_id = ?", player.Id, game.Id).QueryRow(&ps)
	if err == orm.ErrNoRows {
		return nil, err
	}
	return &ps, nil
}

// IsUpdated 表示成绩是否更新过
func (ps *PlayerScore) IsUpdated() bool {
	return !ps.UpdateTime.IsZero()
}
