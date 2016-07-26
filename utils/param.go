package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

var (
	ErrInvalidData = errors.New("无效的数据")
)

// 数据类型常量
const (
	DATA_TYPE_STRING   = "string"
	DATA_TYPE_FLOAT64  = "float64"
	DATA_TYPE_INT      = "int"
	DATA_TYPE_UINT8    = "uint8"
	DATA_TYPE_DATE     = "date"
	DATA_TYPE_DATETIME = "datetime"
	DATA_TYPE_ARRAY    = "array"
	DATA_TYPE_UINT32   = "uint32"
	DATA_TYPE_FLOAT32  = "float32"
)

type ParamInfo struct {
	Req      bool
	Datatype string
}

func ParseParam(js *simplejson.Json, params_def map[string]ParamInfo) (map[string]interface{}, error) {
	params_value := make(map[string]interface{})
	var value interface{}
	var err error
	for param_name, param_info := range params_def {
		js_val, ok := js.CheckGet(param_name)
		if !ok {
			if param_info.Req {
				return nil, fmt.Errorf("%s is required", param_name)
			} else {
				continue
			}
		}
		switch param_info.Datatype {
		case DATA_TYPE_STRING:
			value, err = js_val.String()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
			value = strings.TrimSpace(value.(string))
			if value == "" {
				if param_info.Req {
					return nil, fmt.Errorf("%s is required", param_name)
				}
			}
		case DATA_TYPE_UINT32:
			value_uint64, err := js_val.Uint64()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
			value = uint32(value_uint64)
		case DATA_TYPE_INT:
			value, err = js_val.Int()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
		case DATA_TYPE_FLOAT64:
			value, err = js_val.Float64()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
		case DATA_TYPE_FLOAT32:
			value_float64, err := js_val.Float64()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
			value = float32(value_float64)
		case DATA_TYPE_UINT8:
			valueInt, err := js_val.Int()
			if err != nil {
				if param_info.Req {
					return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
				} else {
					return nil, fmt.Errorf("%s 参数类型不正确", param_name)
				}
			}
			if valueInt < 0 {
				return nil, fmt.Errorf("%s should great than or equal to 0", param_name)
			}

			if valueInt > 255 {
				return nil, fmt.Errorf("%s should be less than 256", param_name)
			}
			value = uint8(valueInt)

		case DATA_TYPE_DATE:
			dateStr, err := js_val.String()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
			value, err = ParseDateStr(dateStr, DATE_FORMAT_LAYOUT)
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
		case DATA_TYPE_DATETIME:
			datetimeStr, err := js_val.String()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
			value, err = ParseDateStr(datetimeStr, TIME_FORMAT_LAYOUT)
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
		case DATA_TYPE_ARRAY:
			value, err = js_val.Array()
			if err != nil {
				return nil, fmt.Errorf("Failed in paring %s. errror: %v", param_name, err)
			}
		default:
			return nil, fmt.Errorf("unknown type")
		}
		params_value[param_name] = value
	}

	return params_value, nil
}

func ParseDateStr(date_str string, time_layout string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	result_time, err := time.ParseInLocation(time_layout, date_str, loc)
	if err != nil {
		return result_time, fmt.Errorf("failed in parsing %s as %s to int: %v", date_str, time_layout, err.Error())
	}

	return result_time, nil
}

func GetUInt32FromParam(js *simplejson.Json, key string) (uint32, error) {
	str, err := js.Get(key).String()
	if err != nil {
		i_value, err := js.Get(key).Int()
		if err != nil {
			return 0, fmt.Errorf("failed in paring %s from json: %v", key, err.Error())
		}

		return uint32(i_value), nil
	}

	i_value, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed in convert %s to int: %v", str, err.Error())
	}

	return uint32(i_value), nil
}

func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func IsChina(country string) bool {
	lw_country := strings.ToLower(country)

	return strings.Contains(lw_country, "china") || strings.Contains(lw_country, "中国") || lw_country == "cn"
}

func CheckIdInArray(id uint32, ids []uint32) bool {
	for _, id_item := range ids {
		if id == id_item {
			return true
		}
	}

	return false
}

/*
seo param就是指在url中为了seo 而加入一
些关键词的参数， 例如 beijing_23
sep是用来分隔id和前面关键词的分隔符， 若没有找到这个sep，则会把整个参数返回
这个函数会把最后的参数编程uint返回
*/
func GetUintIdFromSeoParam(param string, sep string) (uint32, error) {
	str_id, err := GetStrIdFromSeoParam(param, sep)
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(str_id)
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

/*
seo param就是指在url中为了seo 而加入一些关键词的参数， 例如 beijing_23
sep是用来分隔id和前面关键词的分隔符， 若没有找到这个sep，则会把整个参数返回
这个函数会直接返回最后的字符串参数
*/
func GetStrIdFromSeoParam(param string, sep string) (string, error) {
	splitted_params := strings.Split(param, sep)

	if len(splitted_params) < 1 {
		return "", ErrInvalidData
	}

	return splitted_params[len(splitted_params)-1], nil
}

func GetEleStrFromJsonStr(json_str, ele_key string) (string, error) {
	var ele_str string

	js, err := simplejson.NewJson([]byte(json_str))
	if err != nil {
		return ele_str, err
	}

	ele_js := js.Get(ele_key)
	data, err := ele_js.Encode()
	if err != nil {
		return ele_str, err
	}

	return string(data[:]), nil
}
