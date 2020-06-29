package models

import(
  "database/sql"
  _"github.com/lib/pq"
  "log"
)

func getConnection() *sql.DB{
  dsn:= "postgres://cbonilla:Banwire1234@127.0.0.1:5432/contracargos?sslmode=disable"
  db, err:=sql.Open("postgres", dsn)
  if err != nil{
      log.Fatal(err)
  }
  return db
}
