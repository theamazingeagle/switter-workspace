package sql

import (
	"log"

	"switter-back/types"
)

func CreateMessage(url, date, text string, userID int64) {
	_, err := dbConn.Exec("INSERT INTO messages(message_url, message_date, message_text, message_userid) VALUES($1, $2, $3, $4)",
		url, date, text, userID)
	if err != nil {
		log.Println("sql.CreateUser err: ", err)
	}
}
func GetMessage(messageID string) types.Message {
	row := dbConn.QueryRow("SELECT * FROM messages WHERE message_id=$1", messageID)
	if row == nil {
		log.Println("sql.GetMessage err: ", row)
	}
	message := &types.Message{}
	row.Scan(&message.ID, &message.Url, &message.UserID, &message.Date, &message.Text)
	return *message
}

// GetMessages() returns all messages
func GetMessages() []types.Message {
	rows, err := dbConn.Query("SELECT * FROM messages")
	if err != nil {
		log.Println("sql.GetMessages err: ", err)
	}
	//messages := make([]*types.Message,0)
	messages := []types.Message{}
	for rows.Next() {
		message := &types.Message{}
		rows.Scan(&message.ID, &message.Url, &message.UserID, &message.Date, &message.Text)
		log.Println("extracted contains: ", message)
		messages = append(messages, *message)
	}

	return messages
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
