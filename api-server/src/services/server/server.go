package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/astaxie/beego/httplib"
	"github.com/chinarun/utils"
)

var (
	g_core_server_host string
)

const (
	API_PATH_API_SERVER = "/api/api_server"
)

func get_core_server_host() string {
	if len(g_core_server_host) <= 0 {
		g_core_server_host, _ = utils.Cfg.GetString("beego", "core_server_host")
	}

	return g_core_server_host
}

func PreServerRunning() error {
	InitShutdownHook()

	err := registerApiServer()
	if err != nil {
		return err
	}

	return nil
}

func PostServerRunning() {
	fmt.Println("PostServerRunning")
	unregisterApiServer()
}

func InitShutdownHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		PostServerRunning()
		os.Exit(1)
	}()
}

func registerApiServer() error {
	return callRegisterApi("APIServerRegister")
}

func unregisterApiServer() error {
	return callRegisterApi("APIServerUnRegister")
}

func callRegisterApi(func_name string) error {
	req := httplib.Post(get_core_server_host() + API_PATH_API_SERVER)
	req.Body(fmt.Sprintf("{\"cmd\":\"%s\"}", func_name))

	res_str, err := req.String()
	if err != nil {
		utils.Logger.Error("failed in request %s", func_name)
		return fmt.Errorf("failed in request %s", func_name)
	}

	err = utils.CheckResultJsonStr(res_str)
	if err != nil {
		return err
	}

	return nil
}
