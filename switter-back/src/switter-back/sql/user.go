package sql

import (
	"fmt"
	"log"
	"switter-back/types"
)

// CreateUser
func CreateUser(userName, password, email string) error {
	_, err := dbConn.Exec("INSERT INTO users(user_name,user_password,user_email) VALUES($1, $2, $3)", userName, password, email)
	if err != nil {
		log.Println("sql.CreateUser err: ", err)
		return fmt.Errorf("sql.CreateUser Error: ", err)
	}
	return nil
}
func GetUser(ID int64) *types.User {
	row := dbConn.QueryRow("SELECT * FROM users WHERE user_id=$1", ID)
	if row == nil {
		log.Println("sql.GetUser err: ", row)
	}
	user := &types.User{}
	row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	log.Println("sql.GetUser result: ", user)
	return user
}

func GetUserByEmail(email string) *types.User {
	row := dbConn.QueryRow("SELECT * FROM users WHERE user_email=$1", email)
	if row == nil {
		log.Println("sql.GetUserByEmail err: ", row)
		return nil
	}
	user := &types.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	if err != nil {
		log.Println("sql.GetUserByEmail error, no row: ", err)
		return nil
	}
	log.Println("sql.GetUserByEmail result: ", user)
	return user
}

func UpdateUserName(ID int64, newName string) {
	_, err := dbConn.Exec("UPDATE users SET user_name = $1 WHERE user_id = $2", newName, ID)
	if err != nil {
		log.Println("sql.UpdateUserName err: ", err)
	}
}
func UpdateUserPassword(ID int64, newPass string) {
	_, err := dbConn.Exec("UPDATE users SET user_password = $1 WHERE user_id = $2", newPass, ID)
	if err != nil {
		log.Println("sql.UpdateUserPassword err: ", err)
	}
}
func UpdateUserEmail(ID int64, newEmail string) {
	_, err := dbConn.Exec("UPDATE users SET user_email = $1 WHERE user_id = $2", newEmail, ID)
	if err != nil {
		log.Println("sql.UpdateUserEmail err: ", err)
	}
}
func DeleteUser(ID int64) {
	_, err := dbConn.Exec("DELETE  FROM users WHERE user_id=$1", ID)
	if err != nil {
		log.Println("sql.DeleteUser err: ", err)
	}
}
