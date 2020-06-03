package router

import (
	sql "switter-back/sql"
	//"github.com/theamazingeagle/switter-back/types"
	/* ------ */
	"encoding/json"
	"net/http"
)

func Auth()                                             {}
func GetMessage(w http.ResponseWriter, r *http.Request) {}
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
