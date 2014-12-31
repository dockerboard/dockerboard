package controllers

import (
	"fmt"
	"io"
	"net/http"
)

type ImageActionsController struct{}

func (ia *ImageActionsController) History(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/images/%s/history", params.Get(":id"))
	q, err := NewRequest("GET", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if !q.ValidateStatusCode(204, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}
