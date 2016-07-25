package orderSrv

import (
	"models"
	"time"
)

type PayState int

const (
	ORDER_PAY_STATE_WAIT_PAY PayState = 0 //等待支付
	ORDER_PAY_STATE_PAYED    PayState = 1 //已支付
	ORDER_PAY_STATE_CANCELED PayState = 2 //已取消
	ORDER_PAY_STATE_REFUNDED PayState = 3 //已退款
)

var payStateDesc = []string{
	"等待支付",
	"已支付",
	"已取消",
	"已退款",
}

func (p PayState) String() string {
	return payStateDesc[p]
}

const (
	order_expires_time = time.Duration(30 * time.Minute)
)

type OrderState struct {
	order *models.Order
}

func NewOrderState(order *models.Order) *OrderState {
	os := new(OrderState)
	os.order = order
	return os
}

func NewOrderStateWithOrderNo(orderNo string) (*OrderState, error) {
	os := new(OrderState)
	order := models.NewOrder()
	cond := map[string]interface{}{
		models.DB_ORDER_ORDER_NO: orderNo,
	}
	if err := order.FindBy(cond); err == models.ErrOrderNotFound {
		return nil, ErrInvalidOrder
	}
	os.order = order
	return os, nil
}

func (os *OrderState) IsCanceled() bool {
	return !os.order.CancelTime.IsZero()
}

func (os *OrderState) IsPayed() bool {
	return !os.order.PayTime.IsZero()
}

func (os *OrderState) IsRefund() bool {
	return !os.order.RefundTime.IsZero()
}

// 判断是否过了支付时间
func (os *OrderState) IsExpires() bool {
	return time.Since(os.order.SubmitTime) > order_expires_time
}

func (os *OrderState) State() PayState {
	if os.IsPayed() {
		return ORDER_PAY_STATE_PAYED
	}

	if os.IsRefund() {
		return ORDER_PAY_STATE_REFUNDED
	}

	if os.IsCanceled() {
		return ORDER_PAY_STATE_CANCELED
	}
	return ORDER_PAY_STATE_WAIT_PAY
}
