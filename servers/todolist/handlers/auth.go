package handlers

import (
	"encoding/json"
	"errors"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/users"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

const SuccessSignOut = "signed out"

var ErrContentTypeNotJSON = errors.New("content type must be type type 'application/json'")
var ErrForbiddenAccess = errors.New("you are not allowed to make this change")
var ErrInternal = errors.New("something went wrong, please try again later")
var ErrInvalidBody = errors.New("content body is invalid")
var ErrInvalidCredentials = errors.New("email and password do not match")
var ErrInvalidResourcePath = errors.New("invalid resource path")
var ErrRequestMethodNotAllowed = errors.New("request method not allowed")
var ErrUnauthorized = errors.New("please sign in")

// SpecificUserHandler handles the following:
// POST - Create a new user
func (ctx *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Validate request
		if !strings.HasPrefix(r.Header.Get(ContentTypeHeader), ContentTypeJSON) {
			http.Error(w, ErrContentTypeNotJSON.Error(), http.StatusUnsupportedMediaType)
			return
		}

		// Decode request body
		newUser := &users.NewUser{}
		if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
			http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
			return
		}
		unregisteredUser, err := newUser.ToUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the newly created user to store
		registeredUser, err := ctx.UserStore.Insert(unregisteredUser)
		if err != nil {
			if err == users.ErrUserAlreadyExisted {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Begin a session for the newly registered user
		if _, err := sessions.BeginSession(
			ctx.SigningKey,
			ctx.SessionStore,
			sessions.NewSessionState(registeredUser),
			w); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(registeredUser); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

// SpecificUserHandler handles the following:
// GET - Get the current user
// PATCH - Update the current user
func (ctx *HandlerContext) SpecificUserHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	// Ensure user is logged in
	if currentSession.User == nil {
		http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	// Extract id from path
	var requestedUserID int64
	if mux.Vars(r)["userID"] == "me" {
		requestedUserID = currentSession.User.ID
	} else {
		id, err := strconv.ParseInt(mux.Vars(r)["userID"], 10, 64)
		if err != nil {
			http.Error(w, ErrInvalidResourcePath.Error(), http.StatusBadRequest)
			return
		}
		requestedUserID = id
	}

	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Search for the requested user
		requestedUser, err := ctx.UserStore.GetByID(requestedUserID)
		if err != nil {
			if err == users.ErrUserNotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(requestedUser); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPatch:
		// Validate request
		if (mux.Vars(r)["userID"] != "me") || (requestedUserID != currentSession.User.ID) {
			http.Error(w, ErrForbiddenAccess.Error(), http.StatusBadRequest)
			return
		}
		if !strings.HasPrefix(r.Header.Get(ContentTypeHeader), ContentTypeJSON) {
			http.Error(w, ErrContentTypeNotJSON.Error(), http.StatusUnsupportedMediaType)
			return
		}

		// Decode request body
		requestedUpdates := &users.Updates{}
		if err := json.NewDecoder(r.Body).Decode(requestedUpdates); err != nil {
			http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
			return
		}

		// Apply updates to store
		updatedUser, err := ctx.UserStore.Update(requestedUserID, requestedUpdates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Apply updates to session store
		currentSession.User = updatedUser
		if err := ctx.SessionStore.Save(sessionID, currentSession); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

// SessionsHandler handles the following:
// POST - Create a new authenticated session given valid credentials
func (ctx *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Validate request
		if !strings.HasPrefix(r.Header.Get(ContentTypeHeader), ContentTypeJSON) {
			http.Error(w, ErrContentTypeNotJSON.Error(), http.StatusUnsupportedMediaType)
			return
		}

		// Decode request body
		providedCredentials := &users.Credentials{}
		if err := json.NewDecoder(r.Body).Decode(providedCredentials); err != nil {
			http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
			return
		}

		// Validate credentials
		currentUser, err := ctx.UserStore.GetByUsername(providedCredentials.Username)
		if err != nil {
			// Pretend to process valid credentials
			bcrypt.CompareHashAndPassword([]byte(providedCredentials.Password), []byte(providedCredentials.Password))

			http.Error(w, ErrInvalidCredentials.Error(), http.StatusUnauthorized)
			return
		}
		if err := currentUser.Authenticate(providedCredentials.Password); err != nil {
			http.Error(w, ErrInvalidCredentials.Error(), http.StatusUnauthorized)
			return
		}

		// Begin a session for the logged in user
		if _, err := sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, sessions.NewSessionState(currentUser), w); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(currentUser); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

// SpecificSessionHandler handles the following:
// DELETE - Sign out user
func (ctx *HandlerContext) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		// Validate request
		if mux.Vars(r)["sessionID"] != "mine" {
			http.Error(w, ErrInvalidResourcePath.Error(), http.StatusForbidden)
			return
		}

		// End the session
		if _, err := sessions.EndSession(r, ctx.SigningKey, ctx.SessionStore); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Response to request
		w.Write([]byte(SuccessSignOut))
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}
