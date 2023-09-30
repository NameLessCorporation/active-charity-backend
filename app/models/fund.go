package models

import "time"

type Fund struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	OwnerID     uint64    `db:"owner_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
