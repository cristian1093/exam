package controllers

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/models"
)

type LoginUser struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

func (request *LoginUser) validationLoginRequest() (messages []string, err error) {
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
		messages = append(messages, "ERROR PHONE")
		err = fmt.Errorf("ERROR PHONE")
	}

	return
}

func Login(c *gin.Context) {

	var (
		request  LoginUser
		messages []string
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	if messages, err := request.validationLoginRequest(); err != nil {
		BadRequest(c, messages)
		return
	}
	fmt.Println(messages)

	response := models.Login(request.Phone, request.Password)

	if response != nil {
		c.JSON(200, gin.H{"status": response})
	} else {
		c.JSON(401, gin.H{"status": "the login information is incorrect"})
	}
}
