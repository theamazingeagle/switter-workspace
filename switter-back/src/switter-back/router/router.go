package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	//"strconv"
	/* ------ */
	sql "switter-back/sql"
	"switter-back/types"

	"github.com/dgrijalva/jwt-go"
)

var (
	expTime       = 3600
	signingKey    = []byte("mlp976g4bo76t6785gfv56")
	trustedRoutes = map[string]int{
		"/api/login":    0,
		"/api/register": 0,
	}
)

// Claims - token fields
type Claims struct {
	jwt.StandardClaims // default files of JWT
	Email              string
	Password           string
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

// checkToken returned nil if success
func checkToken(accessToken string, signingKey []byte) error {
	splittedToken := strings.Split(accessToken, " ")
	if len(splittedToken) != 2 {
		return fmt.Errorf("invalid token")
	} else {
		jwToken := splittedToken[1]
		tk := &Claims{}
		token, err := jwt.ParseWithClaims(jwToken, tk, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})
		if err != nil {
			return fmt.Errorf(" not possible to parse token: ", err)
		} else {
			if token.Valid {
				return nil
			} else {
				return fmt.Errorf("invalid token")
			}
		}
	}
	//return nil
}

//
func accessMiddleWare(handler http.Handler) http.Handler {
	log.Println("~router.accessMiddleWare ~~~~~")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* some checks? */
		route := r.URL.RequestURI()
		_, trustedRoutExist := trustedRoutes[route]
		if trustedRoutExist {
			handler.ServeHTTP(w, r)
		} else {
			//checkJWT
			authTokenHeader := r.Header.Get("Authorization")
			if len(authTokenHeader) > 0 {
				if checkToken(authTokenHeader, signingKey) == nil {
					handler.ServeHTTP(w, r)
				}
			} else {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("no token"))
			}
		}
	})
}

//
func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	} else {
		//log.Println("route:Login:FormValue: ", r.Form)
		email := strings.TrimSpace(r.FormValue("userEmail"))
		password := strings.TrimSpace(r.FormValue("userPassword"))
		//user := &types.User{}

		if len(email) > 0 && len(password) > 0 {
			user := sql.GetUserByEmail(email)
			//fmt.Printf("router.Login: user db data: %+v\n", user)
			if user != nil {
				if user.Email == email && user.Password == password { //?
					accessKey, err := generateAccessTokenString(email, password, signingKey)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("internal error"))
						return
					}
					authInfo := types.AuthInfo{
						JWT:       accessKey,
						UserID:    user.ID,
						UserName:  user.UserName,
						UserEmail: user.Email,
					}

					body, err := json.Marshal(authInfo)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.WriteHeader(http.StatusOK)
					w.Write(body)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("invalid auth data :)"))
					return
				}
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("incomplete auth data:"))
			return
		}
	}
}

//
func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
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
					// accessKey, err := generateAccessTokenString(email, password, signingKey)
					// if err != nil {
					// 	w.WriteHeader(http.StatusInternalServerError)
					// 	w.Write([]byte("generate access token error"))
					// } else {
					w.WriteHeader(http.StatusOK)
					//w.Write([]byte(accessKey))
					w.Write([]byte("User succefully registered"))
					//}
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("user exist"))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid auth data"))
			return
		}
	}
}

//
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("~router.CreateMessage: ")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("data parsing error"))
		return
	} else {
		userID, _ := strconv.Atoi(strings.TrimSpace(r.FormValue("userID")))
		text := r.FormValue("messageText")
		log.Println("~router.CreateMessage: ", userID, text)
		if userID != 0 && len(text) > 0 {
			err = sql.CreateMessage(text, userID)
			if err != nil {
				log.Println("router.CreateMessage() error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal error"))
				return
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("message created! ;-)"))
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no message data"))
			return
		}
	}
}

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

	//mux.HandleFunc("/api/", )
	mux.HandleFunc("/api/login", Login)
	mux.HandleFunc("/api/register", Register)
	mux.HandleFunc("/api/createmessage", CreateMessage)
	mux.HandleFunc("/api/getmessages", GetMessageList)
	//mux.HandleFunc("/api/refreshtoken", )
	requestHandler := accessMiddleWare(mux)
	//log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(port), nil))
	log.Fatal(http.ListenAndServe(":8080", requestHandler))
}
