package models

import (
	// "github.com/gin-gonic/gin"
	"fmt"
	"log"
)

type Id_Affiliation struct {
	Id_Affiliation string `json:"id_affiliation"`
}

func UploadLayout(data string) bool {

	exito := false

	fmt.Println("insert into chargebacks_temporal(date,afilliation,business_name,date_tx,import,account,ref,acl_type,clarification_type,autorization,clarification_status,commitment_date, request_type,description,chargeback_code,filename) values " + data + ";")

	createDataSQL := ("insert into chargebacks_temporal(date,afilliation,business_name,date_tx,import,account,ref,acl_type,clarification_type,autorization,clarification_status,commitment_date, request_type,description,chargeback_code,filename) values " + data + ";")

	// call procedure(mysql, layout)

	_, err := database.db.Query(createDataSQL)

	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		exito = true
	}

	return exito

}

func ShowAffiliationId(data string) []Id_Affiliation {

	var (
		id_affiliation  Id_Affiliation
		id_affiliations []Id_Affiliation
	)

	getIdAffiliationSQL := ("select aggregator_id from bw_db.affiliations_bw where aggregator_id IN ( " + data + ")")

	rows, err := database.db.Query(getIdAffiliationSQL)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&id_affiliation.Id_Affiliation)
		id_affiliations = append(id_affiliations, id_affiliation)
		if err != nil {
			log.Println(err)
		}

	}
	defer rows.Close()
	return id_affiliations

}

func Truncate() {

	truncateDataSQL := ("TRUNCATE TABLE public.chargebacks_temporal CONTINUE IDENTITY RESTRICT;")
	db := getConnection()
	_, err := db.Query(truncateDataSQL)

	if err != nil {
		log.Println(err)
		panic(err)
	}

}
