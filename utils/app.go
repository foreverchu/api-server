package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	RESULT_OK = 0
)

var env_mode_key = "CHINARUN_API_SERVER_MODE" //环境变量的key, 需要在build/run时读取并动态加载配置文件

const (
	ENV_MODE_NONE        = 0
	ENV_MODE_PRODUCTION  = 1
	ENV_MODE_TESTING     = 2
	ENV_MODE_DEVELOPMENT = 3

	ENV_MODE_PRODUCTION_STR  = "production"
	ENV_MODE_TESTING_STR     = "testing"
	ENV_MODE_DEVELOPMENT_STR = "development"
)

var env_mode_int int

func SetEnvKey(key string) {
	env_mode_key = key
}

// 获取本地环境变量 CHINARUN_API_SERVER_MODE
func getEnvModeMap() map[string]int {
	return map[string]int{
		ENV_MODE_PRODUCTION_STR:  ENV_MODE_PRODUCTION,
		ENV_MODE_TESTING_STR:     ENV_MODE_TESTING,
		ENV_MODE_DEVELOPMENT_STR: ENV_MODE_DEVELOPMENT}
}

func getEnvModeArray() []string {
	return []string{
		ENV_MODE_PRODUCTION_STR,
		ENV_MODE_TESTING_STR,
		ENV_MODE_DEVELOPMENT_STR}
}

func RefreshEnvMode() {
	modes_map := getEnvModeMap()

	var ok bool

	env_mode_int, ok = modes_map[os.Getenv(env_mode_key)]
	if !ok {
		panic("env config setting error, plz set env variable " +
			env_mode_key + " in " + strings.Join(getEnvModeArray(), ",") +
			" via export CHINARUN_API_SERVER_MODE=?")
		os.Exit(1)
	}
}

func GetEnvModeInt() int {
	return env_mode_int
}

func IsDevlopmentMode() bool {
	return GetEnvModeInt() == ENV_MODE_DEVELOPMENT
}

func GetEnvMode() string {
	modes_map := getEnvModeMap()

	mode := os.Getenv(env_mode_key)
	_, ok := modes_map[mode]
	if ok {
		return mode
	}

	panic("env config setting error, plz set env variable " +
		env_mode_key + " in " + strings.Join(getEnvModeArray(), ",") +
		" via export CHINARUN_API_SERVER_MODE=?")
	os.Exit(1)
	return ""
}

func GetRunningFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		Logger.Error(err.Error())
		return ""
	}

	return dir + "/"
}

func GetConfFilePath() string {
	env_mode := GetEnvMode()
	return GetRunningFilePath() + "conf/" + env_mode + ".conf"
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)

	} else {
		return fi.IsDir()

	}

}

func DevErrPrintDefer(retJson interface{}, req_body string) {
	rj := retJson.(map[string]interface{})
	if IsDevlopmentMode() {
		if rj["result"] != RESULT_OK {
			fmt.Printf("error: %d msg: %s \n", rj["result"], rj["msg"])
			fmt.Printf("req body: %s\n", req_body)
		}
	}
}
