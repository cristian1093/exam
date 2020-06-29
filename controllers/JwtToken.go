package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"github.com/tools"
)

// Refresh struct refresh_token request
type Refresh struct {
	RefreshToken string `json:"refresh_token"`
}

// validationTokens valid the request
func (request *Refresh) validationTokens() (messages []string, err error) {

	if request.RefreshToken == "" {
		messages = append(messages, "The refresh_token is necesary")
		err = fmt.Errorf("refresh_token malformed")
	}

	return
}

// RefreshToken handle
func RefreshToken(c *gin.Context) {
	var (
		request  Refresh
		messages []string
		jwts     models.JWT

		response struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}

	if messages, err := request.validationTokens(); err != nil {
		BadRequest(c, messages)
		return
	}

	if token, err := IsValidToken(request.RefreshToken); err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			jwts.User.UserEmail = fmt.Sprintf("%v", claims["user_email"])

			if err := jwts.User.UserFindByEmail(); err != nil {
				//TODO: apply logs system
				if err.Error() == "sql: no rows in result set" {
					messages = append(messages, "The user email doesn´t exist")
					BadRequest(c, messages)
					return
				}

				ServerError(c)
				return
			}

		}
	} else {
		//TODO: apply logs system
		messages = append(messages, "The token no valid")
		BadRequest(c, messages)
		return
	}
	response.AccessToken, _ = jwts.CreateAccessToken()
	response.RefreshToken = request.RefreshToken

	c.JSON(http.StatusOK, response)
}

//IsValidToken valid if the token is ok
func IsValidToken(inToken string) (token *jwt.Token, err error) {

	token, err = jwt.Parse(inToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//TODO: apply logs system
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tools.GetEnv("Secure.JwtSecret")), nil
	})

	return
}
