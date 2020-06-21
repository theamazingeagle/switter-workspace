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
func CloseConn() {
	//close(dbConn)
}

// GetMessages() returns all messages
func GetMessages() []types.MessageInfo {
	rows, err := dbConn.Query(`select m.message_id,m.message_text, m.message_date, u.user_name
							from messages m
							inner join users u on u.user_id=m.message_userid;`)
	if err != nil {
		log.Println("sql.GetMessages err: ", err)
	}
	//messages := make([]*types.Message,0)
	messages := []types.MessageInfo{}
	for rows.Next() {
		message := &types.MessageInfo{}
		rows.Scan(&message.MessageID, &message.Text, &message.Date, &message.UserName)
		log.Println("extracted contains: ", message)
		messages = append(messages, *message)
	}

	return messages
}
