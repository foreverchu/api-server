package utils

import (
	"strconv"
	"encoding/json"
)

func AddInt(v ... interface{})(sum int){
	for _, item := range v {
		switch val := item.(type) {
		case int8:
			sum += int(val)
		case uint8:
			sum += int(val)
		case int32:
			sum += int(val)
		case uint32:
			sum += int(val)
		case uint:
			sum += int(val)
		case int64:
			sum += int(val)
		case uint64:
			sum += int(val)
		case float32:
			sum += int(val)
		case float64:
			sum += int(val)
		case int:
			sum += val
		case string:
			num, err := strconv.Atoi(val)
			if err == nil {
				sum += num;
			}
		case json.Number:
			num, err := val.Int64()
			if err == nil {
				sum += int(num);
			}
		}
	}

	return
}