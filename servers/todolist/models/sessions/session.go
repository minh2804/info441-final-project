package sessions

import (
	"errors"
	"net/http"
	"strings"
)

// ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + HeaderAuthorization + " header")

// ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

const HeaderAuthorization = "Authorization"

const paramAuthorization = "auth"
const schemeBearer = "Bearer "

// BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
// Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	sid, err := NewSessionID(signingKey)
	if err != nil {
		return sid, err
	}

	if err := store.Save(sid, sessionState); err != nil {
		return sid, err
	}

	authToken := schemeBearer + sid.String()
	w.Header().Add(HeaderAuthorization, authToken)

	return sid, nil
}

// GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	authToken := r.Header.Get(HeaderAuthorization)
	if len(authToken) == 0 {
		authToken = r.URL.Query().Get("auth")
		if len(authToken) == 0 {
			return InvalidSessionID, ErrNoSessionID
		}
	}

	if !strings.HasPrefix(authToken, schemeBearer) {
		return InvalidSessionID, ErrInvalidScheme
	}

	sid := strings.TrimPrefix(authToken, schemeBearer)
	return ValidateID(sid, signingKey)
}

// GetState extracts the SessionID from the request,
// gets the associated state from the provided store into
// the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	sid, err := GetSessionID(r, signingKey)
	if err != nil {
		return sid, err
	}
	return sid, store.Get(sid, sessionState)
}

// EndSession extracts the SessionID from the request,
// and deletes the associated data in the provided store, returning
// the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	sid, err := GetSessionID(r, signingKey)
	if err != nil {
		return sid, err
	}
	return sid, store.Delete(sid)
}
