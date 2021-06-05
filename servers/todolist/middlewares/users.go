package middlewares

import (
	"info441-final-project/servers/todolist/handlers"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/users"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandleFunc is a modifed function signature of http.HandlerFunc
type SessionHandleFunc func(http.ResponseWriter, *http.Request, sessions.SessionID, *sessions.SessionState)

// Struct that allows this type to inherit all of the methods of http.ServeMux
type UserMux struct {
	ctx *handlers.HandlerContext
	mux.Router
}

// Constructor for a new instance of AuthenticatedMux
func NewUserMux(ctx *handlers.HandlerContext) *UserMux {
	return &UserMux{ctx: ctx}
}

// Newly defined HandlerFunc that simulates that of http.HandleFunc
func (sm *UserMux) HandleUserFunc(pattern string, handlerFunc SessionHandleFunc) {
	sm.HandleFunc(pattern, sm.ensureUser(handlerFunc))
}

// This is an adapter function that will ensure all handlers will have a user account,
// if a user already existed, then nothing will happen, else a new temporary user account will be created.
func (sm *UserMux) ensureUser(handlerFunc SessionHandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession := &sessions.SessionState{}
		sessionID, err := sessions.GetState(r, sm.ctx.SigningKey, sm.ctx.SessionStore, currentSession)
		if err != nil {
			if err == sessions.ErrNoSessionID || err == sessions.ErrStateNotFound {
				// Create a new session
				sessionID, err = sessions.BeginSession(sm.ctx.SigningKey, sm.ctx.SessionStore, sessions.NewTemporarySessionState(), w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Create a new temporary account
				newUser := &users.NewUser{
					Username:     sessionID.String(),
					Password:     sessionID.String(),
					PasswordConf: sessionID.String(),
					IsTemporary:  true,
				}
				unregisteredUser, err := newUser.ToUser()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				registeredUser, err := sm.ctx.UserStore.Insert(unregisteredUser)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Update current session
				currentSession = sessions.NewSessionState(registeredUser)
				if err := sm.ctx.SessionStore.Save(sessionID, currentSession); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Copy the newly-created session token over to the request
				r.Header.Add(sessions.HeaderAuthorization, w.Header().Get(sessions.HeaderAuthorization))
			} else if err == sessions.ErrInvalidScheme {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
				return
			}
		}
		handlerFunc(w, r, sessionID, currentSession)
	}
}
