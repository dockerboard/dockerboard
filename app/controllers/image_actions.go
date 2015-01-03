package controllers

import (
	"fmt"
	"io"
	"net/http"
)

// Image Actions Controller.
type ImageActionsController struct{}

// Query Parameters for tag an image into a repository.
type tagOptions struct {
	Force string `url:"force"`
	Repo  string `url:"repo"`
	Tag   string `url:"tag"`
}

// Get the history of an image.
// GET /images/:id/history
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

// Tag an image into a repository
// POST /images/:id/tag
func (ia *ImageActionsController) Tag(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/images/%s/tag", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(tagOptions{
		Force: params.Get("force"),
		Repo:  params.Get("repo"),
		Tag:   params.Get("tag"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(201, 400, 404, 409, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}
