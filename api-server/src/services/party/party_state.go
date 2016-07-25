package partySrv

import (
	"errors"
	"math"
	"models"
)

type PartyState int

const (
	NORMAL  PartyState = 0
	NOPARTY PartyState = -1
	CLOSED  PartyState = -2
)

var partyStateDesc = [...]string{
	"赛事正常",
	"无赛事",
	"赛事已关闭",
}

func (ps PartyState) String() string {
	i := int(math.Abs(float64(ps)))
	return partyStateDesc[i]
}

func CheckPartyState(p *models.Party) (PartyState, error) {
	var ps PartyState
	switch {
	case p == nil:
		ps = NOPARTY
	case p.IsClosed():
		ps = CLOSED
	case p != nil && !p.IsClosed():
		ps = NORMAL
	default:
		return PartyState(1), errors.New("获取赛事状态异常")
	}
	return ps, nil
}
