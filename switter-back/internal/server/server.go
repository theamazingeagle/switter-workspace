package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"switter-back/internal/types"
)

type ServerConf struct {
	Addr string `json:"addr"`
}

type AuthDispatcher interface {
	Login(email, password string) (types.AuthInfo, error)
	Register(username, email, password string) (types.AuthInfo, error)
	Refresh(authInfo types.AuthInfo) (types.AuthInfo, error)
	Logout(authInfo types.AuthInfo) error
}
type MessageDispatcher interface{}
type UserDispatcher interface{}

type Server struct {
	authDispatcher    AuthDispatcher
	messageDispatcher MessageDispatcher
	userDispatcher    UserDispatcher
	HTTPServer        http.Server
}

func NewServer(conf ServerConf, authDispatcher AuthDispatcher, messageDispatcher MessageDispatcher, userDispatcher UserDispatcher) *Server {
	server := &Server{
		authDispatcher:    authDispatcher,
		messageDispatcher: messageDispatcher,
		userDispatcher:    userDispatcher,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", server.Login)
	mux.HandleFunc("/api/register", server.Register)
	mux.HandleFunc("/api/refreshtoken", server.Refresh)
	mux.HandleFunc("/api/logout", server.Logout)
	//mux.HandleFunc("/api/createmessage", server.CreateMessage)
	//mux.HandleFunc("/api/getmessages", server.GetMessageList)
	//mux = middleware.AccessMiddleWare(mux)
	server.HTTPServer = http.Server{Addr: conf.Addr, Handler: mux}
	return server
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(s.HTTPServer.Addr, s.HTTPServer.Handler))
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to read form data", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	}

	email := strings.TrimSpace(r.FormValue("userEmail"))
	password := strings.TrimSpace(r.FormValue("userPassword"))
	if len(email) == 0 && len(password) == 0 {
		log.Println("No form data", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	}

	AuthData, err := s.authDispatcher.Login(email, password)
	if err != nil {
		log.Println("Failed to register", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(AuthData)
	if err != nil {
		log.Println("Failed to marshall json", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	}

}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to read form data", err)
		return
	}

	email := strings.TrimSpace(r.FormValue("userEmail"))
	password := strings.TrimSpace(r.FormValue("userPassword"))
	username := strings.TrimSpace(r.FormValue("userName"))
	if len(username) == 0 && len(email) == 0 && len(password) == 0 {
		log.Println("No form data", err)
		return
	}

	AuthData, err := s.authDispatcher.Register(username, email, password)
	if err != nil {
		log.Println("Failed to register", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(AuthData)
	if err != nil {
		log.Println("Failed to marshall json", err)
		return
	}
}

func (s *Server) Refresh(w http.ResponseWriter, r *http.Request) {
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
		log.Println("Failed to marshall json", err)
		return
	}
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	authInfo := types.AuthInfo{}
	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		log.Println("Failed to read incoming data")
		return
	}
	err = s.authDispatcher.Logout(authInfo)
	if err != nil {
		log.Println("Failed to logout")
		return
	}
}

func (s *Server) CreateMessage(w http.ResponseWriter, r *http.Request) {
	message := &types.Message{}
	err := json.NewDecoder(r.Body).Decode(message)
	if err != nil {
		log.Println("router.CreateMessage error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("data parsing error,"))
		return
	}
	userID := r.Context().Value("UserID").(int)
	//log.Println("~router.CreateMessage: ", userID, text)
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
		w.Write([]byte("incomplete message data"))
		return
	}
}

func (s *Server) GetMessageList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte("invalid page index"))
		//return
		page = 0
	}
	messages := sql.GetMessages(page)

	js, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
