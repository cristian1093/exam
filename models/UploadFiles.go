package models

import "log"

type Data struct {
	Id           int64  `json:"chargeback_id"`
	Date         string `json:"date"`
	Affiliation  string `json:"afilliation"`
	BusinessName string `json:"business_name"`
}

func ShowData() []Data {

	var (
		data  Data
		datas []Data
	)
	db := getConnection()

	getIdAffiliationSQL := ("select chargeback_id, date, afilliation, business_name from public.chargebacks ")

	rows, err := db.Query(getIdAffiliationSQL)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&data.Id,
			&data.Date,
			&data.Affiliation,
			&data.BusinessName,
		)
		datas = append(datas, data)
		if err != nil {
			log.Println(err)
		}

	}
	defer rows.Close()
	return datas

}
