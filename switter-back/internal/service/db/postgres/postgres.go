package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"switter-back/internal/types"

	_ "github.com/lib/pq"
)

var (
	ErrNotFound  = errors.New("Not found")
	ErrQueryExec = errors.New("Failed to exec query")
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

// TODO
func CloseConn() {
}

func (p *Postgres) CreateUser(username, password, email string) error {
	_, err := p.conn.Exec("INSERT INTO users(username,password,email) VALUES($1, $2, $3)", username, password, email)
	if err != nil {
		log.Println("postgres.CreateUser err: ", err)
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) GetUserByID(ID int64) (*types.User, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM users WHERE user_id=$1", ID)
	user := &types.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email, &user.RT)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, ErrNotFound
		}
		log.Println("sql.GetUser err: ", row)
		return user, false, ErrQueryExec
	}
	return user, true, nil
}

func (p *Postgres) GetUserByEmail(email string) (*types.User, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM users WHERE email=$1", email)
	user := &types.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.RT)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, ErrNotFound
		}
		log.Println("sql.GetUserByEmail error, no row: ", err)
		return user, false, ErrQueryExec
	}
	log.Println("sql.GetUserByEmail result: ", user)
	return user, true, nil
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

func (p *Postgres) DeleteRefreshTokenByEmail(email string) error {
	_, err := p.conn.Exec("UPDATE users SET refresh_token = NULL WHERE user_email = $1", email)
	if err != nil {
		log.Println("Failed to delete refresh token: ", err)
	}
	return nil
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
	row.Scan(&message.ID, &message.UserID, &message.Date, &message.Text)
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

func (p *Postgres) GetMessages(page int64) []types.Message {
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
	messages := []types.Message{}
	for rows.Next() {
		message := &types.Message{}
		rows.Scan(&message.ID, &message.Text, &message.Date, &message.UserID)
		//log.Println("extracted contains: ", message)
		messages = append(messages, *message)
	}
	return messages
}
