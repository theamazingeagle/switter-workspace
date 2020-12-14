package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"switter-back/internal/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthConf struct {
	JWTSigningKey string
	RTSigningKey  string
	Exptime       int
	SigningMethod string
	HashingCost   int
}

type AuthDispatcher struct {
}

//Claims - token fields
type Claims struct {
	jwt.StandardClaims // default files of JWT
	Email              string
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

type Claims struct {
	jwt.StandardClaims        // default files of JWT
	EXP                string `json:"exp"`
	IAT                string `json:"iat"`
	Email              string `json:"Email"`
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
func checkToken(accessTokenHeader string, signingKey []byte) error {
	// split "Bearer" and token
	splittedHeader := strings.Split(accessTokenHeader, " ")
	if len(splittedHeader) != 2 {
		return fmt.Errorf("invalid token")
	} else {
		jwToken := splittedHeader[1]
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

func getUserEmailFromJWT(JWTHeader string) (string, error) {
	splittedHeader := strings.Split(JWTHeader, " ")
	if len(splittedHeader) != 2 {
		return "", fmt.Errorf("invalid token")
	} else {
		jwToken := splittedHeader[1]
		tk := &Claims{}
		token, err := jwt.ParseWithClaims(jwToken, tk, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})
		if err != nil {
			return "", fmt.Errorf(" not possible to parse token: ", err)
		} else {
			if token.Valid {
				// tk = token.Claims.(*Claims)
				// log.Println("~ jwt ~ email ~ : ", tk.Email)
				return token.Claims.(*Claims).Email, nil
			} else {
				return "", fmt.Errorf("invalid token")
			}
		}
	}
}

//
