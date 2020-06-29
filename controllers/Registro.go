package controllers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ListenRegistrationRequest struct {
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *ListenRegistrationRequest) validationListenRegistrationRequest() (messages []string, err error) {
	if len(request.Name) < 1 {
		messages = append(messages, "The neme is not valid")
		err = fmt.Errorf("The neme is not valid")
	}

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

	words := strings.Fields(request.Name)

	division_of_words := strings.Split(request.Name, " ")

	var word_arrangement []int
	for i := 0; i < len(words); i++ {
		runes := []rune(division_of_words[i])
		for _, result := range string(runes[0:1]) {

			switch {
			case unicode.IsUpper(result):
				word_arrangement = append(word_arrangement, 1)
			}

		}
	}
	sum_of_words := 0
	for _, num := range word_arrangement {
		sum_of_words += num
	}

	if sum_of_words != len(words) {
		messages = append(messages, "Each word must have capital letters in their initials")
		err = fmt.Errorf("Each word must have capital letters in their initials")
	} else if len(words) <= 1 {
		messages = append(messages, "The name must have more than one word, each word must have capital letters in their initials")
		err = fmt.Errorf("The name must have more than one word, each word must have capital letters in their initials")
	}
	return
}

func Registration(c *gin.Context) {
	var (
		request  ListenRegistrationRequest
		messages []string
	)
	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	if messages, err := request.validationListenRegistrationRequest(); err != nil {
		BadRequest(c, messages)
		return
	}
	fmt.Println(messages)

	respuesta := models.ShowPhoneUser(request.Phone)
	var test int
	for _, row := range respuesta {
		fmt.Println(row.UserPhone)

		test += row.UserPhone

	}

	if request.Phone == test {

		c.JSON(412, gin.H{"status": "The phone already belongs to another user"})
	} else {

		defer models.InsertUser(request.Name, request.Phone, request.Email, request.Password)

		c.JSON(200, gin.H{"status": "Successful registration"})
	}

}
