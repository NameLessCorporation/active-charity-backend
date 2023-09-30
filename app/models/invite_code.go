package models

type InviteCode struct {
	ID             uint64
	Code           string
	OrganizationID uint64
	OwnerID        uint64
}
