package handlers

import "net/http"

// Returns stats (tasks created, tasks done) for the entire lifespan of the account
func (h *HandlerContext) AllStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	} else {
		http.Error(w, "Only Get Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Returns stats (tasks created, tasks done) for a specific length of time (year, month, week)
func (h *HandlerContext) SpecificStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.UserStore.

	} else {
		http.Error(w, "Only Get Allowed", http.StatusMethodNotAllowed)
		return
	}
}
