package orderSrv

import (
	"models"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
)

const (
	ORDER_REFRESH_ROUTINE_INTERVAL = (60 * time.Second)
	ORDER_EXPIRE_INTERVAL          = (30 * time.Minute)

	IS_NULL = "__isnull"
	LT      = "__lt"

	REG_PAY_STATE_CANCEL = 2
)

func CancelExpiredUnPayedOrderRoutine() {
	go func() {
		for {
			CancelExpiredUnPayedOrder()
			time.Sleep(ORDER_REFRESH_ROUTINE_INTERVAL)
		}
	}()
}

func CancelExpiredUnPayedOrder() {
	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("Error in CancelExpiredUnPayedOrder : %s", err.Error())
		}
	}()

	query_result := models.Orm.QueryTable(models.DB_TABLE_ORDER).
		Filter(models.DB_ORDER_PAY_TIME+IS_NULL, true).
		Filter(models.DB_ORDER_REFUND_TIME+IS_NULL, true).
		Filter(models.DB_ORDER_CANCEL_TIME+IS_NULL, true).
		Filter(models.DB_ORDER_SUBMIT_TIME+LT, time.Now().Add(-ORDER_EXPIRE_INTERVAL))

	count, err := query_result.Count()
	if err != nil {
		return
	}

	if count <= 0 {
		return
	}

	var orders []*models.Order

	_, err = query_result.All(&orders)
	if err != nil {
		return
	}

	for _, order := range orders {
		CancelOrder(order)
	}
}

func CancelOrder(order *models.Order) error {
	var err error
	defer func() {
		if err != nil {
			utils.Logger.Error("Error in CancelOrder : %s", err.Error())
		}
	}()

	order.CancelTime = time.Now()

	order_orm := orm.NewOrm()
	order_orm.Begin()

	_, err = order_orm.Update(order, models.DB_ORDER_CANCEL_TIME)
	if err != nil {
		return err
	}

	_, err = order_orm.Raw("UPDATE registration SET pay_state = ? WHERE order_id = ?",
		REG_PAY_STATE_CANCEL, order.Id).Exec()
	if err != nil {
		return err
	}

	err = order_orm.Commit()
	if err != nil {
		err_rollback := order_orm.Rollback()
		if err_rollback != nil {
			utils.Logger.Critical("数据库回滚失败 order: %v", order.Id)
			return err_rollback
		}

		return err
	}

	return nil
}
