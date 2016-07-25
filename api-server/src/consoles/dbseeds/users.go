package main

import (
	"fmt"
	"log"
	"time"

	"models"
)

func CreateUsers(num int) {
	users := make([]*models.User, 0, num)
	for i := 1; i <= num; i++ {
		phone := fmt.Sprintf("1810163516%d", i-1)

		user := &models.User{
			Name:     fake.Name(),
			Phone:    phone,
			Email:    fake.Email(),
			Password: "abd865a104aff8c8e840ad4a1c73aa53", //123456
			Salt:     "QnrEOF",
			Avatar:   "",
			ComeFrom: 1,
			Active:   1,
		}
		users = append(users, user)
	}
	successNum, err := models.Orm.InsertMulti(num, users)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateUsers insert %d users\n", successNum)
}

func CreateProfiles(num int) {
	profiles := make([]*models.Profile, 0, num)
	for i := 1; i <= num; i++ {
		profile := &models.Profile{
			UserId:    uint32(i),
			CreatedAt: time.Now(),
		}
		profiles = append(profiles, profile)
	}
	successNum, err := models.Orm.InsertMulti(num, profiles)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateProfiles insert %d profiels\n", successNum)

}
