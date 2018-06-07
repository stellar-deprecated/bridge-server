package entities

// AuthData represents auth data
type AuthData struct {
	exists   bool
	ID       string `db:"id"`
	Domain   string `db:"domain"`
	AuthData string `db:"auth_data"`
}

// GetID returns ID of the entity
func (e *AuthData) GetID() *int64 {
	return nil
}

// SetID sets ID of the entity
func (e *AuthData) SetID(id int64) {
	//
}

// IsNew returns true if the entity has not been persisted yet
func (e *AuthData) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *AuthData) SetExists() {
	e.exists = true
}
