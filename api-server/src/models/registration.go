package models

const (
	DB_REG_PLAYER_ID        = "player_id"
	DB_REG_GAME_ID          = "game_id"
	DB_REG_META_DATA_REG_ID = "reg_id"
)

type Registration struct {
	Id uint32

	Game     *Game   `orm:"rel(fk)"`
	Order    *Order  `orm:"rel(fk)"`
	Player   *Player `orm:"rel(fk)"`
	PayState uint8   // 0表示等待支付 1表示已经支付 2表示取消订单 3表示订单被退款
}

// GetRegistrationByGameID 通过game_id获取注册列表
func GetRegistrationByGameID(game_id string) ([]*Registration, error) {
	query_result := Orm.QueryTable(DB_TABLE_REGISTRATION).Filter("game_id", game_id)

	registrations := []*Registration{}
	_, err := query_result.All(&registrations)

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

// GetRegistrationByPlayerID 通过player_id获取支付状态
func GetRegistrationByPlayerID(player_id string) ([]*Registration, error) {
	query_result := Orm.QueryTable(DB_TABLE_REGISTRATION).Filter("player_id", player_id)

	registrations := []*Registration{}
	_, err := query_result.All(&registrations)

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

// IsWaiting 表示等待支付
func (reg *Registration) IsWaiting() bool {
	return reg.PayState == 0
}

// IsPayed 表示已经支付
func (reg *Registration) IsPayed() bool {
	return reg.PayState == 1
}

// IsCanceled 表示已经取消
func (reg *Registration) IsCanceled() bool {
	return reg.PayState == 2
}

// IsRefunded 表示订单被退款
func (reg *Registration) IsRefunded() bool {
	return reg.PayState == 3
}
