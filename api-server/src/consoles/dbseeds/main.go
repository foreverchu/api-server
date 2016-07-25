package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"models"
	"os"
	"path/filepath"

	"github.com/chinarun/utils"
	"github.com/dlintw/goconf"
	"github.com/manveru/faker"
)

var fake *faker.Faker

func init() {
	env := utils.GetEnvMode()
	if env == utils.ENV_MODE_PRODUCTION_STR {
		log.Fatal("you should never do this in production mode")
	}

	basePath, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}

	rootPath := filepath.Dir(filepath.Dir(basePath))

	log.Print("currnet path :", rootPath)

	conf_file := flag.String("config", rootPath+"/conf/"+env+".conf", "set run config file.")
	flag.Parse()
	l_conf, err := goconf.ReadConfigFile(*conf_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", conf_file, err)
		os.Exit(1)
	}

	utils.InitCfg(l_conf)

	models.InitModels()

	fake, err = faker.New("en")
	if err != nil {
		log.Fatal(err)
	}

	fake.Rand = rand.New(rand.NewSource(42))

}

func main() {
	num := 10
	CreateUserAndProfile(num)
	CreatePartyAndGame(num)
	CreatePlayerAndRegAndOrder(num)

	TruncateTable(models.DB_TALBE_MESSAGE)
	CreateMessages(num)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func TruncateTable(table string) {
	log.Printf("start to truncate table: %s\n", table)
	var err error
	_, err = models.Orm.Raw("set foreign_key_checks = ?", 0).Exec()
	_, err = models.Orm.Raw("truncate `" + table + "`").Exec()
	_, err = models.Orm.Raw("set foreign_key_checks = ?", 1).Exec()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("truncate table: %s finished\n", table)
}

func CreateUserAndProfile(num int) {
	TruncateTable(models.DB_TABLE_USER)
	CreateUsers(10)

	TruncateTable(models.DB_TABLE_PROFILE)
	CreateProfiles(10)
}

func CreatePartyAndGame(num int) {
	TruncateTable(models.DB_TABLE_PARTY)
	CreateParty(10)

	TruncateTable(models.DB_TABLE_PARTY_DETAIL)
	CreatePartyDetail(10)

	TruncateTable(models.DB_TABLE_GAME)
	CreateGames(10)

}

func CreatePlayerAndRegAndOrder(num int) {
	TruncateTable(models.DB_TABLE_PLAYER)
	//CreatePlayers(num)

	TruncateTable(models.DB_TABLE_ORDER)
	//CreateOrders(num)

	TruncateTable(models.DB_TABLE_REGISTRATION)
	//CreateRegistration(num)
}
