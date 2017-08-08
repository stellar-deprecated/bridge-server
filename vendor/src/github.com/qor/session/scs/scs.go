package scs

import (
	"encoding/json"
	"net/http"

	scssession "github.com/alexedwards/scs/session"
	"github.com/qor/session"
)

// New initialize session manager for SCS
func New(engine scssession.Engine, opts ...scssession.Option) *SCS {
	return &SCS{Engine: engine, Options: opts}
}

// SCS session manager struct for SCS
type SCS struct {
	scssession.Engine
	Options []scssession.Option
}

// Add value to session data, if value is not string, will marshal it into JSON encoding and save it into session data.
func (scs SCS) Add(req *http.Request, key string, value interface{}) error {
	if str, ok := value.(string); ok {
		return scssession.PutString(req, key, str)
	}
	result, _ := json.Marshal(value)
	return scssession.PutString(req, key, string(result))
}

// Pop value from session data
func (scs SCS) Pop(req *http.Request, key string) string {
	result, _ := scssession.PopString(req, key)
	return result
}

// Get value from session data
func (scs SCS) Get(req *http.Request, key string) string {
	result, _ := scssession.GetString(req, key)
	return result
}

// Flash add flash message to session data
func (scs SCS) Flash(req *http.Request, message session.Message) error {
	var messages []session.Message
	if err := scs.Load(req, "_flashes", &messages); err != nil {
		return err
	}
	messages = append(messages, message)
	return scs.Add(req, "_flashes", messages)
}

// Flashes returns a slice of flash messages from session data
func (scs SCS) Flashes(req *http.Request) []session.Message {
	var messages []session.Message
	scs.PopLoad(req, "_flashes", &messages)
	return messages
}

// Load get value from session data and unmarshal it into result
func (scs SCS) Load(req *http.Request, key string, result interface{}) error {
	value := scs.Get(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// PopLoad pop value from session data and unmarshal it into result
func (scs SCS) PopLoad(req *http.Request, key string, result interface{}) error {
	value := scs.Pop(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// Middleware returns a new session manager middleware instance
func (scs SCS) Middleware(handler http.Handler) http.Handler {
	return scssession.Manage(scs.Engine, scs.Options...)(handler)
}
