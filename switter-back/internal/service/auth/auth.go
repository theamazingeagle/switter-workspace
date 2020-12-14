package auth

type AuthConf struct {
	JWTSigningKey string
	RTSigningKey  string
	Exptime       int
	SigningMethod string
	HashingCost   int
}
