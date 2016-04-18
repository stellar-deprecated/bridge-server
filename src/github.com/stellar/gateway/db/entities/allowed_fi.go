package entities

import (
	"time"
)

// AllowedFi represents allowed FI
type AllowedFi struct {
	exists    bool
	ID        *int64    `db:"id"`
	Name      string    `db:"name"`
	Domain    string    `db:"domain"`
	PublicKey string    `db:"public_key"`
	AllowedAt time.Time `db:"allowed_at"`
}

// GetID returns ID of the entity
func (e *AllowedFi) GetID() *int64 {
	return e.ID
}

// SetID sets ID of the entity
func (e *AllowedFi) SetID(id int64) {
	e.ID = &id
}

// IsNew returns true if the entity has not been persisted yet
func (e *AllowedFi) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *AllowedFi) SetExists() {
	e.exists = true
}
