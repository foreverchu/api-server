package main

import (
	"fmt"
	"log"
	"models"
	"time"
)

func CreateGames(num int) {
	games := make([]*models.Game, 0, num)

	for i := 1; i <= num; i++ {
		game := &models.Game{
			Name:       fmt.Sprintf("%d公里跑", i),
			Limitation: 10,
			RmbPrice:   1,
			UsdPrice:   0,
			GenderReq:  0,
			MinAgeReq:  0,
			MaxAgeReq:  0,
			StartTime:  time.Now().AddDate(0, 0, 1),
			EndTime:    time.Now().AddDate(0, 0, 2),
			CloseTime:  time.Now().AddDate(0, 0, 2),
			Party:      &models.Party{Id: uint32(i)},
		}
		games = append(games, game)
	}

	successNum, err := models.Orm.InsertMulti(num, games)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateGames insert games %d \n", successNum)
}
