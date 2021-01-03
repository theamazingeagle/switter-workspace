package types

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RT       string `json:"refresh_token"`
}

type Message struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Date   string `json:"date"`
	Text   string `json:"text"`
}

type AuthInfo struct {
	JWT string `json:"jwt"`
	RT  string `json:"refresh_token"`
}

type Claims struct {
	jwt.StandardClaims
	Email  string `json:"Email"`
	UserID int    `json:"UserID"`
}
