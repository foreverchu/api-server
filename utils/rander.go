package utils

import (
	"math/rand"
	"time"
)

var (
	Rander *rand.Rand
)

func InitRander() {
	Rander = rand.New(rand.NewSource(time.Now().UnixNano()))
}
