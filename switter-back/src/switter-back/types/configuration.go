package types

type SqlConfiguration struct {
	HostName string `json:"hostname"`
	DriverName string `json:"drivername"`
	DBName string `json:"dbname"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Port       int16  `json:"port"`
}

//--------------------------------
type AppConfiguration struct {
	Host string           `json:"host"`
	Port int           `json:"port"`
	SQL  SqlConfiguration `json:"sql"`
}