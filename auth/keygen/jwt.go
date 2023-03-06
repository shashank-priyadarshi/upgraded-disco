package keygen

import (
	"errors"
	"fmt"
	"server/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string
}

type Token interface {
	GenerateToken() (string, error)
	ValidateToken() error
}

var secretkey = config.FetchConfig().SECRETKEY

func (user User) GenerateToken() (string, error) {

	expirationTime := time.Now().AddDate(0, 0, 7)
	claims := jwt.MapClaims{
		"iss": user.Username,
		"exp": expirationTime.Unix(),
		"sub": user.Username,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretkey)
	if err != nil {
		return "", errors.New("error while generating token")
	}
	return signedToken, err
}

func (user User) ValidateToken(tokenString string) (err error) {
	// parse token, verify signing method is RSA
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid token")
		}
		return secretkey, nil
	})
	if err != nil {
		return err
	}

	// Verify that the "sub" claim matches the expected username
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok || sub != user.Username {
			fmt.Println("Invalid subject")
			return
		}

		// Verify that the "exp" claim matches the expected expiration time
		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			fmt.Println("Invalid expiration time")
			return
		}

		fmt.Println("Valid token")
	} else {
		fmt.Println("Invalid token")
	}
	return nil
}
