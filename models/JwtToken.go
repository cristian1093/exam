package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tools"
)

// JWT token struct
type JWT struct {
	User User `json:"user,omitempty"`

	Token string `json:"-"`
}

// CreateAccessToken method genarte access token jwt
func (jwts JWT) CreateAccessToken() (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = jwts.User

	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err = token.SignedString([]byte(tools.GetEnv("Secure.JwtSecret")))

	return
}

// CreateRefreshToken method genarte refresh token jwt
func (jwts JWT) CreateRefreshToken() (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_email"] = jwts.User.UserEmail

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err = token.SignedString([]byte(tools.GetEnv("Secure.JwtSecret")))
	return
}

// CreateRestoreToken method genarte refresh token jwt
func (jwts JWT) CreateRestoreToken() (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["token"] = jwts.Token
	claims["id"] = jwts.User.UserId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err = token.SignedString([]byte(tools.GetEnv("Secure.JwtSecret")))
	return
}
