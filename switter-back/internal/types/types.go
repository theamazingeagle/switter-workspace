package types

import "github.com/dgrijalva/jwt-go"

type UserID int64

type User struct {
	ID       UserID `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RT       string `json:"refresh_token"`
}

type MessageID int64

type Message struct {
	ID     MessageID `json:"id"`
	UserID UserID    `json:"user_id"`
	Date   string    `json:"date"`
	Text   string    `json:"text"`
}

type FullMessageData struct {
	ID       MessageID `json:"id"`
	Text     string    `json:"text"`
	Date     string    `json:"date"`
	UserName string    `json:"username"`
	UserID   UserID    `json:"user_id"`
}

type AuthInfo struct {
	JWT string `json:"jwt"`
	RT  string `json:"refresh_token"`
}

type Claims struct {
	jwt.StandardClaims
	Email  string `json:"Email"`
	UserID UserID `json:"UserID"`
}
