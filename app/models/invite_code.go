package models

import "time"

type InviteCode struct {
	ID             uint64
	Code           string
	OrganizationID uint64    `db:"organization_id"`
	OwnerID        uint64    `db:"owner_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
