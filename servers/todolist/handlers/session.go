package handlers

import (
	"info441-final-project/servers/todolist/models/sessions"
	"net/http"

	"github.com/gorilla/mux"
)

type SessionHandlerFunc func(http.ResponseWriter, *http.Request, *sessions.SessionState)

// Struct that allows this type to inherit all of the methods of http.ServeMux
type SessionMux struct {
	ctx *HandlerContext
	mux.Router
}

// Constructor for a new instance of AuthenticatedMux
func NewSessionMux(ctx *HandlerContext) *SessionMux {
	return &SessionMux{ctx: ctx}
}

// Newly defined HandlerFunc that simulates that of http.HandleFunc
func (sm *SessionMux) HandleSessionFunc(pattern string, handler SessionHandlerFunc) {
	sm.HandleFunc(pattern, sm.ensureSession(handler))
}

// This is an adapter function that will ensure all handlers will have a session created,
// if a session already exists for the handler, then nothing will happen.
// The session state is essential an array of Task (i.e., todo list)
func (sm *SessionMux) ensureSession(handlerFunc SessionHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionState := &sessions.SessionState{}
		_, err := sessions.GetState(r, sm.ctx.SigningKey, sm.ctx.SessionStore, sessionState)
		if err != nil {
			if err == sessions.ErrNoSessionID {
				// Create a new session with empty todo list for the session state
				sessionState = sessions.NewTemporarySessionState()
				if _, err := sessions.BeginSession(
					sm.ctx.SigningKey,
					sm.ctx.SessionStore,
					sessionState,
					w); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				// Copy the newly-created session token over to request
				r.Header.Add(sessions.HeaderAuthorization, w.Header().Get(sessions.HeaderAuthorization))
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		handlerFunc(w, r, sessionState)
	}
}
