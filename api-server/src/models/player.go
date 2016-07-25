package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
)

var (
	ErrPlayerNotFound = errors.New("models.Player.FindBy")
)

const (
	DB_PLAYER_CERTIFICATE_TYPE = "certificate_type"
	DB_PLAYER_CERTIFICATE_NO   = "certificate_no"
)

// Player 表示一个选手的信息
type Player struct {
	Id                     uint32
	UserId                 uint32
	Name                   string `orm:"size(32)"`
	CertificateType        uint8
	CertificateNo          string `orm:"size(20)"`
	Mobile                 string `orm:"size(18)"`
	Email                  string `orm:"size(32)"`
	Country                string `orm:"size(32)"`
	Province               string `orm:"size(32)"`
	City                   string `orm:"size(32)"`
	Address1               string `orm:"size(32)"`
	Address2               string `orm:"size(32)"`
	Zip                    string `orm:"size(6)"`
	Gender                 uint8
	BirthDate              time.Time `type(date)"`
	EmergencyContactName   string    `orm:"size(32)"`
	EmergencyContactMobile string    `orm:"size(18)"`
	TShirtSize             uint8

	ExtraInfoJson string `orm:"type(text)"`
}

func NewPlayer() *Player {
	return new(Player)
}

func (p *Player) FindBy(conditions map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.Logger.Debug("models.Player.FindBy : error : %s", err.Error())
		}
	}()
	qs := Orm.QueryTable(p)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err = qs.One(p)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return ErrPlayerNotFound
	}
	return nil

}

// Plyaers 表示一个比赛所有的参数选手
type Players []Player

// GetPlayersByGameId 通过game_id获取一个比赛所有的参赛选手
func GetPlayersByGameId(game_id string) (*Players, error) {
	players := Players{}
	err := Orm.Raw("SELECT p.*,r.game_id FROM "+DB_TABLE_PLAYER+" as p inner join "+DB_TABLE_REGISTRATION+" as r on p.id = r.player_id where r.game_id = ?", game_id).QueryRow(&players)
	if err == orm.ErrNoRows {
		return nil, err
	}
	return &players, nil
}

// Count 获取选手个数
func (ps Players) Count() int {
	return len(ps)
}
