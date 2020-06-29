package models

import (
	"log"
)

// "github.com/gin-gonic/gin"
type User struct {
	UserName     string `json:"username"`
	UserPhone    int    `json:"user_phone"`
	UserEmail    string `json:"user_email"`
	UserId       string `json:"user_id"`
	UserPassword string `json:"user_password"`
}

func InsertUser(name string, phone int, email string, password string) {

	createUserDataSQL := ("insert into Yofio.users (name, phone, email, password) values (?,?,?,?)")

	_, err := database.db.Query(createUserDataSQL, name, phone, email, password)

	if err != nil {
		log.Println(err)
		panic(err)

	}

}

func ShowPhoneUser(phone int) []User {

	var (
		phone_user  User
		phone_users []User
	)

	ShowPhone := ("select phone from users where phone IN (?) limit 1")

	rows, err := database.db.Query(ShowPhone, phone)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&phone_user.UserPhone)
		phone_users = append(phone_users, phone_user)
		if err != nil {
			log.Println(err)
		}

	}
	defer rows.Close()
	return phone_users
}

func (user *User) UserFindByEmail() (err error) {
	querySelect := "select id, name, phone, email, password from users where phone = ?"

	row := database.db.QueryRow(querySelect, user.UserPhone)
	err = row.Scan(&user.UserId, &user.UserName, &user.UserPhone, &user.UserEmail, &user.UserPassword)
	return
}
