package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"switter-back/internal/server"
	"switter-back/internal/service/auth"
	"switter-back/internal/service/db/postgres"
)

type AppConf struct {
	Server server.ServerConf     `json:"server"`
	DB     postgres.PostgresConf `json:"db"`
	Auth   auth.AuthConf         `json:"auth"`
}

func main() {
	appConf, err := loadConfig("./")

	sql.CreateConn(AppConf.SQL)
	router.Start(configuration.AppConf.Host, configuration.AppConf.Port)
}

func loadConfig(path string) (AppConf, error) {
	jsonFile, err := os.Open(path + "/conf.json")
	if err != nil {
		log.Println("config.LoadConfig(), config file opening error: ", err)
		return AppConf{}, err
	}
	defer jsonFile.Close()
	byteFileContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("config.LoadConfig(), config file readin error: ", err)
		return AppConf{}, err
	}
	appConf := AppConf{}
	err = json.Unmarshal([]byte(byteFileContent), appConf)
	if err != nil {
		log.Println("config.LoadConfig(), config file decoding error: ", err)
		return AppConf{}, err
	}
	return appConf, nil
}
