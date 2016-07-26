package utils

import (
	"fmt"
	"time"
)

func GenerateOrderNo() (string, error) {
	now := time.Now()
	ns := now.UnixNano() % 1000000
	return fmt.Sprintf("%v%06d%04d", now.Format("20060102150405"),
		ns, Rander.Intn(10000))[2:], nil
}
