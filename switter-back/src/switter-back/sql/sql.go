package sql

import (
	"database/sql"
	"fmt"
	"log"

	"switter-back/types"

	_ "github.com/lib/pq"
)

var (
	dbConn *sql.DB
)

func CreateConn(conf types.SqlConfiguration) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.HostName, conf.Port, conf.UserName, conf.Password, conf.DBName)
	var err error
	dbConn, err = sql.Open(conf.DriverName, dsn)
	if err != nil {
		log.Println("sql.CreateConn err: ", err)
	}
}

// CloseConn ...
func CloseConn() {
	//close(dbConn)
}
