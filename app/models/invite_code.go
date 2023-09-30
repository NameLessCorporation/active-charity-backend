package models

import "time"

type InviteCode struct {
	ID             uint64
	Code           string
	OrganizationID uint64
	OwnerID        uint64
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
