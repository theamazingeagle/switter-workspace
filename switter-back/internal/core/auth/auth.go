package auth

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"switter-back/internal/types"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrJWTInvalid             = errors.New("JWT is invalid")
	ErrRTInvalid              = errors.New("Refresh token is invalid")
	ErrJWTNotCreated          = errors.New("JWT not created")
	ErrRTNotCreated           = errors.New("Refresh token not created")
	ErrNoData                 = errors.New("No data")
	ErrNotParse               = errors.New("Failed to parse token")
	ErrUserExist              = errors.New("User exist")
	ErrUserNotFound           = errors.New("User not found")
	ErrUserNotCreated         = errors.New("User not created")
	ErrUserNotGet             = errors.New("Failed to get user")
	ErrPasswordMatch          = errors.New("Password not match")
	ErrRefreshTokenNotDeleted = errors.New("Refresh token not deleted")
)

type Storage interface {
	GetUserByEmail(email string) (*types.User, bool, error)
	CreateUser(username, password, email string) (types.User, error)
	DeleteRefreshToken(userID types.UserID) error
}

type AuthDispatcher struct {
	conf    AuthConf
	storage Storage
}

type AuthConf struct {
	JWTSigningKey string `json:"jwtsignkey"`
	RTSigningKey  string `json:"rtsignkey"`
	Exptime       int    `json:"exptime"`
	SigningMethod string `json:"signmethod"`
	HashingCost   int    `json:"hashingcost"`
}

func New(conf AuthConf, storage Storage) *AuthDispatcher {
	return &AuthDispatcher{conf: conf, storage: storage}
}

func (a *AuthDispatcher) Login(email, password string) (types.AuthInfo, error) {
	user, exist, err := a.storage.GetUserByEmail(email)
	if !exist {
		return types.AuthInfo{}, ErrUserNotFound
	}
	if user.Password != password {
		return types.AuthInfo{}, ErrPasswordMatch
	}
	jwt, err := makeJWT(*user, a.conf.JWTSigningKey, a.conf.Exptime)
	if err != nil {
		return types.AuthInfo{}, err
	}
	rt, err := makeRefreshToken(jwt, a.conf.RTSigningKey)
	if err != nil {
		return types.AuthInfo{}, err
	}
	authInfo := types.AuthInfo{
		JWT: jwt,
		RT:  rt,
	}
	return authInfo, nil
}

func (a *AuthDispatcher) Register(username, email, password string) (types.AuthInfo, error) {
	_, exist, err := a.storage.GetUserByEmail(email)

	if exist {
		log.Println("User exist")
		return types.AuthInfo{}, ErrUserExist
	} else if err != nil {
		log.Println("Quering error")
		return types.AuthInfo{}, ErrUserNotFound
	}
	user, err := a.storage.CreateUser(username, password, email)
	if err != nil {
		log.Println("Failed to create user")
		return types.AuthInfo{}, ErrUserNotCreated
	}
	jwt, err := makeJWT(user, a.conf.JWTSigningKey, a.conf.Exptime)
	if err != nil {
		log.Println("Failed to make JWT")
		return types.AuthInfo{}, ErrJWTNotCreated
	}
	rt, err := makeRefreshToken(jwt, a.conf.RTSigningKey)
	if err != nil {
		log.Println("Failed to make refresh token")
		return types.AuthInfo{}, ErrRTNotCreated
	}
	authInfo := types.AuthInfo{
		JWT: jwt,
		RT:  rt,
	}
	return authInfo, nil
}

func (a *AuthDispatcher) Refresh(authInfo types.AuthInfo) (types.AuthInfo, error) {
	// check jwt signing
	tk := &types.Claims{}
	token, err := jwt.ParseWithClaims(authInfo.JWT, tk, func(token *jwt.Token) (interface{}, error) {
		return a.conf.JWTSigningKey, nil
	})
	if err != nil {
		log.Println(" not possible to parse token: ", err)
		return types.AuthInfo{}, ErrNotParse
	}
	// check signature
	parts := strings.Split(authInfo.JWT, ".")
	err = token.Method.Verify(strings.Join(parts[0:2], "."), token.Signature, a.conf.JWTSigningKey)
	if err != nil {
		log.Println("JWT's signature is invalid", err)
		return types.AuthInfo{}, ErrJWTInvalid
	}
	// check rt
	user, exist, err := a.storage.GetUserByEmail(tk.Email)
	if err != nil {
		log.Println("Failed to get refresh token", err)
		return types.AuthInfo{}, ErrNoData
	}
	if !exist {
		log.Println("User not found", err)
		return types.AuthInfo{}, ErrUserNotFound
	}
	if user.RT != authInfo.RT {
		log.Println("Invalid refresh token", err)
		return types.AuthInfo{}, ErrRTInvalid
	}
	// TO DO: check jwt and rt equality

	// make jwt
	authInfo.JWT, err = makeJWT(*user, a.conf.JWTSigningKey, a.conf.Exptime)
	if err != nil {
		log.Println("JWT not created: ", err)
		return types.AuthInfo{}, ErrJWTNotCreated
	}
	//make rt by jwt
	authInfo.RT, err = makeRefreshToken(authInfo.JWT, a.conf.RTSigningKey)
	if err != nil {
		log.Println("Refresh token not created: ", err)
		return types.AuthInfo{}, ErrRTNotCreated
	}

	return authInfo, nil
}

// Logout - set rt to null
func (a *AuthDispatcher) Logout(userID types.UserID) error {
	err := a.storage.DeleteRefreshToken(userID)
	if err != nil {
		log.Println(" not possible to parse token: ", err)
		return ErrRefreshTokenNotDeleted
	}
	return nil
}

func makeJWT(user types.User, signingKey string, expTime int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &types.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expTime) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email:  user.Email,
		UserID: user.ID,
	})
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		log.Println("Failed to make JWT: ", err)
		return "", ErrJWTNotCreated
	}
	return tokenString, nil
}

func makeRefreshToken(jwt, rtSecrerKey string) (string, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) < 3 {
		return "", fmt.Errorf("Invalid JWT")
	}
	t := rtSecrerKey + parts[2][:8] + strconv.Itoa(int(time.Now().UnixNano()))
	hasher := sha512.New()
	_, err := hasher.Write([]byte(t))
	if err != nil {
		return "", fmt.Errorf("Failed to create hash")
	}
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil
}
