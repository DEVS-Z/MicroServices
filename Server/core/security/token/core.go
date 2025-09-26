package token

type TokenData struct {
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
	Rol      int    `json:"rol"`
}
