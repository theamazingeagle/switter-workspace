package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"switter-back/internal/types"

	"github.com/dgrijalva/jwt-go"
)

type Callback func(w http.ResponseWriter, r *http.Request)

type AuthDispatcher interface {
	Check(token string) bool
}

func AccessMiddleWare(JWTsigningKey string, handlerFunc Callback) Callback {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerParts := strings.Split(header, "Bearer ")
		if len(headerParts) < 2 {
			log.Println("No Token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("NoToken"))
			return
		}
		token := headerParts[1]

		if !check(token, JWTsigningKey) {
			log.Println("Unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("TokenExpired"))
		}
		userID, _ := getUserIDFromJWT(token, JWTsigningKey)
		newContext := context.WithValue(r.Context(), "UserID", userID)
		handlerFunc(w, r.WithContext(newContext))
		return
	}
}

func check(JWT, JWTSigningKey string) bool {
	tk := &types.Claims{}
	token, err := jwt.ParseWithClaims(JWT, tk, func(token *jwt.Token) (interface{}, error) {
		return JWTSigningKey, nil
	})
	if err != nil {
		log.Println(" not possible to parse token: ", err)
		return false
	}
	if token.Valid {
		log.Println("token invalid")
		return true
	}
	return false
}

func getUserIDFromJWT(JWTtoken, signingKey string) (int, error) {
	claims := &types.Claims{}
	token, err := jwt.ParseWithClaims(JWTtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return 0, fmt.Errorf(" not possible to parse token: ", err)
	}
	return token.Claims.(*types.Claims).UserID, nil
}
