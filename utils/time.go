package utils

import (
	"time"
)

// 获得当前时间: 毫秒 1s=1000ms
func GetNowTimeInMs() int64 {
	return time.Now().UnixNano() / 1000000
}

func GetNowTimeInSec() int64 {
	return GetNowTimeInMs() / 1000
}
