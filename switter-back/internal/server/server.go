package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"switter-back/internal/types"
)

type ServerConf struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type AuthDispatcher interface{}
type MessageDispatcher interface{}
type UserDispatcher interface{}

type Server struct {
	authDispatcher    AuthDispatcher
	messageDispatcher MessageDispatcher
	userDispatcher    UserDispatcher
	server            http.Server
}

func NewServer(conf ServerConf) *Server {
	return &Server{}
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	}
}

//
func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no auth data"))
		return
	}
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {}

//
func (s *Server) CreateMessage(w http.ResponseWriter, r *http.Request) {
	//log.Println("~router.CreateMessage: ")
	//err := r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	message := &types.NewMessage{}
	err := decoder.Decode(message)
	// ------
	if err != nil {
		log.Println("router.CreateMessage error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("data parsing error,"))
		return
	} else {
		//userID, _ := strconv.Atoi(strings.TrimSpace(r.FormValue("userID")))
		//text := r.FormValue("messageText")
		userEmail, _ := getUserEmailFromJWT(r.Header.Get("Authorization"))
		user := sql.GetUserByEmail(userEmail)
		userID := int(user.ID)
		text := message.Text

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

func (s *Server) Run() {
	mux := http.NewServeMux()

	//mux.HandleFunc("/api/", )
	//mux.HandleFunc("/api/login", Login)
	//mux.HandleFunc("/api/register", Register)
	mux.HandleFunc("/api/createmessage", CreateMessage)
	mux.HandleFunc("/api/getmessages", GetMessageList)
	//mux.HandleFunc("/api/refreshtoken", )
	requestHandler := accessMiddleWare(mux)
	//log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(port), nil))
	log.Fatal(http.ListenAndServe(":8080", requestHandler))
}
