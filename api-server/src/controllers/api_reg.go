package controllers

import (
	"err_code"
	"fmt"
	md "models"
	"services/registration"
	"strings"
	"sync/atomic"
	"time"

	"github.com/chinarun/utils"
	"github.com/chinarun/utils/limitation"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	simplejson "github.com/bitly/go-simplejson"
	guid "github.com/daltoniam/goguid"
)

const (
	HP_PLAYER_ID                = "player_id"
	HP_PLAYER_NAME              = "name"
	HP_CERTIFICATE_TYPE         = "certificate_type"
	HP_CERTIFICATE_NO           = "certificate_no"
	HP_MOBILE                   = "mobile"
	HP_REG_EMAIL                = "email"
	HP_ADDR1                    = "address1"
	HP_ADDR2                    = "address2"
	HP_ZIPCODE                  = "zip"
	HP_GENDER                   = "gender"
	HP_BIRTH_DATE               = "birth_date"
	HP_BLOOD_TYPE               = "blood_type"
	HP_PLAYER_HEIGHT            = "height"
	HP_PLAYER_WEIGHT            = "weight"
	HP_EMERGENCY_CONTACT_NAME   = "emergency_contact_name"
	HP_EMERGENCY_CONTACT_MOBILE = "emergency_contact_mobile"
	HP_T_SHIRT_SIZE             = "t_shirt_size"
	HP_INDUSTRY                 = "industry"
	HP_JOB_LEVEL                = "job_level"
	HP_INCOME                   = "income"
	HP_FAMILY_STATUS            = "family_status"
	HP_TOKEN                    = "token"
	HP_QUERY_INTERVAL           = "query_interval"
	HP_COUNT                    = "count"
	HP_QUEUE_STATE              = "queue_state"
	HP_PLAYERS                  = "players"
	HP_ORDER_NO                 = "order_no"
	HP_BALANCE                  = "balance"
	HP_NOTICE_TYPE              = "notice_type"
	HP_NOTICE_CONTENT           = "content"
	HP_NOTICE_RECIPIENT         = "recipient"
	HP_PAY_STATUS               = "pay_status"
	HP_BALANCES                 = "balances"
)

const (
	REG_ROUTINE_COUNT          = 3 //这个参数应该可以通过外部接口来调整
	REG_CHANNEL_CAPACITY       = 1024
	REG_EXPIRE_INTERVAL        = 300 //
	REG_MAX_PLAYER_IN_ONEORDER = 100
)

const (
	CURRENCY_TYPE_NONE = 0
	CURRENCY_TYPE_RMB  = 1
	CURRENCY_TYPE_USD  = 2
)

//报名排队状态, 0：需要继续等待,1：排队成功,请根据给定reg_token报名, 2：报名名额已满
const (
	REG_QUEUE_STATE_QUEUING        = 0
	REG_QUEUE_STATE_PROCCESSED     = 1
	REG_QUEUE_STATE_FULL           = 2
	REG_QUEUE_STATE_ERROR          = 3 //处理过程遇到错误
	REG_QUEUE_STATE_ALREADY_SIGNUP = 4 //已经报名
)

// 定义证件类型
const (
	CERT_TYPE_NONE                   = -1
	CERT_TYPE_ID_CARD                = 0 //身份证
	CERT_TYPE_PASSPORT               = 1 //护照
	CERT_TYPE_MILITARY_ID_CARD       = 2 //军官证
	CERT_TYPE_MTP                    = 3 //台胞证
	CERT_TYPE_EEP_HK_MACAU           = 4 //港澳通行证
	CERT_TYPE_HK_MACAU_RETURN_PERMIT = 5 //港澳回乡证
)

type RegItem struct {
	token            string
	user_id          uint32
	players          []*md.Player
	game             *md.Game
	order_no         string
	last_access_time int64 //记录这个数据被最后一次更新或最后一次查询的时间, 单位秒
	queue_idx        uint64
	queue_state      uint32 //排队结果
	msg              string //出错信息会放在这里
}

type GameLimitation struct {
	limitation int32 //名额限量
	balance    int32 //名额余额
}

var (
	g_reg_queue              chan RegItem //用户报名用channel
	g_reg_idx                uint64
	g_cur_processing_reg_idx uint64
	g_processed_map          *utils.SafeMap
	g_queued_reg_item_map    *utils.SafeMap //等待处理的 token -> *RegItem
	g_game_limitation_map    *utils.SafeMap //game id -> *GameLimitation, 每个比赛的名额限制
)

func InitRegRoutine() {
	g_reg_queue = make(chan RegItem, REG_CHANNEL_CAPACITY)
	g_processed_map = utils.NewSafeMap()
	g_queued_reg_item_map = utils.NewSafeMap()
	g_game_limitation_map = utils.NewSafeMap()
	fillGameLimitationMapFromDB()

	for i := 0; i < REG_ROUTINE_COUNT; i++ {
		go func() {
			reg_item := <-g_reg_queue
			signUp(&reg_item)
		}()
	}

	dumpRegData()
}

func dumpRegData() {
	time.AfterFunc(60*time.Second, func() {
		utils.Logger.Debug("报名队列: %s", utils.Sdump(g_queued_reg_item_map))
		utils.Logger.Debug("处理队列: %s", utils.Sdump(g_processed_map))
		utils.Logger.Debug("比赛队列: %s", utils.Sdump(g_game_limitation_map))
		utils.Logger.Debug("g_cur_processing_reg_ids: %s", utils.Sdump(g_cur_processing_reg_idx))
		utils.Logger.Debug("g_reg_idx: %s", utils.Sdump(g_reg_idx))
		utils.Logger.Debug("g_reg_queue: %s", utils.Sdump(g_reg_queue))
		dumpRegData()
	})
}

func RefreshGameLimitationMap() {
	fillGameLimitationMapFromDB()
}

//这是暂时的方法，以后要考中控服务器来控制比赛名额
//这个方法以后可用在中控服务器
func fillGameLimitationMapFromDB() error {
	g_game_limitation_map = utils.NewSafeMap()

	queryselector := md.Orm.QueryTable(md.DB_TABLE_GAME)
	queryselector = queryselector.Filter("start_time__gte", time.Now())

	var games []md.Game

	count, err := queryselector.All(&games)
	if err == orm.ErrNoRows {
		return nil
	} else if err != nil {
		utils.Logger.Error("failed in query game data: %v", err.Error())
		return err
	}

	for i := 0; i < int(count); i++ {
		balance := games[i].Limitation

		reg_count, err := md.Orm.QueryTable(md.DB_TABLE_REGISTRATION).Filter("game_id", games[i].Id).Count()
		if err != nil && err != orm.ErrNoRows {
			utils.Logger.Error("failed in query order data: game_id ", games[i].Id)
			continue
		}
		balance = games[i].Limitation - uint32(reg_count)

		g_game_limitation_map.Set(games[i].Id, &GameLimitation{limitation: int32(games[i].Limitation), balance: int32(balance)})
	}

	return nil
}

func GetOrderByID(id uint32) (*md.Order, error) {
	queryselector := md.Orm.QueryTable(md.DB_TABLE_ORDER).Filter("id", id)

	count, err := queryselector.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("%s is not a valid order id", id)
	}

	order := new(md.Order)
	err = queryselector.One(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetPlayerByCertificate(certificate_type uint8, certificate_no string) (*md.Player, error) {
	queryselector := md.Orm.QueryTable(md.DB_TABLE_PLAYER).
		Filter("certificate_type", certificate_type).
		Filter("certificate_no", certificate_no)

	count, err := queryselector.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("证件信息：%d %v 在数据库内不存在", certificate_type, certificate_no)
	}

	player := new(md.Player)
	err = queryselector.One(player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func GetOrderByNo(order_no string) (*md.Order, error) {
	queryselector := md.Orm.QueryTable(md.DB_TABLE_ORDER).Filter("order_no", order_no)

	count, err := queryselector.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("%s is not a valid order no", order_no)
	}

	order := new(md.Order)
	err = queryselector.One(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetOrderRegistrations(order_id uint32) (*[]md.Registration, error) {
	var registrations []md.Registration
	queryselector := md.Orm.QueryTable(md.DB_TABLE_REGISTRATION)
	_, err := queryselector.Filter("order_id", order_id).All(&registrations)
	if err != nil {
		return nil, err
	}

	return &registrations, nil

}

func setOrderRegistrationPayState(order_id uint32, pay_state uint8) error {
	update_orm := orm.NewOrm()
	update_orm.Begin()

	registrations, err := GetOrderRegistrations(order_id)
	if err != nil {
		utils.Logger.Error("failed in querying order registration: %v", err)
		return fmt.Errorf("failed in querying order registration: %v", err)
	}

	for _, reg := range *registrations {
		reg.PayState = pay_state
		_, err = update_orm.Update(reg)
		if err != nil {
			utils.Logger.Error("failed in updating order registration pay state: %v", err)
			return fmt.Errorf("failed in updating order registration pay state: %v", err)
		}
	}

	err = update_orm.Commit()
	if err != nil {
		utils.Logger.Error("数据库事务提交出错: %v", err.Error())

		err_rollback := update_orm.Rollback()
		if err_rollback != nil {
			utils.Logger.Critical("setOrderRegistrationPayState: 数据库回滚失败 order: %d ", order_id)
			return err
		}

		return err
	}

	return nil
}

func insertRegistrationData(signup_orm orm.Ormer, order *md.Order, players *[]*md.Player) error {

	for _, player := range *players {

		registration := new(md.Registration)
		registration.Order = order
		registration.Game = order.Game
		registration.Player = player

		_, err := signup_orm.Insert(registration)
		if err != nil {
			utils.Logger.Error("数据库写入出错: %v", err.Error())
			return err
		}
	}

	return nil
}

//报名, 采用事务
func signUp(item *RegItem) error {
	g_queued_reg_item_map.Delete(item.token)
	g_processed_map.Set(item.token, item)

	err := checkIfPlayersAlreadyRegGameByCertificate(&item.players, item.game.Id)
	if err != nil {
		item.msg = err.Error()
		item.queue_state = REG_QUEUE_STATE_ALREADY_SIGNUP
		return err
	}

	signup_orm := orm.NewOrm()
	signup_orm.Begin()

	err = updatePlayersDB(signup_orm, &item.players)
	if err != nil {
		utils.Logger.Error("failed in call updatePlayerDB: %v", err.Error())
		return err
	}

	item.last_access_time = utils.GetNowTimeInSec()

	if item.queue_idx > g_cur_processing_reg_idx {
		g_cur_processing_reg_idx = item.queue_idx
	}

	game_limitation, _ := g_game_limitation_map.Get(item.game.Id).(*GameLimitation)
	if game_limitation == nil || (game_limitation.balance <= 0 && game_limitation.limitation > 0) {
		item.queue_state = REG_QUEUE_STATE_FULL
		item.msg = "比赛名额已满"
		return fmt.Errorf("比赛名额已满 %d", item.game.Id)
	}

	game_limitation.balance = atomic.AddInt32(&game_limitation.balance, -1)

	order := new(md.Order)
	order.SubmitTime = time.Now()
	order.Game = item.game
	order.UserId = item.user_id
	order.CurrencyType = CURRENCY_TYPE_RMB
	order.Price = uint32(item.game.RmbPrice)
	order.OrderNo, err = utils.GenerateOrderNo()
	if err != nil {
		utils.Logger.Error("生成订单id出错: %v", err.Error())
		item.msg = fmt.Sprintf("生成订单id出错: %v", err.Error())
		item.queue_state = REG_QUEUE_STATE_ERROR
		return err
	}

	_, err = signup_orm.Insert(order)
	if err != nil {
		utils.Logger.Error("数据库写入出错: %v", err.Error())
		item.msg = fmt.Sprintf("数据库写入出错: %v", err.Error())
		item.queue_state = REG_QUEUE_STATE_ERROR
		return err
	}

	err = insertRegistrationData(signup_orm, order, &item.players)
	if err != nil {
		utils.Logger.Error("数据库写入出错: %v", err.Error())
		item.msg = fmt.Sprintf("数据库写入出错: %v", err.Error())
		item.queue_state = REG_QUEUE_STATE_ERROR
		return err
	}

	err = signup_orm.Commit()
	if err != nil {
		utils.Logger.Error("数据库事务提交出错: %v", err.Error())
		item.msg = fmt.Sprintf("数据库事务提交出错: %v", err.Error())
		item.queue_state = REG_QUEUE_STATE_ERROR

		err_rollback := signup_orm.Rollback()
		if err_rollback != nil {
			item.msg = fmt.Sprintf("数据库回滚失败 order: %v game: %v", order.OrderNo, item.game.Id)
			utils.Logger.Critical("%v", item.msg)
			item.queue_state = REG_QUEUE_STATE_ERROR
			return err
		}

		return err
	}

	postSignupSuccess(item)

	item.order_no = order.OrderNo
	item.queue_state = REG_QUEUE_STATE_PROCCESSED

	return nil
}

func postSignupSuccess(item *RegItem) {
	notifyPlayersSignupSuccess(item)
}

func notifyPlayersSignupSuccess(item *RegItem) {
	//这里应该用短信和email通知选手报名成功，催促他尽快支付
}

func QueryInterval() (t int) {
	queue_len := len(g_reg_queue)

	if queue_len < 10 {
		return 1
	} else if queue_len < 100 {
		return 3
	} else {
		return 5
	}
}

func GetPlayerInfoParameterMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{
		HP_NAME:                     {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_CERTIFICATE_TYPE:         {PARAM_REQUESTED, utils.DATA_TYPE_UINT8},
		HP_CERTIFICATE_NO:           {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_MOBILE:                   {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_REG_EMAIL:                {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_COUNTRY:                  {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_PROVINCE:                 {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_CITY:                     {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ADDR1:                    {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ADDR2:                    {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ZIPCODE:                  {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_GENDER:                   {PARAM_REQUESTED, utils.DATA_TYPE_UINT8},
		HP_BIRTH_DATE:               {PARAM_REQUESTED, utils.DATA_TYPE_DATE},
		HP_EMERGENCY_CONTACT_NAME:   {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_EMERGENCY_CONTACT_MOBILE: {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_T_SHIRT_SIZE:             {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
	}
}

func ParsePlayerInfoFromJson(player *md.Player, param_values map[string]interface{}) (err error) {
	player.Name = param_values[HP_NAME].(string)
	player.CertificateNo = param_values[HP_CERTIFICATE_NO].(string)
	player.Mobile = param_values[HP_MOBILE].(string)
	player.Email = param_values[HP_REG_EMAIL].(string)
	player.Country = param_values[HP_COUNTRY].(string)
	player.Province = param_values[HP_PROVINCE].(string)
	player.City = param_values[HP_CITY].(string)
	player.Address1 = param_values[HP_ADDR1].(string)
	player.Address2 = param_values[HP_ADDR2].(string)
	player.Zip = param_values[HP_ZIPCODE].(string)
	player.EmergencyContactName = param_values[HP_EMERGENCY_CONTACT_NAME].(string)
	player.EmergencyContactMobile = param_values[HP_EMERGENCY_CONTACT_MOBILE].(string)

	player.CertificateType = param_values[HP_CERTIFICATE_TYPE].(uint8)

	player.Gender = param_values[HP_GENDER].(uint8)

	player.BirthDate = param_values[HP_BIRTH_DATE].(time.Time)

	err = validatePlayerData(player)
	if err != nil {
		return err
	}
	player.TShirtSize = param_values[HP_T_SHIRT_SIZE].(uint8)

	return nil
}

func validatePlayerData(player *md.Player) error {

	return nil
}

func GetPalyerByCertificate(cert_type uint8, cert_no string) (*md.Player, error) {
	query_result := md.Orm.QueryTable(md.DB_TABLE_PLAYER).
		Filter(md.DB_PLAYER_CERTIFICATE_TYPE, cert_type).Filter(md.DB_PLAYER_CERTIFICATE_NO, cert_no)

	count, err := query_result.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("%v %v is not a valid cert data", cert_type, cert_no)
	}

	player := new(md.Player)
	err = query_result.One(player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

//检查player的证件号是否已经有报名制定比赛，如果有则返回error，否则nil
func checkIfPlayersAlreadyRegGameByCertificate(players *[]*md.Player, game_id uint32) error {
	for _, player := range *players {
		registration, err := checkIfPlayerAlreadyRegGameByCertificate(player.CertificateType, player.CertificateNo, game_id)
		if err != nil {
			continue
		}

		return fmt.Errorf("证件号码%v已经报名参加该比赛，订单号码为%v", player.CertificateNo, registration.Order.OrderNo)
	}

	return nil
}

func checkIfPlayerAlreadyRegGameByCertificate(cert_type uint8, cert_no string, game_id uint32) (*md.Registration, error) {
	player, err := GetPalyerByCertificate(cert_type, cert_no)
	if err != nil {
		return nil, err
	}

	return checkIfPlayerAlreadyRegGame(player.Id, game_id)
}

//检查玩者是否已经报名比赛，如果是，则返回Registration
func checkIfPlayerAlreadyRegGame(player_id uint32, game_id uint32) (*md.Registration, error) {
	query_result := md.Orm.QueryTable(md.DB_TABLE_REGISTRATION).
		Filter(md.DB_REG_PLAYER_ID, player_id).Filter(md.DB_REG_GAME_ID, game_id)

	count, err := query_result.Count()
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("%d is not signup to game %d", player_id, game_id)
	}

	registration := new(md.Registration)
	err = query_result.One(registration)
	if err != nil {
		return nil, err
	}

	return registration, nil
}

func updatePlayersDB(signup_orm orm.Ormer, players *[]*md.Player) (err error) {
	for _, player := range *players {
		err := updatePlayerDB(signup_orm, player)
		if err != nil {
			return err
		}
	}

	return nil
}

func updatePlayerDB(signup_orm orm.Ormer, player *md.Player) (err error) {
	isPlayerExist := false

	old_player, _ := GetPalyerByCertificate(player.CertificateType, player.CertificateNo)
	if old_player != nil {
		isPlayerExist = true
	}

	if isPlayerExist {
		_, err = signup_orm.Update(player)
	} else {
		var id int64
		id, err = signup_orm.Insert(player)
		if err == nil {
			player.Id = uint32(id)
		} else {

		}
	}

	if err != nil {
		return err
	}

	return nil
}

func getGameIdUsrIDParam(js *simplejson.Json) (uint32, uint32, error) {
	game_id, err := utils.GetUInt32FromParam(js, HP_GAME_ID)
	if err != nil {
		return 0, 0, err
	}

	user_id, err := utils.GetUInt32FromParam(js, HP_USER_ID)
	if err != nil {
		return 0, 0, err
	}

	return game_id, user_id, nil
}

func genRegQueueToken() string {
	tokenStr := guid.NewGUID().String()

	return strings.Trim(tokenStr, "-")
}

func GetGameByIDFromCache(game_id uint32) *md.Game {
	cache_manager := utils.GetCacheManager()

	game_id_str := fmt.Sprintf("%d", game_id)
	key := md.DB_TABLE_GAME + game_id_str
	game, _ := cache_manager.Get(key).(*md.Game)
	if game == nil {
		game, _ = GetGameByID(fmt.Sprintf("%d", game_id))
		if game != nil {
			cache_manager.Set(key, game)
		}

		return game
	}

	return game
}

func RemoveGameFromCache(game_id uint32) {
	cache_manager := utils.GetCacheManager()
	cache_manager.Delete(md.DB_TABLE_GAME + fmt.Sprintf("%d", game_id))
}

//查找game是否还有名额
func IsGameHaveBalance(game_id uint32) bool {
	game_limitation, _ := g_game_limitation_map.Get(game_id).(*GameLimitation)
	if game_limitation == nil {
		return false
	}

	return game_limitation.balance > 0 || game_limitation.limitation == 0
}

func parsePlayersInfo(js *simplejson.Json, user_id uint32) (*[]*md.Player, error) {
	a, err := js.Array()
	if err != nil || len(a) <= 0 {
		return nil, err
	}

	players := make([]*md.Player, 0)
	for i := 0; i < len(a); i++ {
		player_js := js.GetIndex(i)

		param_values, err := utils.ParseParam(player_js, GetPlayerInfoParameterMap())
		if err != nil {
			return nil, err
		}

		player := new(md.Player)
		player.UserId = user_id
		err = ParsePlayerInfoFromJson(player, param_values)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return &players, nil
}

//reg相关接口实现
func RegGetQueueToken(js *simplejson.Json) map[string]interface{} {
	retJson := init_retJson()

	if len(g_reg_queue) >= REG_CHANNEL_CAPACITY {
		retJson_edit(retJson, err_code.ServerErr, "抱歉，报名队列已满，请过会儿再试")
		return retJson
	}

	game_id, user_id, err := getGameIdUsrIDParam(js)
	if err != nil {
		retJson_edit(retJson, err_code.ErrJson, "调用getGameIdUsrIDParam返回错误\n"+err.Error())
		return retJson
	}

	game := GetGameByIDFromCache(game_id)
	if game == nil {
		retJson_edit(retJson, err_code.InvalidGameId, "调用 GetGameByIDFromCache 返回错误")
		return retJson
	}

	if !IsGameHaveBalance(game_id) {
		retJson[HP_QUEUE_STATE] = REG_QUEUE_STATE_FULL
		return retJson
	}

	players, err := parsePlayersInfo(js.Get(HP_PLAYERS), user_id)
	if err != nil {
		retJson_edit(retJson, err_code.ErrJson, "调用 parsePlayersInfo 返回错误\n"+err.Error())
		return retJson
	}

	token := genRegQueueToken()
	reg_item := RegItem{
		token:            token,
		user_id:          user_id,
		players:          *players,
		game:             game,
		order_no:         "",
		last_access_time: utils.GetNowTimeInSec(),
		queue_idx:        atomic.AddUint64(&g_reg_idx, 1)}

	if !g_queued_reg_item_map.Set(token, &reg_item) {
		utils.Logger.Error("failed in setting reg atomic map: key %v", token)
		retJson_edit(retJson, err_code.ServerErr, "failed in setting atomic map")
		return retJson
	}

	g_reg_queue <- reg_item

	retJson[HP_TOKEN] = token
	retJson[HP_QUEUE_STATE] = REG_QUEUE_STATE_QUEUING
	retJson[HP_COUNT] = len(g_reg_queue)
	retJson[HP_QUERY_INTERVAL] = QueryInterval()

	return retJson
}

func RegQueryQueueTokenState(js *simplejson.Json) map[string]interface{} {
	retJson := init_retJson()
	token, err := js.Get(HP_TOKEN).String()
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "不合法的token")
		return retJson
	}

	reg_item, _ := g_queued_reg_item_map.Get(token).(*RegItem)
	if reg_item != nil {
		reg_item.last_access_time = utils.GetNowTimeInSec()

		queue_idx := reg_item.queue_idx - g_cur_processing_reg_idx
		if queue_idx < 1 {
			queue_idx = 1
		}
		retJson[HP_QUEUE_STATE] = REG_QUEUE_STATE_QUEUING
		retJson[HP_COUNT] = queue_idx
		retJson[HP_QUERY_INTERVAL] = QueryInterval()
		retJson_edit(retJson, err_code.OK, "")
	} else if reg_item, _ := g_processed_map.Get(token).(*RegItem); reg_item != nil {
		reg_item.last_access_time = utils.GetNowTimeInSec()
		retJson[HP_ORDER_NO] = reg_item.order_no
		retJson[HP_QUEUE_STATE] = reg_item.queue_state
	} else {
		retJson_edit(retJson, err_code.InvalidData, "无效的token")
		return retJson
	}

	return retJson
}

func RegSetGameLimitation(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()
	//check user permission
	if !utils.IsClientFromSamePrivateNetwork(ctx) {
		retJson_edit(retJson, err_code.InvalidSessionId, "")
		return retJson
	}

	game_id, err := utils.GetUInt32FromParam(js, HP_GAME_ID)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}

	game := GetGameByIDFromCache(game_id)
	if game == nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}

	limitation, err := utils.GetUInt32FromParam(js, HP_LIMITATION)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, "")
		return retJson
	}

	balance, err := utils.GetUInt32FromParam(js, HP_BALANCE)
	if err != nil || limitation < balance {
		balance = limitation
	}

	if !g_game_limitation_map.Set(game_id, &GameLimitation{limitation: int32(limitation), balance: int32(balance)}) {
		utils.Logger.Error("failed in setting game limitation map")
		retJson_edit(retJson, err_code.ServerErr, "failed in setting game limitation map")
		return retJson
	}

	return retJson
}

func RegSetAllGamesLimitation(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()
	//check user permission
	if !utils.IsClientFromSamePrivateNetwork(ctx) {
		retJson_edit(retJson, err_code.InvalidSessionId, "")
		return retJson
	}

	game_limitation_map, err := limitation.ParseBalanceInfoArrayJson(js, HP_BALANCES)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	g_game_limitation_map = game_limitation_map

	return retJson
}

func RegGetGameLimitationInfo(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()
	//check user permission
	if !utils.IsClientFromSamePrivateNetwork(ctx) {
		retJson_edit(retJson, err_code.InvalidSessionId, "")
		return retJson
	}

	game_id, err := utils.GetUInt32FromParam(js, HP_GAME_ID)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}

	game := GetGameByIDFromCache(game_id)
	if game == nil {
		retJson_edit(retJson, err_code.InvalidGameId, "")
		return retJson
	}

	game_limitation, _ := g_game_limitation_map.Get(game_id).(*GameLimitation)
	if game_limitation == nil {
		retJson[HP_LIMITATION] = 0
		retJson[HP_BALANCE] = 0
	} else {
		retJson[HP_LIMITATION] = game_limitation.limitation
		retJson[HP_BALANCE] = game_limitation.balance
	}

	return retJson
}

func RegGetAllGamesBalanceInfo(js *simplejson.Json, ctx *context.Context) map[string]interface{} {
	retJson := init_retJson()
	//check user permission
	if !utils.IsClientFromSamePrivateNetwork(ctx) {
		retJson_edit(retJson, err_code.InvalidSessionId, "")
		return retJson
	}

	balances_info, err := registrationSrv.GetGamesBalanceInfo()
	if err != nil {
		retJson_edit(retJson, err_code.NoData, err.Error())
		return retJson
	}

	retJson["balances"] = balances_info

	return retJson
}
