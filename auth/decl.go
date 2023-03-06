package auth

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Account interface {
	addNewUser(raw []byte) error
	verifyCredentials(raw []byte) error
	resetPassword(raw []byte) error
}
