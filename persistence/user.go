package persistence

import (
	"database/sql"
	"log"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func SaveUser (db *sql.DB,u *User) {
	sqlStatement := `INSERT INTO users (age, email, first_name, last_name) VALUES ($1, $2, $3, $4)`
	log.Println("Creating User")
	_, err := db.Exec(sqlStatement,  u.Email, u.FirstName, u.LastName)
	if err != nil {
		panic(err)
	}
	log.Println("Saved user!")
}