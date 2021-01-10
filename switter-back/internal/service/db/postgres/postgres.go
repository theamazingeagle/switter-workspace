package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func (p *Postgres) CreateUser(username, password, email string) (types.User, error) {
	row := p.conn.QueryRow("INSERT INTO users(username,password,email) VALUES($1, $2, $3) RETURNING *", username, password, email)
	user := types.User{}
	err := row.Scan(&user)
	if err != nil {
		log.Println("postgres.CreateUser err: ", err)
		return types.User{}, ErrQueryExec
	}
	return user, nil
}

func (p *Postgres) GetUserByID(ID int64) (*types.User, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM users WHERE id=$1", ID)
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
			return user, false, nil
		}
		log.Println("sql.GetUserByEmail error, no row: ", err)
		return user, false, ErrQueryExec
	}
	return user, true, nil
}

func (p *Postgres) UpdateUserName(ID int64, newName string) {
	_, err := p.conn.Exec("UPDATE users SET username = $1 WHERE id = $2", newName, ID)
	if err != nil {
		log.Println("sql.UpdateUserName err: ", err)
	}
}
func (p *Postgres) UpdateUserPassword(ID int64, newPass string) {
	_, err := p.conn.Exec("UPDATE users SET password = $1 WHERE id = $2", newPass, ID)
	if err != nil {
		log.Println("sql.UpdateUserPassword err: ", err)
	}
}
func (p *Postgres) UpdateUserEmail(ID int64, newEmail string) {
	_, err := p.conn.Exec("UPDATE users SET email = $1 WHERE id = $2", newEmail, ID)
	if err != nil {
		log.Println("sql.UpdateUserEmail err: ", err)
	}
}
func (p *Postgres) DeleteUser(ID int64) {
	_, err := p.conn.Exec("DELETE  FROM users WHERE id=$1", ID)
	if err != nil {
		log.Println("sql.DeleteUser err: ", err)
	}
}

func (p *Postgres) DeleteRefreshToken(userID types.UserID) error {
	_, err := p.conn.Exec("UPDATE users SET refresh_token = NULL WHERE id = $1", userID)
	if err != nil {
		log.Println("Failed to delete refresh token: ", err)
	}
	return nil
}

func (p *Postgres) CreateMessage(userID types.UserID, text string) error {
	_, err := p.conn.Exec(`INSERT INTO messages(msg, user_id)
						  VALUES($1, $2);`,
		text, userID)
	if err != nil {
		log.Println("Failed to create message")
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) GetMessage(messageID types.MessageID) (types.Message, bool, error) {
	row := p.conn.QueryRow("SELECT * FROM messages WHERE id=$1 ;", messageID)
	message := types.Message{}
	err := row.Scan(&message.ID, &message.UserID, &message.Date, &message.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("sql.GetMessage err: ", err)
			return types.Message{}, false, ErrNotFound
		}
		log.Println("sql.GetMessage err: ", err)
		return types.Message{}, false, ErrQueryExec
	}
	return message, true, nil
}

func (p *Postgres) GetMessageList(page int) ([]types.FullMessageData, error) {
	rows, err := p.conn.Query(
		`select m.id, m.msg, to_char(m.msg_date, 'DD Mon YYYY HH24:MI'), u.username, m.user_id
		from messages m
		inner join users u on u.id=m.user_id
		order by m.msg_date desc limit 20 offset $1`, page)
	if err != nil {
		log.Println("sql.GetMessages err: ", err)
		return []types.FullMessageData{}, ErrQueryExec
	}
	messages := []types.FullMessageData{}
	for rows.Next() {
		message := types.FullMessageData{}
		rows.Scan(&message.ID, &message.Text, &message.Date, &message.UserName, &message.UserID)
		messages = append(messages, message)
	}
	return messages, nil
}

func (p *Postgres) UpdateMessage(ID types.MessageID, newText string) error {
	_, err := p.conn.Exec("UPDATE messages SET msg = $1 WHERE id = $2", newText, ID)
	if err != nil {
		log.Println("sql.UpdateMessage err: ", err)
		return ErrQueryExec
	}
	return nil
}

func (p *Postgres) DeleteMessage(userID types.UserID, ID types.MessageID) error {
	_, err := p.conn.Exec("DELETE FROM messages WHERE id=$1", ID)
	if err != nil {
		log.Println("sql.DeleteMessage err: ", err)
		return ErrQueryExec
	}
	return nil
}
