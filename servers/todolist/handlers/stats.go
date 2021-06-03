package handlers

import (
	"encoding/json"
	"fmt"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/stats"
	"info441-final-project/servers/todolist/models/tasks"
	"net/http"
	"strings"
)

// Returns stats (tasks created, tasks done) for the entire lifespan of the account
func (h *HandlerContext) AllStatsHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	switch r.Method {
	case http.MethodGet:
		user := currentSession.User.ID
		if user == 0 {
			http.Error(w, "no user provided", http.StatusBadRequest)
			return
		}

		createdTasks, createdErr := h.StatStore.GetCompletedByID(user)
		if createdErr != nil {
			http.Error(w, fmt.Sprintf("error getting completed tasks: %v", createdErr), http.StatusInternalServerError)
			return
		}

		completedTasks, completedErr := h.StatStore.GetAllByID(user)

		if completedErr != nil {
			http.Error(w, fmt.Sprintf("error getting completed tasks: %v", completedErr), http.StatusInternalServerError)
			return
		}

		results := &stats.QueryResults{
			Completed: completedTasks,
			Created:   createdTasks,
		}

		// respond to the client
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

// Returns stats (tasks created, tasks done) for a specific length of time (year, month, week, custom)
func (h *HandlerContext) PeriodicStatsHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	switch r.Method {
	case http.MethodGet:
		requestedPeriod := strings.ToLower(r.URL.Path)
		if requestedPeriod != "year" && requestedPeriod != "month" && requestedPeriod != "week" && requestedPeriod != "custom" {
			http.Error(w, fmt.Sprintf("%s is not a reconized period, please input only year, month, custom, or week", requestedPeriod), http.StatusBadRequest)
			return
		}
		user := currentSession.User.ID
		if user == 0 {
			http.Error(w, "no user provided", http.StatusBadRequest)
			return
		}

		if requestedPeriod == "custom" {
			SpecificStatsHandler(h, w, r, sessionID, currentSession)
			return
		}

		createdTasks := []*tasks.Task{}

		completedTasks := []*tasks.Task{}

		if requestedPeriod == "year" {
			createdTasks, _ = h.StatStore.GetAllWithinYear(user)
		}

		if requestedPeriod == "month" {
			createdTasks, _ = h.StatStore.GetAllWithinMonth(user)
		}

		if requestedPeriod == "week" {
			createdTasks, _ = h.StatStore.GetAllWithinWeek(user)
		}

		if requestedPeriod == "year" {
			completedTasks, _ = h.StatStore.GetCompletedWithinYear(user)
		}

		if requestedPeriod == "month" {
			completedTasks, _ = h.StatStore.GetCompletedWithinMonth(user)
		}

		if requestedPeriod == "week" {
			completedTasks, _ = h.StatStore.GetCompletedWithinWeek(user)

		}
		results := &stats.QueryResults{
			Completed: completedTasks,
			Created:   createdTasks,
		}

		// respond to the client
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

// Returns stats (tasks created, tasks done) for between two dates (start and stop in query params)
func SpecificStatsHandler(h *HandlerContext, w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	switch r.Method {
	case http.MethodGet:
		user := currentSession.User.ID
		if user == 0 {
			http.Error(w, "no user provided", http.StatusBadRequest)
			return
		}

		begin, beginErr := r.URL.Query()["start"]
		if !beginErr {
			http.Error(w, fmt.Sprintf("error obtaining the begin date: %v", beginErr), http.StatusInternalServerError)
			return
		}
		if len(begin[0]) < 1 {
			http.Error(w, "no begin date supplied", http.StatusBadRequest)
			return
		}

		end, endErr := r.URL.Query()["stop"]
		if !endErr {
			http.Error(w, fmt.Sprintf("error obtaining the end date: %v", endErr), http.StatusInternalServerError)
			return
		}
		if len(end[0]) < 1 {
			http.Error(w, "no end date supplied", http.StatusBadRequest)
			return
		}

		createdTasks, createdErr := h.StatStore.GetCompletedBetweenDates(user, begin[0], end[0])
		if createdErr != nil {
			http.Error(w, fmt.Sprintf("error getting completed tasks: %v", createdErr), http.StatusInternalServerError)
			return
		}

		completedTasks, completedErr := h.StatStore.GetAllByID(user)

		if completedErr != nil {
			http.Error(w, fmt.Sprintf("error getting completed tasks: %v", completedErr), http.StatusInternalServerError)
			return
		}

		results := &stats.QueryResults{
			Completed: completedTasks,
			Created:   createdTasks,
		}

		// respond to the client
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}
