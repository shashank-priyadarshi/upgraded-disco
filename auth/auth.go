package auth

import (
	"encoding/json"
	"fmt"
	"server/auth/keygen"
	sqlconnection "server/db/mysql"
)

func (user *User) ParseCredentials(raw []byte) (token string, err error) {
	// parse raw
	// action: 0 = signup, 1 = loginusingpassword, 2 = loginusingtoken, 3 = resetpassword
	var userData *sqlconnection.User
	err = json.Unmarshal(raw, &userData)
	if err != nil {
		return "", fmt.Errorf("error while unmarshaling user signup data: %v", err)
	}

	if userData.Action == 2 {
		user := keygen.User{
			Username: userData.Username,
		}
		return "", user.ValidateToken(userData.Password)
	} else {
		token, err = userData.EntryPoint()
	}

	return
}
