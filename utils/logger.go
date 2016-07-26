package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
)

var (
	Logger *logs.BeeLogger
)

//可以设置max line size
// max file size
func InitLogger() {
	Logger = logs.NewLogger(10000)
	defer Logger.Flush()
	path, _ := os.Getwd()
	file := path + GetCfgString("log", "filename")
	if !IsExist(file) {
		os.MkdirAll(filepath.Dir(file), os.ModePerm)
	}
	file_cfg := fmt.Sprintf(`{"filename":"%s"}`, file)
	Logger.EnableFuncCallDepth(GetCfgBool("log", "enable_func_call_depth"))
	Logger.SetLogger("file", file_cfg)
	Logger.SetLevel(getLevel())
}

func getLevel() int {
	env_mode := GetEnvModeInt()
	switch env_mode {
	case ENV_MODE_PRODUCTION:
		return logs.LevelWarning
	case ENV_MODE_DEVELOPMENT:
		return logs.LevelDebug
	case ENV_MODE_TESTING:
		return logs.LevelDebug
	default:
		return logs.LevelInformational
	}
}
