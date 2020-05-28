package main

import (
	//sqlpower "github.com/theamazingeagle/switter-back/sql"
	"github.com/theamazingeagle/switter-back/configuration"
	"github.com/theamazingeagle/switter-back/router"
	"github.com/theamazingeagle/switter-back/sql"
	//
	"log"
)
func main() {
	// config loadin ...
	configuration.LoadConfig()
	//log.Printf("CONFIG: %+v\n\n", configuration.AppConf)
	//
	log.Println("starting at ", configuration.AppConf.Port, " port")
	// init sql connection
	sql.CreateConn(configuration.AppConf.SQL)
	//
	//router.Start(configuration.AppConf.Host, configuration.AppConf.Port)
	router.Start(configuration.AppConf.Host, configuration.AppConf.Port)
}
