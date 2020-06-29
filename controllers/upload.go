package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

func UploadLayout(c *gin.Context) {

	var (
		file_name                string
		messages                 []string
		construct_sql            string
		construct_sql_afiliacion string
		columns                  int
		rowsFile                 int
		id_affiliation           string
	)

	c.Request.ParseMultipartForm(2000)

	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	file_name = header.Filename

	result := strings.Split(file_name, ".")

	if strings.Contains(result[0], "INBURSA") == false {

		messages = append(messages, "the file does not contain the correct name")
		BadRequest(c, messages)
		return
	} else if result[1] != "xlsx" {
		messages = append(messages, "the file is not the correct type")
		BadRequest(c, messages)
		return

	} else {

		f, err := os.OpenFile("./files/"+file_name, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Fatal(err)
			return
		}

		defer f.Close()

		io.Copy(f, file)

		xslx, err := excelize.OpenFile("./files/" + file_name)
		if err != nil {
			fmt.Println(err)
			return
		}
		sheets := xslx.GetSheetMap()
		sheet := (sheets[1])
		rows, _ := xslx.GetRows(sheet)

		//count of columns

		columns = 1

		for _, rowColumns := range rows {
			for columns, _ = range rowColumns {
				columns++
			}
		}

		//row count
		rowsFile = 1
		for rowsFile, _ = range rows {
			rowsFile++
		}

		fmt.Println("filas", rowsFile)
		fmt.Println("columnas", columns)

		i := 2
		for i <= rowsFile {
			cell, _ := xslx.GetCellValue("Hoja1", ("A" + strconv.Itoa(i)))
			cell2, _ := xslx.GetCellValue("Hoja1", ("B" + strconv.Itoa(i)))
			cell3, _ := xslx.GetCellValue("Hoja1", ("C" + strconv.Itoa(i)))
			cell4, _ := xslx.GetCellValue("Hoja1", ("D" + strconv.Itoa(i)))
			cell5, _ := xslx.GetCellValue("Hoja1", ("E" + strconv.Itoa(i)))
			cell6, _ := xslx.GetCellValue("Hoja1", ("F" + strconv.Itoa(i)))
			cell7, _ := xslx.GetCellValue("Hoja1", ("G" + strconv.Itoa(i)))
			cell8, _ := xslx.GetCellValue("Hoja1", ("H" + strconv.Itoa(i)))
			cell9, _ := xslx.GetCellValue("Hoja1", ("I" + strconv.Itoa(i)))
			cell10, _ := xslx.GetCellValue("Hoja1", ("J" + strconv.Itoa(i)))
			cell11, _ := xslx.GetCellValue("Hoja1", ("K" + strconv.Itoa(i)))
			cell12, _ := xslx.GetCellValue("Hoja1", ("L" + strconv.Itoa(i)))
			cell13, _ := xslx.GetCellValue("Hoja1", ("M" + strconv.Itoa(i)))
			cell14, _ := xslx.GetCellValue("Hoja1", ("N" + strconv.Itoa(i)))
			cell15, _ := xslx.GetCellValue("Hoja1", ("O" + strconv.Itoa(i)))

			construct_sql_afiliacion += "'" + cell2 + "',"

			construct_sql += "(" + "'" + cell + "'," + "'" + cell2 + "'," + "'" + cell3 + "'," + "'" + cell4 + "'," + "'" + cell5 + "'," + "'" + cell6 + "'," + "'" + cell7 + "'," + "'" + cell8 + "'," + "'" + cell9 + "'," + "'" + cell10 + "'," + "'" + cell11 + "'," + "'" + cell12 + "'," + "'" + cell13 + "'," + "'" + cell14 + "'," + "'" + cell15 + "'," + "'" + file_name + "'),"
			i = i + 1
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		construct_sql = construct_sql[0 : len(construct_sql)-1]

		construct_sql_afiliacion = construct_sql_afiliacion[0 : len(construct_sql_afiliacion)-1]

		data_id_affiliation := models.ShowAffiliationId(construct_sql_afiliacion)

		for _, row := range data_id_affiliation {
			fmt.Println(row.Id_Affiliation)
			id_affiliation += "'" + row.Id_Affiliation + "',"

			validate_affiliation := strings.Replace(construct_sql, row.Id_Affiliation, "true"+row.Id_Affiliation, 1000)

			construct_sql = validate_affiliation

		}
		id_affiliation = id_affiliation[0 : len(id_affiliation)-1]

		bool := models.UploadLayout(construct_sql)

		if bool == true {
			c.JSON(200, gin.H{"status": "successful loading"})
		} else {
			c.JSON(200, gin.H{"status": "error in loading"})
		}

		err = os.Rename("./files/"+file_name, "./files/upload"+file_name)
		if err != nil {
			fmt.Println(err)
		}

		defer models.Truncate()
	}

}
