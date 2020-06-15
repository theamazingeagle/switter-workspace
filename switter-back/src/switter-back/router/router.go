package router

import (
	"log"
	"net/http"
	//"strconv"
)

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
