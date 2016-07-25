package playerSrv

import (
	"errors"
	"models"
)

var (
	ErrPlayerNotFound = errors.New("选手不存在")
)

type Player struct {
	player *models.Player
}

func NewPlayer() *Player {
	return new(Player)
}

func (p *Player) IsExists(column string, value interface{}) bool {
	return models.IsValueExists(models.DB_TABLE_PLAYER, column, value)
}
