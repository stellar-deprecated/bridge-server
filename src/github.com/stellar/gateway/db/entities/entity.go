package entities

type Entity interface {
	GetId() *int64 // Returns nil if object hasn't been persisted yet.
	SetId(int64)   // Used by driver. Sets `Id` field to the id of the row in DB.
	IsNew() bool   // Returns true if object hasn't been persisted in DB yet.
	SetExists()    // Used by driver. Sets internal `exists` flag of Entity to true.
}
