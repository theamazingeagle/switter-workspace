package types

type User struct {
	ID int64 `json:"ID"`
	UserName string `json:"Username"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

type Message struct {
	ID int64 `json:"ID"` 
	Url string `json:"Url"`
	UserID int64 `json:"UserID"`
	Date string `json:Date`
	//RootID int64 `json:RootID`
	Text string `json:Text`
}