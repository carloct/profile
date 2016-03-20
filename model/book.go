package model

import (
	"log"

	"github.com/carloct/slv/shared/database"
)

type Book struct {
	Id      int64  `db:"id"`
	User_id int64  `db:"user_id"`
	Title   string `db:"title"`
}

func BookCreate(userId int64, title string) (*Book, error) {
	book := &Book{User_id: userId, Title: title}

	result, err := database.DB.Exec("INSERT INTO books (user_id, title) VALUES (?, ?)", userId, title)
	if err != nil {
		log.Println("Cannot create book")
		return book, err
	}

	book.Id, err = result.LastInsertId()
	return book, err
}
