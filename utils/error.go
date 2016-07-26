package utils

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

const (
	HP_RESULT = "result"
	HP_MSG    = "msg"
)

func CheckResultJsonStr(res_str string) error {
	js, err := simplejson.NewJson([]byte(res_str))
	if err != nil {
		Logger.Error("Json parsing error %v, str: %s", err.Error(), res_str)
		return err
	}

	result, err := js.Get(HP_RESULT).Int()
	if err != nil {
		Logger.Error("no result code %v, str: %s", err.Error(), res_str)
		return err
	}

	if result != 0 {
		msg, err := js.Get(HP_MSG).String()
		if err != nil {
			Logger.Error("no msg, code: %v", result)
			return err
		}

		Logger.Error("api call failed: code: %v msg: %v", result, msg)
		return fmt.Errorf("api call failed: code: %v msg: %v", result, msg)
	}

	return nil
}
