package main

import (
	"flag"
	"fmt"
	"html/template"
	"models"
	"net/http"
	"os"
	"path"
	_ "routers"
	"runtime"
	"time"

	"controllers"
	"services/message"
	"services/notice"
	"services/order"

	"github.com/astaxie/beego"
	"github.com/chinarun/utils"
	"github.com/dlintw/goconf"
)

// 自定义错误信息.
func ConfError(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").ParseFiles(beego.ViewsPath + "/dberror.html")
	data := make(map[string]interface{})
	data["content"] = "database is now down"
	t.Execute(rw, data)
}

func init_app() {
	//	APP_VER, _ := utils.Cfg.GetString("beego", "app_ver")

	beego.AppName, _ = utils.Cfg.GetString("beego", "app_name")
	beego.HttpPort, _ = utils.Cfg.GetInt("beego", "http_port")

	beego.EnableHttpListen, _ = utils.Cfg.GetBool("beego", "http_on")
	beego.DirectoryIndex = true
	tmp, _ := utils.Cfg.GetInt("beego", "http_server_timeout")
	beego.HttpServerTimeOut = int64(tmp)
	beego.CopyRequestBody = true

	//设置session
	beego.SessionOn, _ = utils.Cfg.GetBool("beego", "sessionon")
	if beego.SessionOn {
		beego.SessionName, _ = utils.Cfg.GetString("beego", "session_name")
		session_gc_maxlifetime, _ := utils.Cfg.GetInt("beego", "session_gc_maxlifetime")
		beego.SessionCookieLifeTime, _ = utils.Cfg.GetInt("beego", "session_cookie_life_time")
		beego.SessionGCMaxLifetime = int64(session_gc_maxlifetime)
		beego.SessionProvider, _ = utils.Cfg.GetString("beego", "session_provider")
		beego.SessionSavePath, _ = utils.Cfg.GetString("beego", "session_savepath")
	}
	beego.AppConfigPath = utils.GetConfFilePath()

	controllers.InitRegRoutine()
	//server.PreServerRunning()
	orderSrv.CancelExpiredUnPayedOrderRoutine()
}

func report() {
	time.AfterFunc(60*time.Second, func() {
		memStat := new(runtime.MemStats)
		runtime.ReadMemStats(memStat)
		info_msg := fmt.Sprintf("使用内存 %d KB, GO程数: %d.", memStat.Alloc/1024, runtime.NumGoroutine())
		beego.Info(info_msg)
		report()
	})
}

func init_dirs() {
	static_path, _ := utils.Cfg.GetString("path", "static_path")
	head_portrait_path, _ := utils.Cfg.GetString("path", "head_portrait")

	dir := path.Join(static_path, head_portrait_path)
	if utils.IsDirExists(dir) != true {
		os.MkdirAll(dir, os.ModePerm)
	}
	beego.SetStaticPath("photo", static_path)
}

func server_static() {
	env_mode := utils.GetEnvModeInt()
	if env_mode != utils.ENV_MODE_PRODUCTION {
		beego.DirectoryIndex = true
		beego.StaticDir["/static"] = "static"
	}
}

func main() {
	// 初始化配置文件.
	utils.RefreshEnvMode()
	env_mode := utils.GetEnvMode()
	fmt.Fprintf(os.Stdout, "current running mode is %s\n", env_mode)
	conf_file := flag.String("config", utils.GetConfFilePath(), "set run config file.")
	flag.Parse()
	l_conf, err := goconf.ReadConfigFile(*conf_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", conf_file, err)
		os.Exit(1)
	}

	utils.InitCfg(l_conf)
	utils.InitLogger()
	utils.InitRander()

	//从配置文件读取发送邮件或者短信必须的参数
	noticeSrv.NoticeParamsInit()

	messageSrv.MsgDeamon()
	models.InitModels()
	init_app()
	init_dirs()

	report()

	utils.Logger.Info("Start ChinaRun function server")
	server_static()

	beego.Run()
}
