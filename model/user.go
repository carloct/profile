package model

import (
	"time"

	"github.com/carloct/profile/shared/database"
)

type User struct {
	Id         uint32    `db:"id"`
	First_name string    `db:"first_name"`
	Last_name  string    `db:"last_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Status     bool      `db:"status"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Deleted    uint8     `db:"deleted"`
}

func UserCreate(firstName string, email string, password string) (User, error) {
	user := User{First_name: firstName, Email: email, Password: password}
	_, err := database.DB.Exec("INSERT INTO users (first_name, email, password, created_at, updated_at) VALUES (?,?,?,?,?)", firstName, email, password, time.Now(), time.Now())
	return user, err
}

func UserByEmail(email string) (User, error) {
	result := User{}

	err := database.DB.Get(&result, "SELECT * FROM users WHERE email = ?", email)

	return result, err
}
