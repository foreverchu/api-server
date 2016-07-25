package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
)

var ErrOrderNotFound = errors.New("order not found")

const (
	DB_ORDER_USERID      = "user_id"
	DB_ORDER_ORDER_NO    = "order_no"
	DB_ORDER_SUBMIT_TIME = "submit_time"
	DB_ORDER_PAY_TIME    = "pay_time"
	DB_ORDER_REFUND_TIME = "refund_time"
	DB_ORDER_CANCEL_TIME = "cancel_time"
	DB_ORDER_PAY_METHOD  = "pay_method"
	DB_ORDER_PAY_ACCOUNT = "pay_account"
)

type Order struct {
	Id           uint32
	UserId       uint32
	OrderNo      string    `orm:"size(32)"`
	SubmitTime   time.Time `orm:"type(datetime)"`
	PayTime      time.Time `orm:"type(datetime)"`
	RefundTime   time.Time `orm:"type(datetime)"`
	CancelTime   time.Time `orm:"type(datetime)"`
	Price        uint32
	CurrencyType uint
	PayMethod    uint
	PayAccount   string `orm:"size(32)"`

	Game *Game `orm:"rel(fk)"`
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) FindBy(conditions map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Debug("models.Order.FindBy : error : %s, cond = %v", err.Error(), conditions)
		}
	}()
	qs := Orm.QueryTable(o)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err = qs.One(o)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrOrderNotFound
	}
	return nil
}

func (o *Order) Update(columns ...string) error {
	affectedRowNum, err := Orm.Update(o, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}
