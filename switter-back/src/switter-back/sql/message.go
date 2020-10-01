package sql

import (
	"fmt"
	"log"
	"strconv"

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

// GetMessages() returns all messages
func GetMessages(page int64) []types.MessageInfo {
	pageStr := strconv.FormatInt(page, 10)
	queryStr := `select m.message_id,m.message_text, to_char(m.message_date, 'DD Mon YYYY HH24:MI'), u.user_name
	from messages m
	inner join users u on u.user_id=m.message_userid
	order by m.message_date desc limit 20 offset ` + pageStr + ` ;`
	rows, err := dbConn.Query(queryStr)
	if err != nil {
		log.Println("sql.GetMessages err: ", err)
	}
	//messages := make([]*types.Message,0)
	messages := []types.MessageInfo{}
	for rows.Next() {
		message := &types.MessageInfo{}
		rows.Scan(&message.MessageID, &message.Text, &message.Date, &message.UserName)
		//log.Println("extracted contains: ", message)
		messages = append(messages, *message)
	}

	return messages
}
