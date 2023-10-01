package models

import "time"

type Activity struct {
	Id        uint64    `db:"id"`
	Name      string    `db:"name"`
	Unit      string    `db:"unit"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
