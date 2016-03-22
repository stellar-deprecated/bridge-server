package entities

import (
	"time"
)

type AllowedUser struct {
	exists      bool
	Id          *int64    `db:"id"`
	FiName      string    `db:"fi_name"`
	FiDomain    string    `db:"fi_domain"`
	FiPublicKey string    `db:"fi_public_key"`
	UserId      string    `db:"user_id"`
	AllowedAt   time.Time `db:"allowed_at"`
}

func (e *AllowedUser) GetId() *int64 {
	return e.Id
}

func (e *AllowedUser) SetId(id int64) {
	e.Id = &id
}

func (e *AllowedUser) IsNew() bool {
	return !e.exists
}

func (e *AllowedUser) SetExists() {
	e.exists = true
}
