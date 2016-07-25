package controllers

import (
	"err_code"
	"fmt"
	md "models"
	gameSrv "services/game"
	"strconv"
	"time"

	"github.com/chinarun/utils"

	"github.com/astaxie/beego/context"
	"github.com/bitly/go-simplejson"
)

const (
	HP_GAME_ID       = "game_id"
	HP_NAME          = "name"
	HP_RMB_PRICE     = "rmb_price"
	HP_USD_PRICE     = "usd_price"
	HP_USER_ID       = "user_id"
	HP_GENDER_REQ    = "gender_req"
	HP_MIN_AGE_REQ   = "min_age_req"
	HP_MAX_AGE_REQ   = "max_age_req"
	HP_PAGE_NO       = "page_no"
	HP_PAGE_SIZE     = "page_size"
	HP_ORDER_BY      = "order_by"
	HP_INCLUDE_CLOSE = "include_close"
)

//party_state
const (
	PARTY_STATE_NORMAL       = 0
	PARTY_STATE_NOPARTY      = -1
	PARTY_STATE_PARTY_CLOSED = -2
)

// 定义性别类型
const (
	GENDER_NONE   = 0
	GENDER_MALE   = 1
	GENDER_FEMALE = 2
)

type game_info struct {
	Game_id   uint32
	Game_name string
}

//games相关接口实现
func getGameEditHttpParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_GAME_ID: {PARAM_REQUESTED, utils.DATA_TYPE_STRING},

		HP_NAME:        {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_LIMITATION:  {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_RMB_PRICE:   {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT64},
		HP_USD_PRICE:   {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT64},
		HP_START_TIME:  {PARAM_OPTIONAL, utils.DATA_TYPE_DATETIME},
		HP_END_TIME:    {PARAM_OPTIONAL, utils.DATA_TYPE_DATETIME},
		HP_GENDER_REQ:  {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_MIN_AGE_REQ: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_MAX_AGE_REQ: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_USER_ID:     {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
	}
}

func getGameCreateHttpParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_NAME:        {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_LIMITATION:  {PARAM_REQUESTED, utils.DATA_TYPE_INT},
		HP_PARTY_ID:    {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_RMB_PRICE:   {PARAM_REQUESTED, utils.DATA_TYPE_FLOAT64},
		HP_USD_PRICE:   {PARAM_REQUESTED, utils.DATA_TYPE_FLOAT64},
		HP_START_TIME:  {PARAM_REQUESTED, utils.DATA_TYPE_DATETIME},
		HP_END_TIME:    {PARAM_REQUESTED, utils.DATA_TYPE_DATETIME},
		HP_GENDER_REQ:  {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_MIN_AGE_REQ: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_MAX_AGE_REQ: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_USER_ID:     {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
	}
}

func CheckGameNameExist(party_id string, game_name string) (bool, error) {
	count, err := md.Orm.QueryTable(md.DB_TABLE_GAME).Filter(md.DB_GAME_FN_PARTY_ID, party_id).
		Filter(md.DB_GAME_FN_NAME, game_name).Count()
	if err != nil {
		utils.Logger.Error("failed in query table select count(*) from game where %s == %s and  %s == %s, err: %v",
			md.DB_GAME_FN_PARTY_ID, party_id, md.DB_GAME_FN_NAME, game_name, err)

		return false, err
	}

	return count > 0, nil

}

func setGameEditValueFromParam(game *md.Game, param_values map[string]interface{}) (int, error) {

	if param_values[HP_NAME] != nil {

		if game.Name != param_values[HP_NAME].(string) {

			party_id := strconv.Itoa(int(game.Party.Id))
			IsExist, err := CheckGameNameExist(party_id, param_values[HP_NAME].(string))
			if IsExist {
				return err_code.GameNameExists, fmt.Errorf("game_name exist")
			} else if err != nil {
				return err_code.ServerErr, fmt.Errorf("服务器查询出错")
			}
		}
		game.Name = param_values[HP_NAME].(string)
	}

	if param_values[HP_LIMITATION] != nil {
		if param_values[HP_LIMITATION].(int) < 0 {
			return err_code.InvalidData, fmt.Errorf("人数限制应该大于0")
		}
		// iValue := param_values[HP_LIMITATION].(int)
		game.Limitation = uint32(param_values[HP_LIMITATION].(int))
	}

	if param_values[HP_RMB_PRICE] != nil {
		if param_values[HP_RMB_PRICE].(float64) < 0 {
			return err_code.InvalidData, fmt.Errorf("赛事人民币价格不能小于0")
		}

		game.RmbPrice = int32(param_values[HP_RMB_PRICE].(float64) * 100)
	}

	if param_values[HP_USD_PRICE] != nil {
		if param_values[HP_USD_PRICE].(float64) < 0 {
			return err_code.InvalidData, fmt.Errorf("赛事美元价格不能小于0")
		}
		game.UsdPrice = int32(param_values[HP_USD_PRICE].(float64) * 100)

	}

	if param_values[HP_START_TIME] != nil {
		game.StartTime = param_values[HP_START_TIME].(time.Time)
	}

	if param_values[HP_END_TIME] != nil {
		game.EndTime = param_values[HP_END_TIME].(time.Time)

	}

	if game.StartTime.After(game.EndTime) {
		return err_code.InvalidData, fmt.Errorf("比赛开始时间不能大于比赛结束时间")
	}

	if param_values[HP_GENDER_REQ] != nil {
		game.GenderReq = param_values[HP_GENDER_REQ].(uint8)

		if game.GenderReq > GENDER_FEMALE {
			return err_code.InvalidData, fmt.Errorf("%s 应小于等于%d", HP_GENDER_REQ, GENDER_FEMALE)
		}
	}

	if param_values[HP_MIN_AGE_REQ] != nil {
		game.MinAgeReq = param_values[HP_MIN_AGE_REQ].(uint8)
	}
	if param_values[HP_MAX_AGE_REQ] != nil {
		game.MaxAgeReq = param_values[HP_MAX_AGE_REQ].(uint8)
	}
	if game.MinAgeReq > game.MaxAgeReq && game.MaxAgeReq != 0 {
		return err_code.InvalidData, fmt.Errorf("最小年龄应该大于最大年龄")
	}
	if game.MaxAgeReq > 100 {
		return err_code.InvalidData, fmt.Errorf("你输入的最大年龄不合理")
	}

	return err_code.OK, nil
}

func setGameValueFromParam(game *md.Game, param_values map[string]interface{}) (int, error) {
	if game.Party == nil {
		party_id, ok := param_values[HP_PARTY_ID].(string)
		if !ok {
			return err_code.InvalidPartyId, fmt.Errorf("%s is required", HP_PARTY_ID)
		}

		party, err := GetPartyByID(party_id)
		if err != nil {
			return err_code.InvalidPartyId, err
		}

		game.Party = party
	}

	game.Name = param_values[HP_NAME].(string)

	if param_values[HP_LIMITATION].(int) < 0 {
		return err_code.InvalidData, fmt.Errorf("人数限制应该大于0")
	}
	iValue := param_values[HP_LIMITATION].(int)
	game.Limitation = uint32(iValue)

	if param_values[HP_RMB_PRICE].(float64) < 0 {
		return err_code.InvalidData, fmt.Errorf("赛事人民币价格不能小于0")
	}
	game.RmbPrice = int32(param_values[HP_RMB_PRICE].(float64) * 100)
	if param_values[HP_USD_PRICE].(float64) < 0 {
		return err_code.InvalidData, fmt.Errorf("赛事美元价格不能小于0")
	}

	game.UsdPrice = int32(param_values[HP_USD_PRICE].(float64) * 100)

	game.StartTime = param_values[HP_START_TIME].(time.Time)

	game.EndTime = param_values[HP_END_TIME].(time.Time)
	if game.StartTime.After(game.EndTime) {
		return err_code.InvalidData, fmt.Errorf("比赛开始时间不能大于比赛结束时间")
	}

	if param_values[HP_GENDER_REQ] != nil {
		game.GenderReq = param_values[HP_GENDER_REQ].(uint8)

		if game.GenderReq > GENDER_FEMALE {
			return err_code.InvalidData, fmt.Errorf("%s 应小于等于%d", HP_GENDER_REQ, GENDER_FEMALE)
		}
	}

	if param_values[HP_MIN_AGE_REQ] != nil {
		game.MinAgeReq = param_values[HP_MIN_AGE_REQ].(uint8)
	}
	if param_values[HP_MAX_AGE_REQ] != nil {
		game.MaxAgeReq = param_values[HP_MAX_AGE_REQ].(uint8)
	}
	if game.MinAgeReq > game.MaxAgeReq && game.MaxAgeReq != 0 {
		return err_code.InvalidData, fmt.Errorf("最小年龄应该大于最大年龄")
	}

	return err_code.OK, nil
}

// 此方法在model层有对应的实现方法
func GetGameByID(game_id string) (*md.Game, error) {
	query_result := md.Orm.QueryTable(md.DB_TABLE_GAME).Filter(md.DB_ID, game_id)

	count, err := query_result.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("%s is not a valid game id", game_id)
	}

	game := new(md.Game)
	err = query_result.One(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func GameEdit(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	//parse input data
	param_values, err := utils.ParseParam(js, getGameEditHttpParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	game, err := GetGameByID(param_values[HP_GAME_ID].(string))
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, err.Error())
		return retJson
	}

	// if ok, _ := CheckCurUserEditPartyPermission(ctx, game.Party); !ok {
	// 	retJson_edit(&retJson, err_code.NoPermission, "")
	// 	return &retJson
	// }
	if !game.CloseTime.IsZero() {
		retJson_edit(retJson, err_code.NoPermission, "已经关闭的比赛不能再修改")
		return retJson
	}
	if time.Now().After(game.EndTime) {
		retJson_edit(retJson, err_code.NoPermission, "已经结束的比赛不能再进行编辑")
		return retJson
	}

	_, err = setGameEditValueFromParam(game, param_values)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	_, err = md.Orm.Update(game)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, err.Error())
		return retJson
	}

	RemoveGameFromCache(game.Id)
	RefreshGameLimitationMap()

	return retJson
}

//查看赛事是否存在，是否关闭
func check_party_state(party_id string) (party_state int, err error) {
	var party md.Party
	query_result := md.Orm.QueryTable(md.DB_TABLE_PARTY).Filter(md.DB_ID, party_id)
	count, err := query_result.Count()
	if err != nil {
		return
	}
	if count == 1 {
		err = query_result.One(&party)
		if err != nil {
			return
		}

		if party.CloseTime.IsZero() {
			party_state = PARTY_STATE_NORMAL
		} else {
			party_state = PARTY_STATE_PARTY_CLOSED
		}
	} else {
		party_state = PARTY_STATE_NOPARTY
	}
	return
}

func GameCreate(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	//parse input data
	param_values, err := utils.ParseParam(js, getGameCreateHttpParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	if param_values[HP_RMB_PRICE].(float64) < 0 || param_values[HP_USD_PRICE].(float64) < 0 {
		retJson_edit(retJson, err_code.PriceIsNegative, err.Error())
		return retJson
	}
	party_state, err := check_party_state(param_values[HP_PARTY_ID].(string))
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, err.Error())
		return retJson
	}

	if party_state != PARTY_STATE_NORMAL {
		if party_state == PARTY_STATE_NOPARTY {
			retJson_edit(retJson, err_code.InvalidPartyId, "")
		} else {
			retJson_edit(retJson, err_code.PartyClosed, "")
		}
		return retJson
	}
	var is_game_name_exist bool

	is_game_name_exist, err = CheckGameNameExist(param_values[HP_PARTY_ID].(string),
		param_values[HP_NAME].(string))
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, err.Error())
		return retJson
	}

	if is_game_name_exist {
		retJson_edit(retJson, err_code.PartyExist, "")
		return retJson
	}

	var game md.Game

	result_code, err := setGameValueFromParam(&game, param_values)
	if err != nil {
		retJson_edit(retJson, result_code, err.Error())
		return retJson
	}

	id, err := md.Orm.Insert(&game)
	if err != nil {
		utils.Logger.Error("数据库写入出错, %v", err)
		retJson_edit(retJson, err_code.ServerErr, "数据库写入出错,"+err.Error())
		return retJson
	}

	retJson["game_id"] = fmt.Sprintf("%d", id)

	RefreshGameLimitationMap()

	return retJson
}

//api game delete
func GameClose(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	game_id, err := js.Get("game_id").String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的比赛id")
		return retJson
	}

	game, _ := GetGameByID(game_id)
	if game == nil {
		retJson_edit(retJson, err_code.InvalidData, "无效的比赛id")
		return retJson
	}

	party := game.Party
	//check permission
	permission, err := CheckCurUserEditPartyPermission(ctx, party)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	if !permission {
		retJson_edit(retJson, err_code.NoPermission, "没有删除次比赛的权限")
		return retJson
	}

	game.CloseTime = time.Now()
	_, err = md.Orm.Update(game)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "关闭赛事失败: "+err.Error())
		return retJson
	}

	RemoveGameFromCache(game.Id)

	RefreshGameLimitationMap()

	return retJson
}

//api query game by party
func QueryPartyGames(js *simplejson.Json) map[string]interface{} {
	retJson := init_retJson()
	partyIdStr, err := js.Get("party_id").String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的赛事id")
		return retJson
	}
	partyId, err := strconv.Atoi(partyIdStr)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "party_id包含非数字字符")
		return retJson
	}
	uPartyId := uint32(partyId)
	party := md.Party{Id: uPartyId}
	err = md.Orm.Read(&party)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "无效的赛事id")
		return retJson
	}

	var games []md.Game
	queryselector := md.Orm.QueryTable(md.DB_TABLE_GAME)
	count, err := queryselector.Filter("party", party).All(&games)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "数据库查询出错")
		return retJson
	}
	var games_info []game_info
	for _, game := range games {
		games_info = append(games_info, game_info{game.Id, game.Name})
	}

	retJson["games"] = games_info
	retJson["count"] = count

	return retJson
}

func GameInputScore(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	game_id, err := js.Get(HP_GAME_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}
	player_id, err := js.Get(HP_PLAYER_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPlayerId, "")
		return retJson
	}
	//成绩
	result, err := js.Get("result").Int()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "比赛成绩出错")
		return retJson
	}

	//比赛状态
	game, err := md.GetGameByID(game_id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "获取比赛出错")
		return retJson
	}
	if game == nil {
		retJson_edit(retJson, err_code.InvalidData, "比赛不存在")
		return retJson
	}

	gs, err := gameSrv.CheckGameState(game)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	if gs == gameSrv.NOT_EXISTS || gs == gameSrv.NOT_START || gs == gameSrv.CLOSED {
		retJson_edit(retJson, err_code.InvalidData, fmt.Sprint(gs))
		return retJson
	}

	//选手状态
	player := md.NewPlayer()
	err = player.FindBy(map[string]interface{}{"Id": player_id})
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPlayerId, "")
		return retJson
	}
	if registered, err := is_player_registered(player, game); !registered {
		if err != nil {
			retJson_edit(retJson, err_code.InvalidData, err.Error())
		} else {
			retJson_edit(retJson, err_code.InvalidData, "选手未参加此比赛")
		}
		return retJson
	}

	//入库, 确保唯一性, db层添加约束or逻辑层?
	//更新or插入?
	var ps *md.PlayerScore
	ps, err = md.GetPlayerScore(player, game)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	//更新
	if ps != nil {
		ps.Result = result
		ps.UpdateTime = time.Now()
		_, err := md.Orm.Update(ps)
		if err != nil {
			utils.Logger.Error("数据库更新出错, %v", err)
			retJson_edit(retJson, err_code.ServerErr, "数据库更新出错,"+err.Error())
			return retJson
		}

	} else {
		ps = &md.PlayerScore{
			Game:       game,
			Player:     player,
			Result:     result,
			CreateTime: time.Now(),
		}

		_, err = md.Orm.Insert(ps)
		if err != nil {
			utils.Logger.Error("数据库写入出错, %v", err)
			retJson_edit(retJson, err_code.ServerErr, "数据库写入出错,"+err.Error())
			return retJson
		}
	}
	return retJson
}

func GameQueryRegInfo(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	game_id, err := js.Get(utils.HP_GAME_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}
	gq, err := gameSrv.NewGameQuery().FindByID(game_id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	regInfo, err := gq.RegisterInfo()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	retJson["apply"] = regInfo.ApplyCount()
	retJson["payed"] = regInfo.PayedCount()
	return retJson
}

func GameQuery(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	game_id, err := js.Get(utils.HP_GAME_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}
	gq, err := gameSrv.NewGameQuery().FindByID(game_id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	game, _ := gq.Game()
	retJson["game"] = game
	retJson["party_id"] = game.Party.Id

	return retJson
}

//判断选手是否报过名
func is_player_registered(player *md.Player, game *md.Game) (bool, error) {
	query_result := md.Orm.QueryTable(md.DB_TABLE_REGISTRATION).Filter(md.DB_REG_PLAYER_ID, player.Id).Filter(md.DB_REG_GAME_ID, game.Id)

	count, err := query_result.Count()
	if err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}
	return true, nil
}

func GameList(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	party_id, err := js.Get(utils.HP_PARTY_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}
	gq, err := gameSrv.NewGameQuery().FindListByPID(party_id)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	gameList, _ := gq.GameList()
	retJson["gameList"] = gameList
	retJson["count"] = len(gameList)

	return retJson
}
