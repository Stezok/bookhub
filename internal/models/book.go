package models

type Book struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	Desc  string `db:"description"`
}
