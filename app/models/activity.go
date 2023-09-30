package models

type Activity struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}
