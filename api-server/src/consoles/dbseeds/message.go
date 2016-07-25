package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chinarun/utils"

	"models"
)

func CreateMessages(num int) {
	messages := make([]*models.Message, 0, num)
	msgtype := []uint8{1, 2, 3, 4, 5, 6}
	state := []uint8{1, 2, 3}
	for i := 1; i <= num; i++ {
		msg := &models.Message{
			UserId:      uint32(i),
			FromTable:   "user",
			FromTableId: uint32(i),
			MsgType:     msgtype[3],
			State:       state[0],
			CreatedAt:   time.Now(),
		}
		msgIdStr := fmt.Sprintf("%10d%10d%3d", msg.FromTableId, msg.UserId, msg.MsgType)
		msg.MessageId = utils.GetMd5Sha1(msgIdStr)

		messages = append(messages, msg)
	}
	successNum, err := models.Orm.InsertMulti(num, messages)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CreateMessages insert %d messages", successNum)
}
