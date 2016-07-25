package messageSrv

import (
	"fmt"
	md "models"
	"services/party"
	"strconv"
	"time"

	"github.com/chinarun/utils"
)

const (
	checkDuration            = time.Duration(3 * time.Hour) //3小时执行一次party 查询
	partyNoticeInAdvanceTime = time.Duration(48 * time.Hour)
)

//轮循函数
func MsgDeamon() {
	go func() {
		for {
			rangeParty()
			time.Sleep(checkDuration)
		}
	}()
}

//party message 产生和写入
func rangeParty() {
	parties := partySrv.NewPartyListQuery().IncludeClose(false).Result()
	now := time.Now()
	for _, party := range parties {
		if party.RegStartTime.Sub(now) < partyNoticeInAdvanceTime {
			var msgType uint8
			msgType = utils.MSG_PARTY_REG
			generatePartyMsg(msgType, party)
		}
		if party.StartTime.Sub(now) < partyNoticeInAdvanceTime {
			var msgType uint8
			msgType = utils.MSG_PARTY_START
			generatePartyMsg(msgType, party)
		}
	}
}

func generatePartyMsg(msgType uint8, party *md.PartyAndDetail) {
	userParties, err := md.GetFollowUserPartiesByPartyId(party.PartyId)
	if err != nil {
		utils.Logger.Error("Get party followers err in generatePartyMsg func")
		return
	}
	var msg md.Message
	var msgs []md.Message
	msg.FromTable = utils.DB_TABLE_PARTY
	msg.FromTableId = party.PartyId
	msg.State = utils.MSG_UNREAD
	for _, userParty := range *userParties {
		msg.UserId = userParty.UserId
		switch msgType {
		case utils.MSG_PARTY_REG:
			msg.MessageId = generateMsgId(party.PartyId, userParty.UserId, msgType)
		case utils.MSG_PARTY_START:
			msg.MessageId = generateMsgId(party.PartyId, userParty.UserId, msgType)
		default:
			utils.Logger.Error("Unknown msg type in generatePartyMsg func, msgType:" + strconv.Itoa(int(msgType)))
		}
		msgs = append(msgs, msg)
	}
	md.InsertMsgs(msgs)
}

//simple message sender 此类消息只会生成一次,所以不用轮循,比如订单
//订单消息,订单生成的时候调用此函数产生一次消息
func OrderMsg(order *md.Order, msgType int) {
	var msg md.Message
	msg.State = utils.MSG_UNREAD
	msg.UserId = order.UserId
	msg.FromTable = utils.DB_TABLE_ORDER
	msg.FromTableId = order.Id
	switch msgType {
	case utils.MSG_ORDER_PAID:
		msg.MsgType = utils.MSG_ORDER_PAID
		msg.MessageId = generateMsgId(order.UserId, order.Id, utils.MSG_ORDER_PAID)
	case utils.MSG_ORDER_REFUND:
		msg.MsgType = utils.MSG_ORDER_REFUND
		msg.MessageId = generateMsgId(order.UserId, order.Id, utils.MSG_ORDER_REFUND)
	default:
		utils.Logger.Error("未知的订单类型参数,请检查msgType: orderId: %s msgType: %s", order.Id, msgType)
		return
	}
	md.InsertMsg(msg)
}

//用户关注消息,在用户关注某人时调用次函数产生一次消息
func FollowMsg(subject *md.User, object *md.User) {
	//subject是产生消息主体,即关注者, object是受体,即被关注者
	var msg md.Message
	msg.State = utils.MSG_UNREAD
	msg.UserId = object.Id
	msg.FromTable = utils.DB_TABLE_USER
	msg.MsgType = utils.MSG_USER_FOLLOW
	msg.FromTableId = subject.Id
	msg.MessageId = generateMsgId(subject.Id, object.Id, utils.MSG_USER_FOLLOW)
	md.InsertMsg(msg)
}

func generateMsgId(id1 uint32, id2 uint32, msgType uint8) (msgId string) {
	msgIdStr := fmt.Sprintf("%10d%10d%3d", id1, id2, msgType)
	msgId = utils.GetMd5Sha1(msgIdStr)
	return msgId
}

func Delete(msgId uint64, uid uint32) (err error) {
	err = md.DeleteMsg(msgId, uid)
	return err
}

func GetMsg(msgId uint64, uid uint32) (msg *md.Message, err error) {
	msg, err = md.GetMsgById(msgId, uid)
	return msg, err
}

func List(uid uint32) (msgs *[]md.Message, err error) {
	msgs, err = md.GetUserMsgs(uid)
	return msgs, err
}
