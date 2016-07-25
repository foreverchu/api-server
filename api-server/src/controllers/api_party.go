package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	"github.com/chinarun/utils"

	"err_code"
	md "models"
	"services/parse_params"
	"services/party"
	"services/party_operation"
	"services/registration"
)

const (
	HP_PARTY_NAME      = "name"
	HP_COUNTRY         = "country"
	HP_PROVINCE        = "province"
	HP_CITY            = "city"
	HP_TAGS            = "tags"
	HP_MONTH           = "month"
	HP_VALID_STATE     = "valid_state"
	HP_ADDR            = "addr"
	HP_LOC_LONG        = "loc_long"
	HP_LOC_LAT         = "loc_lat"
	HP_LIMITATION      = "limitation"
	HP_LIMITATION_TYPE = "limitation_type"
	HP_REG_START_TIME  = "reg_start_time"
	HP_REG_END_TIME    = "reg_end_time"
	HP_START_TIME      = "start_time"
	HP_END_TIME        = "end_time"
	HP_PARTY_ID        = "party_id"
	HP_PARTY_STATE     = "state"

	// Party detail
	HP_SLOGAN        = "slogan"
	HP_LIKE          = "like"
	HP_WEBSITE       = "website"
	HP_TYPE          = "type"
	HP_INTRODUCTION  = "introduction"
	HP_SCHEDULE      = "schedule"
	HP_SCORE         = "score"
	HP_SIGNUP_MALE   = "signup_male"
	HP_SIGNUP_FEMALE = "signup_female"
	HP_PRICE         = "price"
)

const (
	DEF_PARTY_PAGE_SIZE       = 10
	DEF_PARTY_PAGE_NO         = 0
	DEF_PARTY_MONTH           = 0
	DEF_ORDER_QUERY_PAGE_SIZE = 20
	DEF_ORDER_QUERY_PAGE_NO   = 0
)

type reg_info struct {
	order_no      string
	game_name     string
	pay_status    uint8
	price         float32
	currency_type uint
}

type GameRegInfo struct {
	Game_id               uint32 `json:"game_id"`
	Name                  string `json:"name"`
	Payed_player_count    int64  `json:"payed_player_count"`
	Wait_pay_player_count int64  `json:"wait_pay_player_count"`
}

type QueryOrderPlayerInfo struct {
	Id              uint32 `json:"id"`
	Name            string `json:"name"`
	Gender          uint8  `json:"gender"`
	CertificateType uint8  `json:"certificate_type"`
	CertificateNo   string `json:"certificate_no"`
}

type QueryOrderInfo struct {
	GameId       uint32    `json:"game_id"`
	GameName     string    `json:"game_name"`
	OrderNo      string    `json:"order_no"`
	SubmitTime   time.Time `json:"submit_time"`
	PayTime      time.Time `json:"pay_time"`
	RefundTime   time.Time `json:"refund_time"`
	CancelTime   time.Time `json:"cancel_time"`
	Price        float32   `json:"price"`
	CurrencyType uint8     `json:"currency_type"`
	PayMethod    uint8     `json:"pay_method"`
	PayAccount   string    `json:"pay_account"`
	UserId       string    `json:"user_id"`

	Players []QueryOrderPlayerInfo `json:"players"`
}

type QueryGameInfo struct {
	GameId   string  `json:"game_id"`
	GameName string  `json:"game_name"`
	RmbPrice float32 `json:"rmb_price"`
	UsdPrice float32 `json:"usd_price"`
}

//partys 相关操作接口实现

func getPartyQueryRegOrdersHttpParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_PARTY_ID:         {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_ORDER_NO:         {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_CERTIFICATE_TYPE: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_CERTIFICATE_NO:   {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_PAGE_NO:          {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_PAGE_SIZE:        {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_PAY_STATUS:       {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
	}
}

//api party create
func (c *PartyController) PartyCreate() map[string]interface{} {

	retJson := init_retJson()
	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyCreate : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyCreate : %#v", c.js)
		}
	}()

	if err = c.IsLogin(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	//parse input data
	param_values, err := utils.ParseParam(c.js, parseParamsSrv.GetPartyCreateHttpParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson

	}

	if partySrv.IsPartyNameExist(param_values[parseParamsSrv.HP_PARTY_NAME].(string)) {
		err = errors.New("Party Exist")
		retJson_edit(retJson, err_code.PartyExist, "")
		return retJson
	}

	party_id, err := partyOperationSrv.NewPartyOperationForCreate(c.CurrentUser.User().Id).SetParamsForCreate(param_values).Create()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "数据库插入出错")
		return retJson

	}

	retJson["party_id"] = party_id
	return retJson
}
func GetPartyGames(party_id uint32) (*[]md.Game, error) {
	var games []md.Game
	queryselector := md.Orm.QueryTable(md.DB_TABLE_GAME)
	_, err := queryselector.Filter("party_id", party_id).All(&games)
	if err != nil {
		return nil, err
	}

	return &games, nil
}

func GetPartyByID(party_id string) (*md.Party, error) {
	query_result := md.Orm.QueryTable(md.DB_TABLE_PARTY).Filter(md.DB_ID, party_id)

	count, err := query_result.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("%s is not a valid party id", party_id)
	}

	party := new(md.Party)
	err = query_result.One(party)
	if err != nil {
		return nil, err
	}

	return party, nil
}

func CheckCurUserEditPartyPermission(ctx *context.Context, party *md.Party) (bool, error) {
	return true, nil
}

//api party update
func (c *PartyController) PartyEdit() map[string]interface{} {
	retJson := init_retJson()
	var err error

	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyEdit : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyEdit : %#v", c.js)
		}
	}()
	if err = c.IsLogin(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	param_values, err := utils.ParseParam(c.js, parseParamsSrv.GetPartyEditHttpParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	partyOperation, err := partyOperationSrv.NewPartyOperationForUpdate(param_values[parseParamsSrv.HP_PARTY_ID].(string), c.CurrentUser)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPartyId, "")
		return retJson
	}

	err = partyOperation.SetParamsForUpdate(param_values).Edit()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "")
		return retJson
	}
	return retJson
}

//api party shutdown
func (c *PartyController) PartyClose() map[string]interface{} {
	//todo: check user permission
	retJson := init_retJson()
	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyClose : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyClose : %#v", c.js)
		}
	}()

	if err = c.IsLogin(); err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	party_id, err := c.js.Get("party_id").String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的party_id")
		return retJson
	}

	partyOperation, err := partyOperationSrv.NewPartyOperationForUpdate(party_id, c.CurrentUser)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	err = partyOperation.Close()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	return retJson
}

//api PartyStateUpdate
func (c *PartyController) PartyStateUpdate() map[string]interface{} {
	retJson := init_retJson()

	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyStateUpdate : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyStateUpdate : %#v", c.js)
		}
	}()
	if !c.BaseController.Admin {
		err = errors.New("Not Admin")
		retJson_edit(retJson, err_code.NoPermission, err.Error())
		return retJson
	}

	paramsValues, err := utils.ParseParam(c.js, parseParamsSrv.GetPartyStateUpdateParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	partyOperation, err := partyOperationSrv.NewPartyOperationForUpdate(paramsValues["party_id"].(string), c.CurrentUser)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	err = partyOperation.UpdateState(paramsValues["state"].(uint8))
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	return retJson
}

//api party list
func (c *PartyController) PartyList() map[string]interface{} {
	retJson := init_retJson()

	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyList : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyList : %#v", c.js)
		}
	}()

	param_values, err := utils.ParseParam(c.js, parseParamsSrv.GetPartyListParamsMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	var page_size int
	if param_values[HP_PAGE_SIZE] == nil {
		page_size = DEF_PARTY_PAGE_SIZE
	} else {
		page_size = param_values[HP_PAGE_SIZE].(int)
		if page_size <= 0 {
			page_size = DEF_PARTY_PAGE_SIZE
		}
	}

	var page_no int
	if param_values[HP_PAGE_NO] != nil {
		page_no = param_values[HP_PAGE_NO].(int)
		if page_no < 0 {
			page_no = DEF_PARTY_PAGE_NO
		}
	}

	var month int
	if param_values[HP_MONTH] != nil {
		month = param_values[HP_MONTH].(int)
		if month >= 13 || month <= 0 {
			month = DEF_PARTY_MONTH
		}
	}
	var country string
	if param_values[HP_COUNTRY] != nil {
		country = param_values[HP_COUNTRY].(string)
		if country == "全部" {
			country = ""
		}
	}

	var province string
	if param_values[HP_PROVINCE] != nil {
		province = param_values[HP_PROVINCE].(string)
		if province == "全部" {
			province = ""
		}
	}

	var city string
	if param_values[HP_CITY] != nil {
		city = param_values[HP_CITY].(string)
		if city == "全部" {
			city = ""
		}
	}

	var tags_string string
	if param_values[HP_TAGS] != nil {
		tags_string = param_values[HP_TAGS].(string)
		if tags_string == "全部比赛" {
			tags_string = ""
		}
	}

	var valid_state int
	if param_values[HP_VALID_STATE] != nil {
		valid_state = param_values[HP_VALID_STATE].(int)
		if valid_state < partySrv.DEF_VALID_STATE_WAIT || valid_state > partySrv.DEF_VALID_STATE_ALL {
			valid_state = partySrv.DEF_VALID_STATE_PASS
		}
	}

	var order_field string
	if param_values[HP_ORDER_BY] == nil {
		order_field = md.DB_PARTY_FN_START_TIME
	} else {
		order_field = param_values[HP_ORDER_BY].(string)
	}

	var include_close int
	if param_values[HP_INCLUDE_CLOSE] != nil {
		include_close = param_values[HP_INCLUDE_CLOSE].(int)
	}

	var include_close_bool = false
	if include_close == 1 {
		include_close_bool = true
	}

	plq := partySrv.NewPartyListQuery().PageNo(page_no).PageSize(page_size).
		Country(country).Province(province).City(city).Tags(tags_string).ValidState(valid_state).Month(month).
		IncludeClose(include_close_bool)

	if param_values[HP_USER_ID] != nil {
		user_id := param_values[HP_USER_ID].(string)

		var party_type int
		if param_values[HP_TYPE] != nil {
			party_type = param_values[HP_TYPE].(int)
		}

		plq.SetUserPartyType(user_id, party_type)
	}

	order := make(map[string]bool)
	if order_field != "" {
		order[order_field] = false //true表示asc;false表示desc;默认desc
	}
	plq.OrderBy(order)

	err = plq.Do()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}
	result := plq.Result()
	var count int64
	count = plq.Total()
	retJson["partylist"] = result
	retJson["count"] = count
	return retJson
}

//api party query
func (c *PartyController) PartyQuery() map[string]interface{} {
	retJson := init_retJson()
	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyQuery : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyQuery : %#v", c.js)
		}
	}()

	party_id, err := c.js.Get(HP_PARTY_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的party_id\n"+err.Error())
		return retJson
	}
	partyOperation, err := partyOperationSrv.NewPartyOperationForUpdate(party_id, c.CurrentUser)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}
	if partyOperation.GetPhotoList() != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}
	retJson["party"] = partyOperation.Detail()

	retJson["photos"] = partyOperation.GetPhoto().GetModelsPhotoList()
	return retJson
}

//api QueryRegStateByCert

func QueryRegStateByCert(js *simplejson.Json) map[string]interface{} {
	retJson := init_retJson()
	cert_type_int, err := js.Get(HP_CERTIFICATE_TYPE).Int()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "证件类型参数不正确"+err.Error())
		return retJson
	}
	cert_no, err := js.Get(HP_CERTIFICATE_NO).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "证件号码参数不正确"+err.Error())
		return retJson
	}
	var player md.Player
	queryselector := md.Orm.QueryTable(md.DB_TABLE_PLAYER)
	err = queryselector.Filter("certificatetype", uint8(cert_type_int)).Filter("certificateno", cert_no).One(&player)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "数据库查询失败")
		return retJson
	}

	party_id_str, err := js.Get(HP_PARTY_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "缺少party_id字段或者类型不正确"+err.Error())
		return retJson
	}
	party_id, err := strconv.Atoi(party_id_str)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPartyId, "")
		return retJson
	}

	var registrations []md.Registration
	queryselector = md.Orm.QueryTable(md.DB_TABLE_REGISTRATION)
	_, err = queryselector.Filter("player", player.Id).All(&registrations)

	games, err := GetPartyGames(uint32(party_id))
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPartyId, "")
		return retJson
	}
	games_in_party := make(map[uint32]md.Game)
	for _, game := range *games {
		if game.Party.Id == uint32(party_id) {
			games_in_party[game.Id] = game
		}
	}

	var info []reg_info
	for _, registration := range registrations {
		order, err := GetOrderByID(registration.Order.Id)
		if err != nil {
			retJson_edit(retJson, err_code.InvalidOrderId, "获取注册订单信息失败")
			return retJson
		}
		var game md.Game
		price := float32(order.Price / 100.0)
		game, ok := games_in_party[registration.Game.Id]
		if !ok {
			utils.Logger.Error("Invalid Game id : %v in registration, registration_id: %v", registration.Game.Id, registration.Id)
			retJson_edit(retJson, err_code.ServerErr, "发现无效的game_id")
			return retJson
		}

		info = append(info, reg_info{order.OrderNo, game.Name, registration.PayState, price, order.CurrencyType})
	}
	retJson["order_info"] = info

	return retJson
}

//api PartyQueryRegResult
func (c *PartyController) PartyQueryRegResult() map[string]interface{} {
	retJson := init_retJson()

	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyQueryRegResult : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyQueryRegResult : %#v", c.js)
		}
	}()

	party_id, err := c.js.Get(HP_PARTY_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的party_id\n"+err.Error())
		return retJson
	}

	partyOperation, err := partyOperationSrv.NewPartyOperationForUpdate(party_id, c.CurrentUser)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPartyId, "")
		return retJson
	}

	err = partyOperation.SetGameList()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}
	games_reg_info, err := partyOperation.QueryRegResult()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	retJson["games"] = games_reg_info

	return retJson

}

//返回第一个参数是错误码
func assignQueryOrderInfo(order *md.Order, query_order_info *QueryOrderInfo, game_name string) (int, error) {
	query_order_info.GameId = order.Game.Id
	query_order_info.GameName = game_name
	query_order_info.OrderNo = order.OrderNo
	query_order_info.SubmitTime = order.SubmitTime
	query_order_info.RefundTime = order.RefundTime
	query_order_info.CancelTime = order.CancelTime
	query_order_info.Price = float32(order.Price) / 100.0
	query_order_info.CurrencyType = uint8(order.CurrencyType)
	query_order_info.PayMethod = uint8(order.PayMethod)
	query_order_info.PayAccount = order.PayAccount
	query_order_info.UserId = strconv.Itoa(int(order.UserId))

	var err error

	query_order_info.Players, err = getQueryOrderPlayerInfo(order.Id)
	if err != nil {
		if len(query_order_info.Players) == 0 {
			return err_code.InvalidData, err
		} else {
			return err_code.ServerErr, err
		}
	}

	return err_code.OK, nil
}

func PartyQueryRegOrders(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()

	if !utils.IsClientFromSamePrivateNetwork(ctx) {
		retJson_edit(retJson, err_code.UserNoLogin, "请先登录")
		return retJson
	}

	param_values, err := utils.ParseParam(js, getPartyQueryRegOrdersHttpParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	party_id, err := strconv.Atoi(param_values[HP_PARTY_ID].(string))
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPartyId, err.Error())
		return retJson
	}

	games, err := GetPartyGames(uint32(party_id))
	if err != nil || len(*games) <= 0 {
		retJson_edit(retJson, err_code.InvalidPartyId, "无效party_id："+err.Error())
		return retJson
	}

	game_ids := make([]uint32, len(*games))
	id_game_map := make(map[uint32]*md.Game)

	for i, game := range *games {
		game_ids[i] = game.Id
		id_game_map[game.Id] = &game
	}

	var query_orders []QueryOrderInfo
	var order *md.Order

	if param_values[HP_ORDER_NO] != nil {
		order, err = GetOrderByNo(param_values[HP_ORDER_NO].(string))
		if err != nil {
			retJson_edit(retJson, err_code.InvalidData, "无效订单号："+err.Error())
			return retJson
		}

		if !utils.CheckIdInArray(order.Game.Id, game_ids) {
			retJson_edit(retJson, err_code.InvalidData, "订单号不属于该赛事")
			return retJson
		}

		query_orders = make([]QueryOrderInfo, 1)

		result, err := assignQueryOrderInfo(order, &query_orders[0], id_game_map[order.Game.Id].Name)
		if result != err_code.OK {
			retJson_edit(retJson, result, err.Error())
			return retJson
		}
	} else {
		if param_values[HP_CERTIFICATE_TYPE] != nil && param_values[HP_CERTIFICATE_NO] != nil {
			player, err := GetPlayerByCertificate(param_values[HP_CERTIFICATE_TYPE].(uint8),
				param_values[HP_CERTIFICATE_NO].(string))
			if err != nil {
				retJson_edit(retJson, err_code.InvalidData, "无效证件信息："+err.Error())
				return retJson
			}

			var registration md.Registration

			queryselector := md.Orm.QueryTable(md.DB_TABLE_REGISTRATION)
			queryselector = queryselector.Filter("player_id", player.Id)
			queryselector.Filter("game_id__in", game_ids)
			err = queryselector.One(&registration)
			if err == orm.ErrNoRows {
				retJson_edit(retJson, err_code.PlayerRegNotExist, "")
				return retJson
			} else if err != nil {
				retJson_edit(retJson, err_code.ServerErr, "数据库查询出错")
				return retJson
			}

			order, err := GetOrderByID(registration.Order.Id)
			if err != nil {
				retJson_edit(retJson, err_code.ServerErr, err.Error())
				return retJson
			}

			query_orders = make([]QueryOrderInfo, 1)

			result, err := assignQueryOrderInfo(order, &query_orders[0], id_game_map[order.Game.Id].Name)
			if result != err_code.OK {
				retJson_edit(retJson, result, err.Error())
				return retJson
			}
		} else {
			//搜索所有赛事订单
			pay_status := registrationSrv.ORDER_PAY_STATE_NONE
			page_no := 0
			page_size := DEF_ORDER_QUERY_PAGE_SIZE

			if param_values[HP_PAGE_NO] != nil {
				page_no = param_values[HP_PAGE_NO].(int)
			}
			if param_values[HP_PAGE_SIZE] != nil {
				page_size = param_values[HP_PAGE_SIZE].(int)
			}
			if param_values[HP_PAY_STATUS] != nil {
				pay_status = param_values[HP_PAY_STATUS].(int)
				if pay_status < registrationSrv.ORDER_PAY_STATE_FIRST || pay_status > registrationSrv.ORDER_PAY_STATE_LAST {
					retJson_edit(retJson, err_code.InvalidData, fmt.Sprintf("无效的%s", HP_PAY_STATUS))
					return retJson
				}
			}

			queryselector := md.Orm.QueryTable(md.DB_TABLE_ORDER)
			queryselector = queryselector.Filter("game_id__in", game_ids)
			if pay_status == registrationSrv.ORDER_PAY_STATE_WAIT_PAY {
				queryselector = queryselector.Filter("pay_time__ispage", true)
			} else if pay_status == registrationSrv.ORDER_PAY_STATE_PAYED {
				queryselector = queryselector.Filter("pay_time__isnull", false).
					Filter("refund_time__isnull", false).Filter("cancel_time__isnull", false)
			} else if pay_status == registrationSrv.ORDER_PAY_STATE_CANCELED {
				queryselector = queryselector.Filter("cancel_time__isnull", true)
			} else if pay_status == registrationSrv.ORDER_PAY_STATE_REFUNDED {
				queryselector = queryselector.Filter("refund_time__isnull", true)
			}

			queryselector = queryselector.Limit(page_size).Offset(page_no * page_size).
				OrderBy("-" + md.DB_ORDER_SUBMIT_TIME)

			var orders []md.Order

			count, err := queryselector.All(&orders)

			if err == orm.ErrNoRows {
				retJson_edit(retJson, err_code.InvalidData, "没有符合条件的订单")
				return retJson
			} else if err != nil {
				retJson_edit(retJson, err_code.ServerErr, "数据库查询出错")
				return retJson
			}

			query_orders = make([]QueryOrderInfo, count)
			for i, order := range orders {
				result, err := assignQueryOrderInfo(&order, &query_orders[i], id_game_map[order.Game.Id].Name)
				if result != err_code.OK {
					retJson_edit(retJson, result, err.Error())
					return retJson
				}
			}
		}
	}

	retJson["orders"] = query_orders

	return retJson
}

func getQueryOrderPlayerInfo(order_id uint32) ([]QueryOrderPlayerInfo, error) {
	var players_info []QueryOrderPlayerInfo

	queryselector := md.Orm.QueryTable(md.DB_TABLE_REGISTRATION)
	queryselector = queryselector.Filter("order_id", order_id)

	var registrations []md.Registration

	count, err := queryselector.All(&registrations)
	if err != nil {
		return players_info, err
	}

	if count <= 0 {
		utils.Logger.Error("无效的订单id: %d", order_id)
		return players_info, fmt.Errorf("无效的订单id: %d", order_id)
	}

	players_info = make([]QueryOrderPlayerInfo, count)
	for i, reg := range registrations {
		players_info[i].Id = reg.Player.Id

		player := md.NewPlayer()
		err := player.FindBy(map[string]interface{}{"Id": reg.Player.Id})
		if err != nil {
			return players_info, err
		}

		players_info[i].Name = player.Name
		players_info[i].Gender = player.Gender
		players_info[i].CertificateType = player.CertificateType
		players_info[i].CertificateNo = player.CertificateNo
	}

	return players_info, nil
}

//api query game by party
func (c *PartyController) PartyQueryGames() map[string]interface{} {
	retJson := init_retJson()

	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("controllers.PartyQueryRegResult : %s", err.Error())
		} else {
			utils.Logger.Debug("controllers.PartyQueryRegResult : %#v", c.js)
		}
	}()

	party_id, err := c.js.Get(HP_PARTY_ID).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的party_id\n"+err.Error())
		return retJson
	}

	partyOperation, err := partyOperationSrv.NewPartyOperationForUpdate(party_id, c.CurrentUser)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidPartyId, "")
		return retJson
	}

	err = partyOperation.SetGameList()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	games_reg_info := partyOperation.PartyQueryGames()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "")
		return retJson
	}

	retJson["games"] = games_reg_info
	retJson["count"] = len(games_reg_info)

	return retJson
}
