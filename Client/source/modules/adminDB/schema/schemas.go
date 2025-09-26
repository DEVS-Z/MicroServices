package schema

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BackupPartialParams struct {
	Credentials Credentials `json:"credentials"`
	Tables      []string    `json:"tables"`
}