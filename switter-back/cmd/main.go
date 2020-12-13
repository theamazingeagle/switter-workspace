package main

import (
	"encoding/json"
	"go-http/configuration"
	"go-http/sql"
	"io/ioutil"
	"os"

	//
	"log"
)

type SqlConfiguration struct {
	HostName   string `json:"hostname"`
	DriverName string `json:"drivername"`
	DBName     string `json:"dbname"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Port       int16  `json:"port"`
}

//--------------------------------
type AppConfiguration struct {
	Host string           `json:"host"`
	Port int              `json:"port"`
	SQL  SqlConfiguration `json:"sql"`
}

var (
	AppConf *types.AppConfiguration
)

func main() {
	loadConfig()
	log.Println("starting at ", configuration.AppConf.Port, " port")
	sql.CreateConn(AppConf.SQL)
	router.Start(configuration.AppConf.Host, configuration.AppConf.Port)
}

// LoadConfig from app home dir
func loadConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("config.LoadConfig #1, getting working directory error: ", err)
	}
	jsonFile, err := os.Open(workDir + "/conf.json")
	if err != nil {
		log.Println("config.LoadConfig #2, config file opening error: ", err)
	}
	defer jsonFile.Close()
	byteFileContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("config.LoadConfig #3, config file readin error: ", err)
	}
	AppConf = &types.AppConfiguration{}
	err = json.Unmarshal([]byte(byteFileContent), AppConf)
	if err != nil {
		log.Println("config.LoadConfig #4, config file decoding error: ", err)
	}

}
