package sqlconnection

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

func pushUserToUsers(user struct {
	Name     string
	Email    string
	Password string
	Phone    string
}) error {
	db, err := createConnection("")
	if err != nil {
		return fmt.Errorf("error while creating connection for pushUserToUsers: %v", err)
	}

	// Insert the password record
	passwordResult, err := db.Exec("INSERT INTO passwords (password, salt) VALUES (?, ?)", user.Password, "user_salt")
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
	_, err = db.Exec("INSERT INTO users (name, email, phone, password_id) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.Phone, passwordID)
	// close <- true
	if err != nil {
		return fmt.Errorf("error while inserting user data: %v", err)
	}
	return nil
}

func SearchUserInUsers(user struct {
	Name     string
	Email    string
	Password string
	Phone    string
}) error {
	db, err := createConnection("")
	if err != nil {
		return fmt.Errorf("error while creating connection for SearchUserInUsers: %v", err)
	}

	// Search the user record
	rows, err := db.Query("SELECT * FROM users WHERE email = ? OR phone = ?", user.Email, user.Phone)
	if err != nil {
		// close <- true
		return fmt.Errorf("error while checking if user exists: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows
	if rows.Next() {
		// close <- true
		return errors.New("user already exists")
	} else {
		err := pushUserToUsers(user)
		if err != nil {
			// close <- true
			return fmt.Errorf("error while adding user to db in SearchUserInUsers sql.go: %v", err)
		}
	}
	// close <- true
	return nil
}

func VerifyCredentials(user struct {
	Identifier string
	Password   string
}) error {
	db, err := createConnection("")
	if err != nil {
		return fmt.Errorf("error while creating connection for SearchUserInUsers: %v", err)
	}

	// Search the user record
	var rows *sql.Rows
	if strings.Contains(user.Identifier, "@") {
		rows, err = db.Query("SELECT * FROM users WHERE email = ?", user.Identifier)
	} else {
		rows, err = db.Query("SELECT * FROM users WHERE phone = ?", user.Identifier)
	}

	if err != nil {
		// close <- true
		return fmt.Errorf("error while checking if user exists: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows
	if rows.Next() {
		// close <- true
		var password string
		var id int
		var name string
		var email string
		var phone string
		var passwordID int
		var createdAt string

		err = rows.Scan(&id, &name, &email, &phone, &createdAt, &passwordID)
		if err != nil {
			return errors.New("error while retrieving user data")
		}

		row := db.QueryRow("SELECT password FROM passwords WHERE id = ?", passwordID)
		err = row.Scan(&password)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("password does not exist")
			} else {
				return errors.New("error while retrieving password")
			}
		}

		if strings.EqualFold(password, user.Password) {
			return nil
		}
	}
	// close <- true
	return errors.New("user does not exist")
}

func ResetPassword(user struct {
	Identifier string
}) error {
	// Find user based on email/phone
	// Read password id cell value from response
	// Replace password in passwords table using password id
	db, err := createConnection("")
	if err != nil {
		return fmt.Errorf("error while creating connection for SearchUserInUsers: %v", err)
	}

	// Search the user record
	var rows *sql.Rows
	if strings.Contains(user.Identifier, "@") {
		rows, err = db.Query("SELECT * FROM users WHERE email = ?", user.Identifier)
	} else {
		rows, err = db.Query("SELECT * FROM users WHERE phone = ?", user.Identifier)
	}

	if err != nil {
		// close <- true
		return fmt.Errorf("error while checking if user exists: %v", err)
	}
	defer rows.Close()
	// Iterate over the rows
	if rows.Next() {
		// close <- true
		var password string
		var id int
		var name string
		var email string
		var phone string
		var passwordID int
		var createdAt string

		err = rows.Scan(&id, &name, &email, &phone, &createdAt, &passwordID)
		if err != nil {
			return errors.New("error while retrieving user data")
		}

		row := db.QueryRow("SELECT password FROM passwords WHERE id = ?", passwordID)
		err = row.Scan(&password)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("password does not exist")
			} else {
				return errors.New("error while retrieving password")
			}
		}
	}
	return nil
}
