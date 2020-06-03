package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"switter-back/types"
)

/*
type sqlConfiguration struct {
	Drivername string `json:"drivername"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Port       int16  `json:"port"`
}

//--------------------------------
type appConfiguration struct {
	Host string           `json:"host"`
	Port int16            `json:"port"`
	SQL  sqlConfiguration `json:"sql"`
}
*/
var (
	AppConf *types.AppConfiguration
)

// LoadConfig from app home dir
func LoadConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("config.LoadConfig #1, getting working directory error: ", err)
	}
	// open file
	jsonFile, err := os.Open(workDir + "/conf.json")
	//jsonFile, err := os.Open(homeDir + "conf.json")
	//log.Println("	path: ", homeDir + "conf.json")
	if err != nil {
		log.Println("config.LoadConfig #2, config file opening error: ", err)
	}
	defer jsonFile.Close()
	// read file content
	byteFileContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("config.LoadConfig #3, config file readin error: ", err)
	}
	//var encodedConfigData map[string]interface{}
	AppConf = &types.AppConfiguration{}
	err = json.Unmarshal([]byte(byteFileContent), AppConf)
	if err != nil {
		log.Println("config.LoadConfig #4, config file decoding error: ", err)
	}

}
