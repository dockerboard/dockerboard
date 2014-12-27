package controllers

import (
	"io"
	"net/http"
)

type SystemController struct{}

func (s *SystemController) Info(w http.ResponseWriter, r *http.Request) {
	endpoint := "/info"
	q, err := NewRequest("GET", endpoint)
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
