package model

type Page struct {
	Id      int    `db:"id"`
	Book_id int    `db:"book_id"`
	Counter string `db:"counter"`
}
