package session

import "net/http"

// ManagerInterface session manager interface
type ManagerInterface interface {
	// Add value to session data, if value is not string, will marshal it into JSON encoding and save it into session data.
	Add(req *http.Request, key string, value interface{}) error
	// Get value from session data
	Get(req *http.Request, key string) string
	// Pop value from session data
	Pop(req *http.Request, key string) string

	// Flash add flash message to session data
	Flash(req *http.Request, message Message) error
	// Flashes returns a slice of flash messages from session data
	Flashes(req *http.Request) []Message

	// Load get value from session data and unmarshal it into result
	Load(req *http.Request, key string, result interface{}) error
	// PopLoad pop value from session data and unmarshal it into result
	PopLoad(req *http.Request, key string, result interface{}) error

	// Middleware returns a new session manager middleware instance.
	Middleware(http.Handler) http.Handler
}

// Message message struct
type Message struct {
	Message string
	Type    string
}
