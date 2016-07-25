package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/chinarun/utils"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Orm orm.Ormer
)

func setDebug() {
	if utils.GetCfgBool("def_db", "db_debug") == true {
		orm.Debug = true
	}
}

func setDriver() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
}

// 设置连接池大小
func setMaxIdelConn() {
	orm.SetMaxIdleConns("default", 1000)
}

func InitModels() {
	setDebug()
	setDriver()

	var db_def string

	host := utils.GetCfgString("def_db", "db_host")
	user := utils.GetCfgString("def_db", "db_user")
	pass := utils.GetCfgString("def_db", "db_pass")
	name := utils.GetCfgString("def_db", "db_name")
	port := utils.GetCfgString("def_db", "db_port")
	socket := utils.GetCfgString("def_db", "db_socket")
	charset := utils.GetCfgString("def_db", "db_charset")

	if host == "localhost" || host == "127.0.0.1" {
		db_def = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=%s&loc=Local", user, pass, socket, name, charset)
	} else {
		db_def = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local", user, pass, host, port, name, charset)
	}

	orm.RegisterDataBase("default", utils.GetCfgString("def_db", "db_type"), db_def)

	//	注册数据库表
	orm.RegisterModel(new(User), new(Order), new(Party))
	orm.RegisterModel(new(Game), new(Player), new(Registration))
	orm.RegisterModel(new(PlayerScore), new(Tag), new(PartyDetail))
	orm.RegisterModel(new(Photo), new(Profile), new(ThirdPartyRegister))
	orm.RegisterModel(new(Message), new(PartyTagMap), new(EmailConfirm))
	orm.RegisterModel(new(UserParty), new(Smscode))

	Orm = orm.NewOrm()

	setMaxIdelConn()
}
