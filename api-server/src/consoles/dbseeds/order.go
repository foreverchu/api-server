package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"time"

	"models"
)

func CreateOrders(num int) {
	orders := make([]*models.Order, 0, num)

	for i := 1; i <= num; i++ {
		iString := strconv.Itoa(i + 10)
		orderNo := fmt.Sprintf("%x", md5.Sum([]byte(iString)))

		order := &models.Order{
			UserId:       uint32(i),
			OrderNo:      orderNo,
			SubmitTime:   time.Now(),
			Price:        1,
			CurrencyType: 1,
			PayMethod:    2,
			Game:         &models.Game{Id: uint32(i)},
		}
		orders = append(orders, order)
	}

	successNum, err := models.Orm.InsertMulti(num, orders)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateOrders insert orders %d \n", successNum)
}
