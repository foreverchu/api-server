package orderSrv

import (
	"errors"
	"fmt"
	"models"
	"strconv"
	"time"

	"github.com/chinarun/utils"
)

const (
	PAY_METHOD_NONE   = 0
	PAY_METHOD_ALIPAY = 1
	PAY_METHOD_WECHAT = 2
	PAY_METHOD_PAYPAL = 3
	PAY_METHOD_OTHER  = 4
)

var (
	ErrInvalidOrder = errors.New("非法订单")
	ErrPayed        = errors.New("订单已经支付")
	ErrRefund       = errors.New("订单已经退订")
	ErrExpires      = errors.New("订单已经过期")
	ErrCanceled     = errors.New("订单已经取消")
)

type Order struct {
	orderNo string
	order   *models.Order
	state   *OrderState
	game    *models.Game
	party   *models.Party
}

func NewOrder(orderNo string) (o *Order, err error) {
	o = &Order{}
	o.orderNo = orderNo
	if err = o.FindOrder(); err != nil {
		return nil, err
	}
	return
}

func (o *Order) FindOrder() (err error) {
	o.order = models.NewOrder()
	cond := map[string]interface{}{
		models.DB_ORDER_ORDER_NO: o.orderNo,
	}
	if err = o.order.FindBy(cond); err == models.ErrOrderNotFound {
		return ErrInvalidOrder
	}

	o.state = NewOrderState(o.order)
	return
}

func (o *Order) Game() (*models.Game, error) {
	if o.game != nil {
		return o.game, nil
	}
	o.game = &models.Game{}
	cond := map[string]interface{}{
		models.DB_ID: o.order.Game.Id,
	}
	if err := o.game.FindBy(cond); err != models.ErrGameNotFound {
		return nil, err
	}
	return o.game, nil
}

func (o *Order) Party() (*models.Party, error) {
	if o.party != nil {
		return o.party, nil
	}
	o.party = &models.Party{}
	game, err := o.Game()
	if err != nil {
		return nil, err
	}
	cond := map[string]interface{}{
		models.DB_ID: game.Party.Id,
	}
	if err := o.party.FindBy(cond); err != models.ErrPartyNotFound {
		return nil, err
	}
	return o.party, nil
}

func generateOrderNumber() string {
	now := time.Now()
	ns := now.UnixNano() % 1000000
	return fmt.Sprintf("%v%06d%04d", now.Format("20060102150405"),
		ns, utils.Rander.Intn(10000))[2:]
}

func (o *Order) GenerateOrderNo() string {
	return generateOrderNumber()
}

func (o *Order) Valid() (err error) {
	if o.state.IsPayed() {
		return ErrPayed
	}

	if o.state.IsRefund() {
		return ErrRefund
	}

	if o.state.IsExpires() {
		return ErrExpires
	}

	if o.state.IsCanceled() {
		return ErrCanceled
	}

	return
}

func (o *Order) GetNo() string {
	return o.order.OrderNo
}

func (o *Order) GetPrice() int {
	return int(o.order.Price)
}

// 商品描述
func (o *Order) GetDesc() string {
	var gameName, partyName string
	game, err := o.Game()
	if err != nil {
		gameName = game.Name
	}

	party, err := o.Party()
	if err != nil {
		partyName = party.Name
	}
	return partyName + "-" + gameName
}

func (o *Order) GetDetail() string {
	return o.GetDesc()
}

func (o *Order) GetProductId() string {
	return strconv.Itoa(int(o.order.Game.Id))
}

func (o *Order) Update(params map[string]interface{}) (err error) {
	for k, v := range params {
		switch k {
		case models.DB_ORDER_PAY_TIME:
			o.order.PayTime = v.(time.Time)
		case models.DB_ORDER_PAY_METHOD:
			o.order.PayMethod = uint(v.(int))
		case models.DB_ORDER_PAY_ACCOUNT:
			o.order.PayAccount = v.(string)
		}
	}
	return o.order.Update(models.DB_ORDER_PAY_METHOD, models.DB_ORDER_PAY_TIME, models.DB_ORDER_PAY_ACCOUNT)
}
