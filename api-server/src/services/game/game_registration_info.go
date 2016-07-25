package gameSrv

// GameRegInfo 用于获取一个比赛的报名信息, 目前暂时有报名人数与付款人数
type GameRegInfo struct {
	applyCount uint
	payedCount uint
}

// Apply 表示报名人数 uint
func (ri GameRegInfo) ApplyCount() uint {
	return ri.applyCount
}

// Payed 表示已经支付的人数
func (ri GameRegInfo) PayedCount() uint {
	return ri.payedCount
}
