package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/chinarun/utils"
)

type Message struct {
	Id          uint64
	MessageId   string
	UserId      uint32
	FromTable   string
	FromTableId uint32
	MsgType     uint8
	State       uint8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//插入多条Message数据
func InsertMsgs(msgs []Message) {
	if _, err := Orm.InsertMulti(len(msgs), msgs); err != nil {
		utils.Logger.Error("Insert multi msgs err:" + err.Error())
	}
}

//插入单条Message数据
func InsertMsg(msg Message) {
	if _, _, err := Orm.ReadOrCreate(msg, msg.MessageId); err != nil {
		utils.Logger.Error("Insert msg err:" + err.Error())
	}
}

func GetFollowUserPartiesByPartyId(partyId uint32) (userPartys *[]UserParty, err error) {
	_, err = Orm.QueryTable(utils.DB_TABLE_USER_PARTY).Filter(strconv.Itoa(int(partyId))).All(&userPartys)
	if err != nil {
		utils.Logger.Error("failed in query users following a party, partyid: %d", partyId)
		return
	}
	return
}

func DeleteMsg(msgId uint64, uid uint32) error {
	msg := Message{Id: msgId}
	err := Orm.Read(&msg)
	if err != nil {
		return errors.New("消息不存在")
	}
	if msg.UserId != uid {
		return errors.New("消息不存在") //实际错误为此消息不属于此用户
	}
	msg.State = utils.MSG_DELETED
	_, err = Orm.Update(&msg)
	if err != nil {
		return errors.New("删除消息失败,请重试")
	}
	return nil

	return nil
}

func GetMsgById(msgId uint64, uid uint32) (msg *Message, err error) {
	msg = &Message{Id: msgId}
	if Orm.Read(msg) == nil {
		if msg.State == utils.MSG_DELETED || msg.UserId != uid {
			return nil, errors.New("消息不存在")
		}
		return msg, nil
	}
	return nil, errors.New("消息不存在")
}

func GetUserMsgs(uid uint32) (msgs *[]Message, err error) {
	var messages []Message
	_, err = Orm.QueryTable(utils.DB_TABLE_MSG).Filter(utils.MSG_USER_ID, uid).Exclude("state", utils.MSG_DELETED).All(&messages)
	if err != nil {
		return nil, err
	}
	return &messages, nil
}
