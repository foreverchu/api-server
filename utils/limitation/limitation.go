package limitation

import (
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/chinarun/utils"
)

const (
	PARAM_REQUESTED = true
	PARAM_OPTIONAL  = false
)

const (
	HP_GAME_ID    = "game_id"
	HP_LIMITATION = "limitation"
	HP_BALANCE    = "balance"
)

type BalanceInfo struct {
	GameId     string `json:"game_id"`
	Limitation int    `json:"limitation"`
	Balance    int    `json:"balance"`
}

func getGameBalanceParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_GAME_ID:    {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_LIMITATION: {PARAM_REQUESTED, utils.DATA_TYPE_INT},
		HP_BALANCE:    {PARAM_REQUESTED, utils.DATA_TYPE_INT},
	}
}

func ParseBalanceInfoArrayJson(js *simplejson.Json, balances_key string) (*utils.SafeMap, error) {
	js_val, ok := js.CheckGet(balances_key)
	if !ok {
		utils.Logger.Error("no %s in json %v", balances_key, js)
		return nil, fmt.Errorf("no %s in json %v", balances_key, js)
	}

	balances, err := js_val.Array()
	if err != nil {
		return nil, fmt.Errorf("Failed in paring json. error: %v", err.Error())
	}

	game_limitation_map := utils.NewSafeMap()
	for i, _ := range balances {
		params, err := utils.ParseParam(js_val.GetIndex(i), getGameBalanceParametersMap())
		if err != nil {
			return nil, fmt.Errorf("Failed in paring game balance %v. errror: %v", js, err.Error())
		}

		game_limitation_map.Set(params[HP_GAME_ID].(string),
			BalanceInfo{
				GameId:     params[HP_GAME_ID].(string),
				Limitation: params[HP_LIMITATION].(int),
				Balance:    params[HP_BALANCE].(int),
			})
	}

	return game_limitation_map, nil
}
