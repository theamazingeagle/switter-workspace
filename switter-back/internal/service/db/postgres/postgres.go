package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"switter-back/internal/types"

	_ "github.com/lib/pq"
)

type PostgresConf struct {
	HostName   string `json:"hostname"`
	DriverName string `json:"drivername"`
	DBName     string `json:"dbname"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Port       int    `json:"port"`
}

type Postgres struct {
	conn *sql.DB
}

func NewPostgres(conf PostgresConf) (*Postgres, error) {
	postgres := &Postgres{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.HostName, conf.Port, conf.UserName, conf.Password, conf.DBName)
	var err error
	postgres.conn, err = sql.Open(conf.DriverName, dsn)
	if err != nil {
		log.Println("NewPostgres err: ", err)
		return &Postgres{}, err
	}
	return postgres, nil
}

// CloseConn ...
func CloseConn() {
	//close(dbConn)
}

//

func (p *Postgres) CreateUser(userName, password, email string) error {
	_, err := p.conn.Exec("INSERT INTO users(user_name,user_password,user_email) VALUES($1, $2, $3)", userName, password, email)
	if err != nil {
		log.Println("sql.CreateUser err: ", err)
		return fmt.Errorf("sql.CreateUser Error: ", err)
	}
	return nil
}
func (p *Postgres) GetUser(ID int64) *types.User {
	row := p.conn.QueryRow("SELECT * FROM users WHERE user_id=$1", ID)
	if row == nil {
		log.Println("sql.GetUser err: ", row)
	}
	user := &types.User{}
	row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email, &user.RT)
	log.Println("sql.GetUser result: ", user)
	return user
}

func (p *Postgres) GetUserByEmail(email string) *types.User {
	row := p.conn.QueryRow("SELECT * FROM users WHERE user_email=$1", email)
	if row == nil {
		log.Println("sql.GetUserByEmail err: ", row)
		return nil
	}
	user := &types.User{}
	// user_id | user_name | user_email | user_password
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.RT)
	if err != nil {
		log.Println("sql.GetUserByEmail error, no row: ", err)
		return nil
	}
	log.Println("sql.GetUserByEmail result: ", user)
	return user
}

func (p *Postgres) UpdateUserName(ID int64, newName string) {
	_, err := p.conn.Exec("UPDATE users SET user_name = $1 WHERE user_id = $2", newName, ID)
	if err != nil {
		log.Println("sql.UpdateUserName err: ", err)
	}
}
func (p *Postgres) UpdateUserPassword(ID int64, newPass string) {
	_, err := p.conn.Exec("UPDATE users SET user_password = $1 WHERE user_id = $2", newPass, ID)
	if err != nil {
		log.Println("sql.UpdateUserPassword err: ", err)
	}
}
func (p *Postgres) UpdateUserEmail(ID int64, newEmail string) {
	_, err := p.conn.Exec("UPDATE users SET user_email = $1 WHERE user_id = $2", newEmail, ID)
	if err != nil {
		log.Println("sql.UpdateUserEmail err: ", err)
	}
}
func (p *Postgres) DeleteUser(ID int64) {
	_, err := p.conn.Exec("DELETE  FROM users WHERE user_id=$1", ID)
	if err != nil {
		log.Println("sql.DeleteUser err: ", err)
	}
}

func (p *Postgres) CreateMessage(text string, userID int) error {
	_, err := p.conn.Exec(`INSERT INTO messages(message_text, message_userid)
						  VALUES($1, $2);`,
		text, userID)
	if err != nil {
		return fmt.Errorf("sql.CreateUser err: ", err)
	}
	return nil
}

func (p *Postgres) GetMessage(messageID string) types.Message {
	row := p.conn.QueryRow("SELECT * FROM messages WHERE message_id=$1 ;", messageID)
	if row == nil {
		log.Println("sql.GetMessage err: ", row)
	}
	message := &types.Message{}
	row.Scan(&message.ID, &message.Url, &message.UserID, &message.Date, &message.Text)
	return *message
}

func (p *Postgres) UpdateMessage(newText string, userID int) {
	_, err := p.conn.Exec("UPDATE messages SET message_text = $1 WHERE user_id = $2", newText, userID)
	if err != nil {
		log.Println("sql.UpdateMessage err: ", err)
	}
}

func (p *Postgres) DeleteMessage(ID int64) {
	_, err := p.conn.Exec("DELETE FROM messages WHERE message_id=$1", ID)
	if err != nil {
		log.Println("sql.DeleteMessage err: ", err)
	}
}

// GetMessages() returns all messages
func (p *Postgres) GetMessages(page int64) []types.MessageInfo {
	pageStr := strconv.FormatInt(page, 10)
	queryStr := `select m.message_id,m.message_text, to_char(m.message_date, 'DD Mon YYYY HH24:MI'), u.user_name
	from messages m
	inner join users u on u.user_id=m.message_userid
	order by m.message_date desc limit 20 offset ` + pageStr + ` ;`
	rows, err := p.conn.Query(queryStr)
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
