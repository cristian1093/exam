package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/models"

	jwt "github.com/dgrijalva/jwt-go"
)

type LoginUser2 struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

var mySigningKey = []byte("mysupersecret")

func GenerateJWT(name string, email string, id string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err

	}

	return tokenString, nil

}

func (request *LoginUser2) validationLoginRequest2() (messages []string, err error) {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(request.Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range request.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true

		}
	}

	if (hasMinLen != true) || (hasUpper != true) || (hasLower != true) || (hasNumber != true) {
		messages = append(messages, "Password must be at least one lowercase letter, one uppercase letter, and one number and contain at least 8 characters")
		err = fmt.Errorf("Password must be at least one lowercase letter, one uppercase letter, and one number and contain at least 8 characters")
	}

	if isPhone, _ := regexp.MatchString(`([0-9]*52).{10,10}$`, strconv.Itoa(request.Phone)); !isPhone {
		messages = append(messages, "The phone does not have the correct format (lada México)")
		err = fmt.Errorf("The phone does not have the correct format (lada México)")
	}

	return
}

func Login2(c *gin.Context) {

	var (
		request  LoginUser2
		messages []string
		user     models.User
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	if messages, err := request.validationLoginRequest2(); err != nil {
		BadRequest(c, messages)
		return
	}

	user.UserPhone = request.Phone
	if err := user.UserFindByEmail(); err != nil {
		log.Println("Error not find user by phone in database", err)
		if err.Error() == "sql: no rows in result set" {
			messages = append(messages, "The user phone doesn´t exists")
			BadRequest(c, messages)
			return
		}
		ServerError(c)
		return
	}

	respon := models.Login(request.Phone, request.Password)

	var password string
	var phone int

	for _, row := range respon {

		password += row.Password
		phone += row.Phone

	}

	if password != request.Password {

		c.JSON(401, gin.H{"messages": "The user phone or password not valid"})

	} else {

		tokenString, err := GenerateJWT(user.UserName, user.UserEmail, user.UserId)

		if err != nil {
			fmt.Println("Errror generating token string")
		}

		fmt.Println(tokenString)

		c.JSON(200, gin.H{"token": tokenString})
	}
}
