package utils

import (
	"github.com/dlintw/goconf"
)

var Cfg *goconf.ConfigFile

func InitCfg(cfg *goconf.ConfigFile) {
	Cfg = cfg
}

func GetCfgString(section, option string) (retStr string) {
	retStr, _ = Cfg.GetString(section, option)
	return
}
func GetCfgInt(section, option string) (retInt int) {
	retInt, _ = Cfg.GetInt(section, option)
	return
}
func GetCfgBool(section, option string) (retBool bool) {
	retBool, _ = Cfg.GetBool(section, option)
	return
}
