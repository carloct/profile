package model

import (
	"log"
	"time"

	"github.com/carloct/profile/shared/database"
)

type Closet struct {
	Id        int       `db:"id"`
	UserId    uint32    `db:"user_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}

func ClosetCreate(userId uint32, name string) error {
	_, err := database.DB.Exec("INSERT INTO closets (user_id, name, created_at, updated_at) VALUES (?, ?, ?, ?)", userId, name, time.Now(), time.Now())
	if err != nil {
		log.Println("Cannot create closet")
	}
	return err
}

func Closets() ([]Closet, error) {
	closets := []Closet{}
	err := database.DB.Select(&closets, "SELECT * FROM closets")
	if err != nil {
		log.Println("Error Select closets")
	}
	return closets, err
}

func ClosetById(id int) (Closet, error) {
	closet := Closet{}
	err := database.DB.Get(&closet, "SELECT * FROM closets WHERE id = ?", id)
	return closet, err
}
