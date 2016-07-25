package main

import (
	"fmt"
	"log"
	"time"

	"models"
)

func CreateParty(num int) {
	parties := make([]*models.Party, 0, num)

	cities := []string{
		"上海", "北京", "深圳",
	}

	for i := 1; i <= num; i++ {
		party := &models.Party{
			Name:           fmt.Sprintf("赛事%d", i),
			UserId:         uint32(i),
			Limitation:     100,
			LimitationType: 0,
			Country:        fake.Country(),
			Province:       "上海",
			City:           cities[i%len(cities)],
			Addr:           fake.StreetAddress(),
			LocLat:         float32(fake.Latitude()),
			LocLong:        float32(fake.Longitude()),
			RegStartTime:   time.Now().AddDate(0, 0, 10), //最早
			RegEndTime:     time.Now().AddDate(0, 0, 12), //次早
			StartTime:      time.Now().AddDate(0, 0, 15), //比赛开始时间要晚于注册开始时间,
			EndTime:        time.Now().AddDate(0, 0, 20),
			ValidState:     1,
		}

		parties = append(parties, party)
	}

	successNum, err := models.Orm.InsertMulti(num, parties)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateParty insert %d parties\n", successNum)

}

func CreatePartyDetail(num int) {
	parties := make([]*models.PartyDetail, 0, num)

	for i := 0; i < num; i++ {
		party := &models.PartyDetail{
			PartyId:      uint32(i + 1),
			Slogan:       "风靡全球 ，最好玩，美女帅哥老外最多的5公里跑步" + fake.Name(),
			Like:         uint32(i + 1),
			Website:      "www.chinarun.com",
			Type:         "5公里",
			Price:        "125元/人",
			Introduction: "风靡全球 ，最好玩，美女帅哥老外最多的5公里跑步风靡全球 ，最好玩，美女帅哥老外最多的5公里跑步",
			Schedule:     "09月25日　08：00 公开组男子5公里出发",
			Score:        4.6,
			SignupMale:   uint32(fake.Rand.Intn(65536)),
			SignupFemale: uint32(fake.Rand.Intn(65536)),
		}

		parties = append(parties, party)
	}

	successNum, err := models.Orm.InsertMulti(num, parties)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreatePartyDetail insert %d partieDetails\n", successNum)

}
