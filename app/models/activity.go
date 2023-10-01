package models

import "time"

type Activity struct {
	Id        uint64    `db:"id"`
	Name      string    `db:"name"`
	Unit      string    `db:"unit"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ActivityAnalytics struct {
	Id  uint64  `db:"id"`
	Avg float32 `db:"avg"`
	Min uint32  `db:"min"`
	Max uint32  `db:"max"`
}

type FavouriteActivity struct {
	Count int32  `db:"count"`
	Name  string `db:"name"`
}

type MostEarned struct {
	Sum  int32  `db:"sum"`
	Name string `db:"name"`
}
