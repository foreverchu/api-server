package main

import (
	"fmt"
	"log"
	"models"
)

func CreatePlayers(num int) {
	players := make([]*models.Player, 0, num)

	for i := 1; i <= num; i++ {
		randNum := random(1000, 9999)
		certNo := fmt.Sprintf("30122719881212%d", randNum)

		player := &models.Player{
			UserId:          uint32(i),
			Name:            fake.Name(),
			CertificateType: 1,
			CertificateNo:   certNo,
		}
		players = append(players, player)
	}

	successNum, err := models.Orm.InsertMulti(num, players)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("CreatePlayers insert players %d \n", successNum)
}
