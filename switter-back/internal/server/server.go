package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"switter-back/internal/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var (
	ErrTokenInvalid = errors.New("Token invalid")
)

type ServerConf struct {
	Addr          string `json:"addr"`
	JWTSigningKey string `json:"jwtsignkey"`
}
type AuthCred struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Callback func(w http.ResponseWriter, r *http.Request)

type AuthDispatcher interface {
	Login(email, password string) (types.AuthInfo, error)
	Register(username, email, password string) (types.AuthInfo, error)
	Refresh(authInfo types.AuthInfo) (types.AuthInfo, error)
	Logout(userID types.UserID) error
}

type MessageDispatcher interface {
	GetListPage(page int) ([]types.FullMessageData, error)
	GetMessage(msgID types.MessageID) (types.Message, error)
	CreateMessage(userID types.UserID, message string) error
	UpdateMessage(userID types.UserID, msgID types.MessageID, message string) error
	DeleteMessage(userID types.UserID, msgID types.MessageID) error
}
type UserDispatcher interface{}

type Server struct {
	conf              ServerConf
	authDispatcher    AuthDispatcher
	messageDispatcher MessageDispatcher
	userDispatcher    UserDispatcher
	HTTPServer        http.Server
}

func NewServer(conf ServerConf, authDispatcher AuthDispatcher, messageDispatcher MessageDispatcher) *Server {
	server := &Server{
		conf:              conf,
		authDispatcher:    authDispatcher,
		messageDispatcher: messageDispatcher,
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/api/auth/login", server.login)
	mux.HandleFunc("/api/auth/register", server.register)
	mux.HandleFunc("/api/auth/refreshtoken", server.refresh)
	mux.HandleFunc("/api/auth/logout", server.logout)
	mux.HandleFunc("/api/message/create", server.accessMiddleWare(server.createMessage))
	mux.HandleFunc("/api/message/all", server.accessMiddleWare(server.getMessageList))
	mux.HandleFunc("/api/message/delete/{id:[0-9]+}", server.accessMiddleWare(server.deleteMessage))

	server.HTTPServer = http.Server{Addr: conf.Addr, Handler: mux}
	return server
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(s.HTTPServer.Addr, s.HTTPServer.Handler))
}

func (s *Server) accessMiddleWare(handlerFunc Callback) Callback {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerParts := strings.Split(header, "Bearer ")
		if len(headerParts) < 2 {
			log.Println("No token")
			sendMessage(w, http.StatusUnauthorized, "No token")
			return
		}
		token := headerParts[1]
		if !check(token, s.conf.JWTSigningKey) {
			log.Println("Token expired")
			sendMessage(w, http.StatusUnauthorized, "Token expired")
			return
		}
		userID, err := getUserIDFromJWT(token, s.conf.JWTSigningKey)
		if err != nil {
			log.Println("failed to get User ID from JWT, ID: ", userID)
			return
		}

		newContext := context.WithValue(r.Context(), "UserID", userID)
		handlerFunc(w, r.WithContext(newContext))
		return
	}
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	authCred := AuthCred{}
	err := json.NewDecoder(r.Body).Decode(&authCred)
	if err != nil {
		log.Println("Bad auth data", err)
		sendMessage(w, http.StatusBadRequest, "Bad auth data")
		return
	}

	if len(authCred.Email) == 0 && len(authCred.Password) == 0 {
		log.Println("No auth data", err)
		sendMessage(w, http.StatusUnauthorized, "No auth data")
		return
	}

	AuthData, err := s.authDispatcher.Login(authCred.Email, authCred.Password)
	if err != nil {
		log.Println("Failed to authenticate", err)
		sendMessage(w, http.StatusBadRequest, "Failed to authenticate")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(AuthData)
	if err != nil {
		log.Println("Failed to make answer", err)
		sendMessage(w, http.StatusBadRequest, "Failed to make answer")
		return
	}
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	authCred := AuthCred{}
	err := json.NewDecoder(r.Body).Decode(&authCred)
	if err != nil {
		log.Println("Bad auth data", err)
		sendMessage(w, http.StatusBadRequest, "Bad auth data")
		return
	}

	if len(authCred.Username) == 0 && len(authCred.Email) == 0 && len(authCred.Password) == 0 {
		log.Println("No form data", err)
		sendMessage(w, http.StatusUnauthorized, "No form data")
		return
	}

	AuthData, err := s.authDispatcher.Register(authCred.Username, authCred.Email, authCred.Password)
	if err != nil {
		log.Println("Failed to register", err)
		sendMessage(w, http.StatusInternalServerError, "No form data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(AuthData)
	if err != nil {
		log.Println("Failed to make answer", err)
		sendMessage(w, http.StatusInternalServerError, "Failed to make answer")
		return
	}
}

func (s *Server) refresh(w http.ResponseWriter, r *http.Request) {
	authInfo := types.AuthInfo{}
	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		log.Println("Failed to read incoming data")
		return
	}
	authInfo, err = s.authDispatcher.Refresh(authInfo)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&authInfo)
	if err != nil {
		log.Println("Failed to marshall answer", err)
		return
	}
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("UserID").(types.UserID)
	err := s.authDispatcher.Logout(userID)
	if err != nil {
		log.Println("Failed to logout")
		sendMessage(w, http.StatusInternalServerError, "Failed to logout")
		return
	}
	r.Header.Set("Authorization", "")
	sendMessage(w, http.StatusOK, "log out")
}

func (s *Server) createMessage(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("UserID").(types.UserID)
	message := &types.Message{}
	err := json.NewDecoder(r.Body).Decode(message)
	if err != nil {
		log.Println("Failed to read json, error: ", err)
		sendMessage(w, http.StatusInternalServerError, "internal error")
		return
	}
	if userID == 0 && len(message.Text) == 0 {
		log.Println("Empty data, error: ")
		sendMessage(w, http.StatusBadRequest, "incomplete message data")
	}

	err = s.messageDispatcher.CreateMessage(userID, message.Text)
	if err != nil {
		log.Println("router.CreateMessage() error: ", err)
		sendMessage(w, http.StatusInternalServerError, "internal error")
		return
	}
	sendMessage(w, http.StatusCreated, "message created")
}

func (s *Server) getMessageList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 0
	}

	messages, err := s.messageDispatcher.GetListPage(int(page))
	if err != nil {
		sendMessage(w, http.StatusInternalServerError, "failed to get data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		log.Println("Failed to marshall answer", err)
		return
	}
}

func (s *Server) deleteMessage(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("UserID").(types.UserID)
	messageID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("invalid message id")
		sendMessage(w, http.StatusBadRequest, "invalid message id")
		return
	}
	err = s.messageDispatcher.DeleteMessage(userID, types.MessageID(messageID))
	if err != nil {
		log.Println("failed to delete message", err)
		sendMessage(w, http.StatusInternalServerError, "failed to delete message")
		return
	}
	sendMessage(w, http.StatusOK, "deleted")
}

func check(JWT, JWTSigningKey string) bool {
	tk := &types.Claims{}
	token, err := jwt.ParseWithClaims(JWT, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSigningKey), nil
	})
	if err != nil {
		log.Println("Not possible to parse token: ", err)
		return false
	}
	if !token.Valid {
		log.Println("Token invalid")
		return false
	}
	return true
}

func getUserIDFromJWT(JWTtoken, signingKey string) (types.UserID, error) {
	claims := &types.Claims{}
	token, err := jwt.ParseWithClaims(JWTtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		log.Println("Token invalid")
		return 0, ErrTokenInvalid
	}
	return token.Claims.(*types.Claims).UserID, nil
}

func sendMessage(w http.ResponseWriter, status interface{}, message string) {
	w.WriteHeader(status.(int))
	w.Write([]byte(message))
}
