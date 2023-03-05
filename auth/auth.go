package auth

import (
	"encoding/json"
	"fmt"
	sqlconnection "server/db/mysql"
)

func AddNewUser(raw []byte) error {
	// parse raw
	var userData *SignUp
	err := json.Unmarshal(raw, &userData)
	if err != nil {
		return fmt.Errorf("error while unmarshaling user signup data: %v", err)
	}
	err = sqlconnection.SearchUserInUsers(struct {
		Name     string
		Email    string
		Password string
		Phone    string
	}(*userData))
	if err != nil {
		return fmt.Errorf("error while adding user to db in AddNewUser auth.go: %v", err)
	}
	return nil
}

func VerifyCredentials(raw []byte) error {
	var userData *Login
	err := json.Unmarshal(raw, &userData)
	if err != nil {
		return fmt.Errorf("error while unmarshaling user login data: %v", err)
	}
	identifier := ""
	if userData.Email == "" {
		identifier = userData.Phone
	} else {
		identifier = userData.Email
	}
	err = sqlconnection.VerifyCredentials(struct {
		Identifier string
		Password   string
	}{
		Identifier: identifier,
		Password:   userData.Password,
	})
	if err != nil {
		return fmt.Errorf("error while searching user in db in VerifyCredentials auth.go: %v", err)
	}
	return nil
}

func ResetPassword(raw []byte) error {
	return nil
}
