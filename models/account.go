package models

type RegisterUser struct {
	Name, Email, Password, Username string 
}

type Login struct {
	ID, Password string 
}

type ResetPassword struct {
	Email, Username string
}

type DeleteUser struct {
	Email, Username, Password string 
}
