package controllers

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

type PagoMembresia struct {
	Monto            int    `json:"monto"`
	Tarjeta          string `json:"tarjeta"`
	FechaVencimiento string `json:"fecha_vencimiento"`
	CVV              int    `json:"cvv"`
}

func (request *PagoMembresia) validatioMontoRequest() (messages []string, err error) {

	if request.Monto < 100000 {
		messages = append(messages, "El monto es superior s los 100 mil pesos mexicanos")
		err = fmt.Errorf("El monto es superior s los 100 mil pesos mexicanos")
	}
	if isPhone, _ := regexp.MatchString(`[0-9]*[1-9]`, request.Tarjeta); !isPhone {
		messages = append(messages, "El Primer Campo No puede ser 0")
		err = fmt.Errorf("El Primer Campo No puede ser 0")
	}

	return
}

func Membresia(c *gin.Context) {

	var (
		request  PagoMembresia
		messages []string
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	if messages, err := request.validatioMontoRequest(); err != nil {
		BadRequest(c, messages)
		return

	}

	fmt.Println(messages)

	c.JSON(200, gin.H{"token": "mempresia"})
}
