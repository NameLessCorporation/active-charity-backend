package models

import "time"

type Graph struct {
	Timestamp time.Time `db:"created_at"`
	Value     uint32    `db:"value"`
}
