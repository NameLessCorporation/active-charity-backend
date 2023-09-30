package models

type Fund struct {
	ID          uint64 `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	OwnerID     uint64 `db:"owner_id"`
}
