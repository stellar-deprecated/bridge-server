package entities

import (
	"time"
)

type AllowedFi struct {
	exists    bool
	Id        *int64    `db:"id"`
	Name      string    `db:"name"`
	Domain    string    `db:"domain"`
	PublicKey string    `db:"public_key"`
	AllowedAt time.Time `db:"allowed_at"`
}

func (e *AllowedFi) GetId() *int64 {
	return e.Id
}

func (e *AllowedFi) SetId(id int64) {
	e.Id = &id
}

func (e *AllowedFi) IsNew() bool {
	return !e.exists
}

func (e *AllowedFi) SetExists() {
	e.exists = true
}
