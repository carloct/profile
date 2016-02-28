package model

import (
	"github.com/carloct/profile/shared/database"
	"log"
)

type Item struct {
	Id    int    `db:"id"`
	Url   string `db:"url"`
	Image string `db:"image"`
}

func ItemCreate(userId uint32, closetId int, url string, image string) error {
	_, err := database.DB.Exec("INSERT INTO items (user_id, closet_id, url, image) VALUES (?, ?, ?, ?)", userId, closetId, url, image)
	if err != nil {
		log.Println("Cannot create item")
	}
	return err
}
