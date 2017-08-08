package gorilla

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	gorillaContext "github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/qor/qor/utils"
	"github.com/qor/session"
)

// New initialize session manager for gorilla
func New(sessionName string, store sessions.Store) *Gorilla {
	return &Gorilla{SessionName: sessionName, Store: store}
}

// Gorilla session manager struct for gorilla
type Gorilla struct {
	SessionName string
	Store       sessions.Store
}

var (
	writer utils.ContextKey = "gorilla_writer"
	reader utils.ContextKey = "gorilla_reader"
)

func (gorilla Gorilla) getSession(req *http.Request) (*sessions.Session, error) {
	if r, ok := req.Context().Value(reader).(*http.Request); ok {
		return gorilla.Store.Get(r, gorilla.SessionName)
	}
	return gorilla.Store.Get(req, gorilla.SessionName)
}

func (gorilla Gorilla) saveSession(req *http.Request) {
	if session, err := gorilla.getSession(req); err == nil {
		if w, ok := req.Context().Value(writer).(http.ResponseWriter); ok {
			if err := session.Save(req, w); err != nil {
				fmt.Printf("No error should happen when saving session data, but got %v", err)
			}
		}
	}
}

// Add value to session data, if value is not string, will marshal it into JSON encoding and save it into session data.
func (gorilla Gorilla) Add(req *http.Request, key string, value interface{}) error {
	defer gorilla.saveSession(req)

	session, err := gorilla.getSession(req)
	if err != nil {
		return err
	}

	if str, ok := value.(string); ok {
		session.Values[key] = str
	} else {
		result, _ := json.Marshal(value)
		session.Values[key] = string(result)
	}

	return nil
}

// Pop value from session data
func (gorilla Gorilla) Pop(req *http.Request, key string) string {
	defer gorilla.saveSession(req)

	if session, err := gorilla.getSession(req); err == nil {
		if value, ok := session.Values[key]; ok {
			delete(session.Values, key)
			return fmt.Sprint(value)
		}
	}
	return ""
}

// Get value from session data
func (gorilla Gorilla) Get(req *http.Request, key string) string {
	if session, err := gorilla.getSession(req); err == nil {
		if value, ok := session.Values[key]; ok {
			return fmt.Sprint(value)
		}
	}
	return ""
}

// Flash add flash message to session data
func (gorilla Gorilla) Flash(req *http.Request, message session.Message) error {
	var messages []session.Message
	if err := gorilla.Load(req, "_flashes", &messages); err != nil {
		return err
	}
	messages = append(messages, message)
	return gorilla.Add(req, "_flashes", messages)
}

// Flashes returns a slice of flash messages from session data
func (gorilla Gorilla) Flashes(req *http.Request) []session.Message {
	var messages []session.Message
	gorilla.PopLoad(req, "_flashes", &messages)
	return messages
}

// Load get value from session data and unmarshal it into result
func (gorilla Gorilla) Load(req *http.Request, key string, result interface{}) error {
	value := gorilla.Get(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// PopLoad pop value from session data and unmarshal it into result
func (gorilla Gorilla) PopLoad(req *http.Request, key string, result interface{}) error {
	value := gorilla.Pop(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// Middleware returns a new session manager middleware instance
func (gorilla Gorilla) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer gorillaContext.Clear(req)
		ctx := context.WithValue(context.WithValue(req.Context(), writer, w), reader, req)
		handler.ServeHTTP(w, req.WithContext(ctx))
	})
}
