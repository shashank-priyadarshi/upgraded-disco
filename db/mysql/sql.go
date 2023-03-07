package sqlconnection

import (
	"database/sql"
	"errors"
	"fmt"
	"server/auth/keygen"
	"server/config"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Action   int    `json:"action"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// var close chan bool
func createConnection(dsn string) (*sql.DB, error) {
	// close <- false
	// create connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return &sql.DB{}, fmt.Errorf("error while creating connection in createConnection: %v", err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Second * 15)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	// go func() {
	// 	select {
	// 	case <-close:
	// 		db.Close()
	// 		log.Println("Closed connection after writing data successfully")
	// 	default:
	// 		time.Sleep(time.Second * 15)
	// 		db.Close()
	// 		log.Println("Connection timed out, failed to write data")
	// 	}
	// }()

	return db, nil
}

func (user *User) pushUserToUsers(db *sql.DB) (err error) {
	// Insert the password record
	passwordResult, err := db.Exec("INSERT INTO passwords (password, salt) VALUES (?, ?)", user.Password, "salt")
	if err != nil {
		// close <- true
		return fmt.Errorf("error while inserting password: %v", err)
	}

	// Retrieve the last inserted password ID
	passwordID, err := passwordResult.LastInsertId()
	if err != nil {
		// close <- true
		return fmt.Errorf("error while retrieving last inserted password ID: %v", err)
	}

	// Insert the user record
	_, err = db.Exec("INSERT INTO users (name, email, username, password_id) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.Username, passwordID)
	// close <- true
	if err != nil {
		return fmt.Errorf("error while inserting user data: %v", err)
	}
	return
}

func (user *User) searchUserInUsers(db *sql.DB) (err error) {
	err = user.pushUserToUsers(db)
	if err != nil {
		// close <- true
		return fmt.Errorf("error while adding user to db in SearchUserInUsers sql.go: %v", err)
	}
	// close <- true
	return nil
}

func verifyCredentials(db *sql.DB, passwordID int, currPassword string) (err error) {
	// Iterate over the rows
	savedPassword, err := fetchPassword(db, passwordID)
	if err != nil {
		// close <- true
		return
	}

	if !strings.EqualFold(savedPassword, currPassword) {
		return errors.New("wrong password")
	}

	// close <- true
	return
}

func resetPassword(db *sql.DB, passwordID int, currPassword string) (err error) {
	savedPassword, err := fetchPassword(db, passwordID)
	if err != nil {
		return
	}
	if !strings.EqualFold(savedPassword, currPassword) {
		return fmt.Errorf("wrong password")
	}
	return
}

func fetchPassword(db *sql.DB, passwordID int) (password string, err error) {
	row := db.QueryRow("SELECT password FROM passwords WHERE id = ?", passwordID)
	err = row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("password does not exist")
		} else {
			return "", errors.New("error while retrieving password")
		}
	}
	return
}

func (user *User) EntryPoint() (token string, err error) {
	// Create connection
	db, err := createConnection(config.FetchConfig().SQLURI)
	if err != nil {
		return "", fmt.Errorf("error while creating connection: %v", err)
	}

	// Search the user record
	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM users WHERE username = ? OR EMAIL = ?", user.Username, user.Email)

	if err != nil {
		// close <- true
		return "", fmt.Errorf("error while checking if user exists: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows
	if rows.Next() {
		if user.Action == 0 {
			return "", errors.New("user already exists")
		}
		// close <- true
		var id int
		var name string
		var email string
		var username string
		var passwordID int
		var createdAt string

		err := rows.Scan(&id, &name, &email, &username, &createdAt, &passwordID)
		if err != nil {
			return "", errors.New("error while retrieving user data")
		}

		if user.Action == 1 {
			tokenUser := keygen.User{
				Username: user.Username,
			}
			token, err = tokenUser.GenerateToken()
			if err != nil {
				return "", err
			}
			return token, verifyCredentials(db, passwordID, user.Password)
		} else {
			return "", resetPassword(db, passwordID, user.Password)
		}
	}

	if user.Action == 0 {
		tokenUser := keygen.User{
			Username: user.Username,
		}
		token, err = tokenUser.GenerateToken()
		if err != nil {
			return "", err
		}
		err = user.searchUserInUsers(db)
	} else {
		err = errors.New("user not found")
	}

	return
}
