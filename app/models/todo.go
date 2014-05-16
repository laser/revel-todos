package models

type Todo struct {
	Title     string `db:"title"`
	Completed bool   `db:"completed"`
	Id        int    `db:"id"`
}
