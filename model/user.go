package model

import (
	"fmt"
)

type User struct {
	Id         uint32    `db:"id"`
	First_name string    `db:"first_name"`
	Last_name  string    `db:"last_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Status_id  uint8     `db:"status_id"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Deleted    uint8     `db:"deleted"`
}

func UserCreate(name string, email string, password string) {
	_, err := database.DB.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES (?,?,?,?)", first_name, last_name, email, password)
	return err
}
