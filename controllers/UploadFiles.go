package controllers

import (
	//"github.com/360EntSecGroup-Skylar/excelize/v2"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

func UploadFiles(c *gin.Context) {

	xlsx := excelize.NewFile()

	xlsx.SetSheetName("Sheet1", "Prueba")
	xlsx.SetCellValue("Sheet1", "A1", "Id")
	xlsx.SetCellValue("Sheet1", "B1", "Date")
	xlsx.SetCellValue("Sheet1", "C1", "Affiliation")
	xlsx.SetCellValue("Sheet1", "D1", "Business Name")
	xlsx.SetCellValue("Sheet1", "E1", "Date TX")
	xlsx.SetCellValue("Sheet1", "F1", "Import")
	xlsx.SetCellValue("Sheet1", "G1", "Account")
	xlsx.SetCellValue("Sheet1", "H1", "Reference")
	xlsx.SetCellValue("Sheet1", "I1", "Act Type")
	xlsx.SetCellValue("Sheet1", "J1", "Clasification Status")
	xlsx.SetCellValue("Sheet1", "K1", "Commitment Date")
	xlsx.SetCellValue("Sheet1", "L1", "	Request Type")
	xlsx.SetCellValue("Sheet1", "M1", "Description")
	xlsx.SetCellValue("Sheet1", "N1", "Chargeback Code")
	xlsx.SetCellValue("Sheet1", "O1", "File Name")
	xlsx.SetCellValue("Sheet1", "P1", "Key")
	xlsx.SetCellValue("Sheet1", "Q1", "Created At")
	xlsx.SetCellValue("Sheet1", "R1", "Updates At")
	xlsx.SetCellValue("Sheet1", "S1", "File Name")

	datos := models.ShowData()

	celda := 2
	for _, row := range datos {
		celda1 := strconv.Itoa((celda))

		xlsx.SetCellValue("Sheet1", "A"+celda1, row.Id)
		xlsx.SetCellValue("Sheet1", "B"+celda1, row.Date)
		xlsx.SetCellValue("Sheet1", "C"+celda1, row.Affiliation)
		xlsx.SetCellValue("Sheet1", "D"+celda1, row.BusinessName)
		// xlsx.SetCellValue("Sheet1", "E"+celda1, "Date TX")
		// xlsx.SetCellValue("Sheet1", "F"+celda1, "Import")
		// xlsx.SetCellValue("Sheet1", "G"+celda1, "Account")
		// xlsx.SetCellValue("Sheet1", "H"+celda1, "Reference")
		// xlsx.SetCellValue("Sheet1", "I"+celda1, "Act Type")
		// xlsx.SetCellValue("Sheet1", "J"+celda1, "Clasification Status")
		// xlsx.SetCellValue("Sheet1", "K"+celda1, "Commitment Date")
		// xlsx.SetCellValue("Sheet1", "L"+celda1, "Request Type")
		// xlsx.SetCellValue("Sheet1", "M"+celda1, "Description")
		// xlsx.SetCellValue("Sheet1", "N"+celda1, "Chargeback Code")
		// xlsx.SetCellValue("Sheet1", "O"+celda1, "File Name")
		// xlsx.SetCellValue("Sheet1", "P"+celda1, "Key")
		// xlsx.SetCellValue("Sheet1", "Q"+celda1, "Created At")
		// xlsx.SetCellValue("Sheet1", "R"+celda1, "Updates At")
		// xlsx.SetCellValue("Sheet1", "S"+celda1, "File Name")

		celda++
	}
	err := xlsx.SaveAs("./files/hola.xlsx")

	if err != nil {
		fmt.Println(err)
	}

	archivo := "./files/hola.xlsx"
	c.JSON(200, gin.H{
		"success": true,
		"liga":    archivo,
	})
}
