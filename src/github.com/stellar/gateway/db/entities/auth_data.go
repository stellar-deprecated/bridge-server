package entities

// AuthData represents auth data
type AuthData struct {
	exists    bool
	ID        *int64 `db:"id"`
	RequestID string `db:"request_id"`
	Domain    string `db:"domain"`
	AuthData  string `db:"auth_data"`
}

// GetID returns ID of the entity
func (e *AuthData) GetID() *int64 {
	return e.ID
}

// SetID sets ID of the entity
func (e *AuthData) SetID(id int64) {
	e.ID = &id
}

// IsNew returns true if the entity has not been persisted yet
func (e *AuthData) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *AuthData) SetExists() {
	e.exists = true
}
