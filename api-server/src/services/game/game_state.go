package gameSrv

import (
	"errors"
	"models"
)

//比赛状态
type GameState int

const (
	NOT_EXISTS GameState = iota
	NOT_START
	HAPPENING
	ENDED
	CLOSED
)

var gameStateDesc = [...]string{
	"比赛不存在",
	"比赛还未开始",
	"比赛正在进行",
	"比赛已经结束",
	"比赛已经关闭",
}

func (gs GameState) String() string {
	return gameStateDesc[gs]
}

// CheckGameState 查看比赛的状态
func CheckGameState(g *models.Game) (GameState, error) {
	var gs GameState
	switch {
	case g == nil:
		gs = NOT_EXISTS
	case !g.IsStarted():
		gs = NOT_START
	case g.IsStarted() && g.IsHappening():
		gs = HAPPENING
	case g.IsEnded() && !g.IsClosed():
		gs = ENDED
	case g.IsClosed():
		gs = CLOSED
	default:
		return GameState(-1), errors.New("获取比赛状态异常")
	}
	return gs, nil
}
