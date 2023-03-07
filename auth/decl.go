package auth

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Action   int    `json:"action"`
}

type Account interface {
	ParseCredentials([]byte) error
}
