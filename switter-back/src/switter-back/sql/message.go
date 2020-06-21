package sql

import (
	"fmt"
	"log"

	"switter-back/types"
)

//
func CreateMessage(text string, userID int) error {
	//_, err := dbConn.Exec("INSERT INTO messages(message_url, message_date, message_text, message_userid) VALUES($1, $2, $3, $4)",
	_, err := dbConn.Exec(`INSERT INTO messages(message_text, message_userid)
						  VALUES($1, $2);`,
		text, userID)
	if err != nil {
		return fmt.Errorf("sql.CreateUser err: ", err)
	}
	return nil
}

//
func GetMessage(messageID string) types.Message {
	row := dbConn.QueryRow("SELECT * FROM messages WHERE message_id=$1 ;", messageID)
	if row == nil {
		log.Println("sql.GetMessage err: ", row)
	}
	message := &types.Message{}
	row.Scan(&message.ID, &message.Url, &message.UserID, &message.Date, &message.Text)
	return *message
}

func UpdateMessage(newText string, userID int) {
	_, err := dbConn.Exec("UPDATE messages SET message_text = $1 WHERE user_id = $2", newText, userID)
	if err != nil {
		log.Println("sql.UpdateMessage err: ", err)
	}
}
func DeleteMessage(ID int64) {
	_, err := dbConn.Exec("DELETE FROM messages WHERE message_id=$1", ID)
	if err != nil {
		log.Println("sql.DeleteMessage err: ", err)
	}
}
