package router
import(
	"net/http"
	"log"
	//"strconv"
	
)

func Start(host string, port int){

	http.HandleFunc("/api/getmessages", GetMessageList)
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
	//http.HandleFunc("/api/", )
		

	//log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(port), nil))
	log.Fatal(http.ListenAndServe(":8080", nil))
}