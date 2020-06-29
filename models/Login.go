package models

import (
	"log"
	"time"
)

// "github.com/gin-gonic/gin"

const (
	TokenLifespan = time.Hour * 24 * 14
)

type LoginOutput struct {
	// Token     string    `json:"token"`
	// ExpiresAt time.Time `json:"expires_at"`
	ID_User  int    `json:"id"`
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

func Login(phone int, password string) []LoginOutput {

	var (
		login_user  LoginOutput
		login_users []LoginOutput
	)

	ShowPhone := ("select id, name,phone,email, password from users where phone = ? and password = ?;")

	rows, err := database.db.Query(ShowPhone, phone, password)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&login_user.ID_User, &login_user.Name, &login_user.Phone, &login_user.Email, &login_user.Password)
		login_users = append(login_users, login_user)
		if err != nil {
			log.Println(err)
		}

	}
	// fmt.Println("token")
	// fmt.Println(database.codec.EncodeToString(strconv.FormatInt(login_user.ID_User, 10)))
	defer rows.Close()
	return login_users

	// var out LoginOutput

	// ShowPhone := ("select id, name from users where password = ?;")
	// fmt.Println("select id, name from users where password = " + password + ";")

	// if err == sql.ErrNoRows {
	// 	log.Println(err)
	// 	panic(err)
	// }

	// if err != nil {

	// 	return out, fmt.Errorf("could not query select user: %v", err)
	// }

	// if err != nil {
	// 	return out, fmt.Errorf("could not create token:%v", err)
	// }

	// out.ExpiresAt = time.Now().Add(TokenLifespan)

	// return out, nil

}
