package router

import (
	"log"
	"net/http"

	//"strconv"
	"fmt"
	sql "switter-back/sql"
	"time"

	"github.com/dgrijalva/jwt-go"

	//"github.com/theamazingeagle/switter-back/types"
	/* ------ */
	"encoding/json"
)

var (
	expTime    = 3600
	signingKey = []byte("mlp976g4bo76t6785gfv56")
)

// Claims - token fields
type Claims struct {
	jwt.StandardClaims // default files of JWT
	Email              string
	Password           string
}

func accessMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* some checks? */
		handler.ServeHTTP(w, r)
	})
}

// generateAccessTokenString return string.string.string JWT token
func generateAccessTokenString(userEmail string, userPassword string, signingKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expTime) * time.Millisecond).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email:    userEmail,
		Password: userPassword,
	})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("router.generateAccessTokenString error: %w", err)
	}
	return tokenString, nil
}

//
func Login(w http.ResponseWriter, r *http.Request) {

	log.Printf("%+v", r.Form)
	//w.Write([]byte(tokenString))
	//return //What???
}

//
func Register(w http.ResponseWriter, r *http.Request) {

}

//
func RefreshToken(w http.ResponseWriter, r *http.Request) {

}

//
func GetMessage(w http.ResponseWriter, r *http.Request) {}

//
func GetMessageList(w http.ResponseWriter, r *http.Request) {
	messages := sql.GetMessages()
	js, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Start() - init and start serving requests
func Start(host string, port int) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/getmessages", GetMessageList)
	//mux.HandleFunc("/api/", )
	mux.HandleFunc("/api/login", Login)
	//mux.HandleFunc("/api/refreshtoken", )
	requestHandler := accessMiddleWare(mux)
	//log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(port), nil))
	log.Fatal(http.ListenAndServe(":8080", requestHandler))
}
