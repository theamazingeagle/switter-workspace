package types

type User struct {
	ID       int64  `json:"ID"`
	UserName string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	RT       string `json:"RT"`
}

type Message struct {
	ID     int64  `json:"ID"`
	Url    string `json:"Url"`
	UserID int64  `json:"UserID"`
	Date   string `json:"Date"`
	//RootID int64 `json:RootID`
	Text string `json:"Text"`
}
type MessageInfo struct {
	MessageID int64  `json:"ID"`
	Text      string `json:"Text"`
	Date      string `json:"Date"`
	UserName  string `json:"Username"`
}
type AuthInfo struct {
	JWT       string `json:"JWT"`
	UserID    int64  `json:"UserID"`
	UserName  string `json:"Username"`
	UserEmail string `json:"UserEmail"`
}
type NewMessage struct {
	Text   string `json:"Text"`
	UserID int    `json:"UserID"`
}
