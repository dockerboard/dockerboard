package controllers

import (
	"io"
	"net/http"
)

// System Controller.
type SystemController struct{}

// Display system-wide information.
// GET /info
func (s *SystemController) Info(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := "/info"
	q, err := NewRequest("GET", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if !q.ValidateStatusCode(204, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}
