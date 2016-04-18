package entities

import (
	"time"
)

// AllowedUser represents allowed user
type AllowedUser struct {
	exists      bool
	ID          *int64    `db:"id"`
	FiName      string    `db:"fi_name"`
	FiDomain    string    `db:"fi_domain"`
	FiPublicKey string    `db:"fi_public_key"`
	UserID      string    `db:"user_id"`
	AllowedAt   time.Time `db:"allowed_at"`
}

// GetID returns ID of the entity
func (e *AllowedUser) GetID() *int64 {
	return e.ID
}

// SetID sets ID of the entity
func (e *AllowedUser) SetID(id int64) {
	e.ID = &id
}

// IsNew returns true if the entity has not been persisted yet
func (e *AllowedUser) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *AllowedUser) SetExists() {
	e.exists = true
}
