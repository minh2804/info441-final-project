package middlewares

import (
	"info441-final-project/servers/todolist/handlers"
	"info441-final-project/servers/todolist/models/sessions"
	"net/http"

	"github.com/gorilla/mux"
)

// SessionHandlerFunc is a modifed function signature of http.HandlerFunc
type SessionHandlerFunc func(http.ResponseWriter, *http.Request, sessions.SessionID, *sessions.SessionState)

// Struct that allows this type to inherit all of the methods of http.ServeMux
type SessionMux struct {
	ctx *handlers.HandlerContext
	mux.Router
}

// Constructor for a new instance of AuthenticatedMux
func NewSessionMux(ctx *handlers.HandlerContext) *SessionMux {
	return &SessionMux{ctx: ctx}
}

// Newly defined HandlerFunc that simulates that of http.HandleFunc
func (sm *SessionMux) HandleSessionFunc(pattern string, handlerFunc SessionHandlerFunc) {
	sm.HandleFunc(pattern, sm.ensureSession(handlerFunc))
}

// This is an adapter function that will ensure all handlers will have a session created,
// if a session already exists for the handler, then nothing will happen.
// The session state is essential an array of Task (i.e., todo list)
func (sm *SessionMux) ensureSession(handlerFunc SessionHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession := &sessions.SessionState{}
		sessionID, err := sessions.GetState(r, sm.ctx.SigningKey, sm.ctx.SessionStore, currentSession)
		if err != nil {
			if err == sessions.ErrNoSessionID {
				// Create a new session with an empty todo list for the session state
				currentSession = sessions.NewTemporarySessionState()
				if sessionID, err = sessions.BeginSession(
					sm.ctx.SigningKey,
					sm.ctx.SessionStore,
					currentSession,
					w); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				// Copy the newly-created session token over to the request
				r.Header.Add(sessions.HeaderAuthorization, w.Header().Get(sessions.HeaderAuthorization))
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		handlerFunc(w, r, sessionID, currentSession)
	}
}
