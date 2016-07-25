package partyOperationSrv

import (
	"errors"
	"models"
	"services/party"
	"strconv"

	"github.com/chinarun/utils"
)

var (
	ErrPartyNotExists = errors.New("赛事不存在")
	ErrFollowParty    = errors.New("关注赛事失败")
)

func FollowParty(userId, partyId string) (err error) {
	var actualErr error

	defer func() {
		if err != nil {
			utils.Logger.Error("userSrv.NewFollow : %s", actualErr.Error())
		} else {
			utils.Logger.Debug("userSrv.NewFollow: success : user_id = %s, party_id = %s", userId, partyId)
		}
	}()

	if ok := partySrv.IsPartyExists(partyId); !ok {
		return ErrPartyNotExists
	}
	partyIdInt, _ := strconv.Atoi(partyId)
	userIdInt, _ := strconv.Atoi(userId)

	userParty := &models.UserParty{
		UserId:  uint32(userIdInt),
		PartyId: uint32(partyIdInt),
	}
	if err = userParty.Follow(); err != nil {
		actualErr = err
		return ErrFollowParty
	}
	return nil
}
