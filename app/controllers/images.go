package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Images Container.
type ImagesController struct{}

// Query Parameters for list images.
type ImagesIndexOptions struct {
	All     string `url:"all"`
	Filters string `url:"filters"`
}

// Query Parameters for remove an images.
type ImagesDestoryOptions struct {
	Force   string `url:"force"`
	NoPrune string `url:"noprune"`
}

// List Images.
// GET /images
func (ic *ImagesController) Index(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := NewRequest("GET", "/images/json", params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(ImagesIndexOptions{
		All:     params.Get("all"),
		Filters: params.Get("filters"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

func (ic *ImagesController) Create(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/images/create")
	q, err := NewRequest("POST", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Parameters(r.PostForm)
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

// Inspect an image.
// GET /images/:id
func (ic *ImagesController) Show(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, _ := url.QueryUnescape(params.Get(":id"))
	endpoint := fmt.Sprintf("/images/%s/json", id)
	q, err := NewRequest("GET", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

// Remove an image.
// DELETE /images/:id
func (ic *ImagesController) Destroy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/images/%s", params.Get(":id"))
	q, err := NewRequest("DELETE", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(ImagesDestoryOptions{
		Force:   params.Get("force"),
		NoPrune: params.Get("noprune"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 404, 409, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}
