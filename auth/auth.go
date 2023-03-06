package auth

import (
	"encoding/json"
	"fmt"
	"server/auth/keygen"
	sqlconnection "server/db/mysql"
)

func (user *User) AddNewUser(raw []byte) error {
	// parse raw
	var userData *sqlconnection.User
	err := json.Unmarshal(raw, &userData)
	if err != nil {
		return fmt.Errorf("error while unmarshaling user signup data: %v", err)
	}
	err = userData.SearchUserInUsers()
	if err != nil {
		return fmt.Errorf("error while adding user to db in AddNewUser auth.go: %v", err)
	}
	return nil
}

func (user *User) VerifyCredentials(raw []byte) (err error) {
	var userData *sqlconnection.User
	err = json.Unmarshal(raw, &userData)
	if err != nil {
		return fmt.Errorf("error while unmarshaling user login data: %v", err)
	}
	identifier := userData.Email
	if identifier != "" {
		user := keygen.User{
			Username: userData.Username,
		}
		err = user.ValidateToken(userData.Password)
		if err != nil {
			return fmt.Errorf("error while verifying user credentials: %v", err)
		}
		return
	}
	err = userData.VerifyCredentials()
	if err != nil {
		return fmt.Errorf("error while searching user in db in VerifyCredentials auth.go: %v", err)
	}
	return
}

func (user User) ResetPassword(raw []byte) error {
	return nil
}
