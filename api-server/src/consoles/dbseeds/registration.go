package main

import (
	"log"
	"models"
)

func CreateRegistration(num int) {
	regs := make([]*models.Registration, 0, num)
	for i := 1; i <= num; i++ {
		reg := &models.Registration{
			Game:     &models.Game{Id: uint32(i)},
			Order:    &models.Order{Id: uint32(i)},
			Player:   &models.Player{Id: uint32(i)},
			PayState: 0,
		}
		regs = append(regs, reg)
	}

	successNum, err := models.Orm.InsertMulti(num, regs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateRegistration insert %d regs\n", successNum)

}
