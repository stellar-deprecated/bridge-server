package manager

import (
	"github.com/gorilla/sessions"
	"github.com/qor/session"
	"github.com/qor/session/gorilla"
)

// SessionManager default session manager
var SessionManager session.ManagerInterface = gorilla.New("_session", sessions.NewCookieStore([]byte("secret")))
