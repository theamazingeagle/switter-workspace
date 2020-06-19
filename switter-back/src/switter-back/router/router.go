package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	//"strconv"
	/* ------ */
	sql "switter-back/sql"
	"switter-back/types"

	"github.com/dgrijalva/jwt-go"
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
		trustedRoutes := []string{
			"/api/login",
			"/api/register",
		}
		for _, route := range trustedRoutes {
			if route == r.URL.RequestURI() {
				handler.ServeHTTP(w, r)
			}
		}

	})
}

// generateAccessTokenString return string.string.string JWT token
func generateAccessTokenString(userEmail string, userPassword string, signingKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expTime) * time.Second).Unix(),
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
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("no data"))
	} else {
		//log.Println("route:Login:FormValue: ", r.Form)
		email := strings.TrimSpace(r.FormValue("userEmail"))
		password := strings.TrimSpace(r.FormValue("userPassword"))
		user := &types.User{}

		if len(email) > 0 && len(password) > 0 {
			user = sql.GetUserByEmail(email)
		}
		if user != nil {
			if user.Email == email && user.Password == password {
				accessKey, err := generateAccessTokenString(email, password, signingKey)
				if err != nil {
					w.Write([]byte("internal error"))
				}
				w.Write([]byte(accessKey))
			}
		}
	}
}

//
func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("no data"))
	} else {
		//log.Println("route:Login:FormValue: ", r.Form)
		email := strings.TrimSpace(r.FormValue("userEmail"))
		password := strings.TrimSpace(r.FormValue("userPassword"))
		username := strings.TrimSpace(r.FormValue("userName"))
		if len(email) > 0 && len(password) > 0 && len(username) > 0 {
			if sql.GetUserByEmail(email) == nil {
				if sql.CreateUser(username, password, email) != nil {
					w.Write([]byte("internal error"))
				} else {
					//send accessToken
					accessKey, err := generateAccessTokenString(email, password, signingKey)
					if err != nil {
						w.Write([]byte("generate access token error"))
					}
					w.Write([]byte(accessKey))
				}
			} else {
				w.Write([]byte("user exist"))
			}
		}
	}
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
	mux.HandleFunc("/api/register", Register)
	//mux.HandleFunc("/api/refreshtoken", )
	requestHandler := accessMiddleWare(mux)
	//log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(port), nil))
	log.Fatal(http.ListenAndServe(":8080", requestHandler))
}
