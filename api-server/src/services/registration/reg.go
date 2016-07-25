package registrationSrv

import (
	"fmt"
	"models"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
	"github.com/chinarun/utils/limitation"
)

const (
	ORDER_PAY_STATE_NONE = 100 //无效状态， 因为是uint8， 所以不用-1

	ORDER_PAY_STATE_FIRST    = 0
	ORDER_PAY_STATE_WAIT_PAY = 0 //等待支付
	ORDER_PAY_STATE_PAYED    = 1 //已支付
	ORDER_PAY_STATE_CANCELED = 2 //已取消
	ORDER_PAY_STATE_REFUNDED = 3 //已退款
	ORDER_PAY_STATE_LAST     = 3
)

func GetGamesBalanceInfo() ([]limitation.BalanceInfo, error) {
	var balance_info []limitation.BalanceInfo

	queryselector := models.Orm.QueryTable(models.DB_TABLE_GAME)
	queryselector = queryselector.Filter("start_time__gte", time.Now())

	var games []models.Game

	count, err := queryselector.All(&games)
	if err == orm.ErrNoRows {
		return balance_info, fmt.Errorf("现在没有符合条件的比赛")
	} else if err != nil {
		utils.Logger.Error("failed in query game data: %v", err.Error())
		return balance_info, err
	}

	balance_info = make([]limitation.BalanceInfo, count)

	for i := 0; i < int(count); i++ {
		game_limitation := games[i].Limitation

		reg_count, err := models.Orm.QueryTable(models.DB_TABLE_REGISTRATION).
			Filter("game_id", games[i].Id).Filter("pay_state__in",
			ORDER_PAY_STATE_WAIT_PAY, ORDER_PAY_STATE_PAYED).Count()

		if err != nil && err != orm.ErrNoRows {
			utils.Logger.Error("failed in query order data: game_id ", games[i].Id)
			continue
		}
		balance := game_limitation - uint32(reg_count)

		balance_info[i] = limitation.BalanceInfo{GameId: strconv.Itoa(int(games[i].Id)), Limitation: int(game_limitation), Balance: int(balance)}
	}

	return balance_info, nil
}
