package controllers

import (
	"io"
	"net/http"
)

// Host Actions Controller.
type HostActionsController struct{}

// Ping the docker server.
// GET /hosts/:id/ping
func (ha *HostActionsController) Ping(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := NewRequest("GET", "/_ping", params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

// Show the docker version information
// GET /hosts/:id/version
func (ha *HostActionsController) Version(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := NewRequest("GET", "/version", params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}