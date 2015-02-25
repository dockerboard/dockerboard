package controllers

import (
	"fmt"
	"io"
	"net/http"
)

// Containers Controller.
type ContainersController struct{}

// Query Parameters for list containers.
type ContainersIndexOptions struct {
	All     string `url:"all"`
	Limit   string `url:"limit"`
	Size    string `url:"size"`
	Since   string `url:"since"`
	Before  string `url:"before"`
	Filters string `url:"filters"`
}

// Query Parameters for remove a container.
type ContainersDestoryOptions struct {
	Force string `url:"force"`
	V     string `url:"v"`
}

// List containers.
// GET /containers
func (cc *ContainersController) Index(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := NewRequest("GET", "/containers/json", params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(ContainersIndexOptions{
		All:     params.Get("all"),
		Limit:   params.Get("limit"),
		Size:    params.Get("size"),
		Since:   params.Get("since"),
		Before:  params.Get("before"),
		Filters: params.Get("filters"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 400, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

func (cc *ContainersController) Create(w http.ResponseWriter, r *http.Request) {
}

// Inspect a container.
// GET /containers/:id
func (cc *ContainersController) Show(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/json", params.Get(":id"))
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

// Remove a container.
// DELETE /containers/:id
func (cc *ContainersController) Destroy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s", params.Get(":id"))
	q, err := NewRequest("DELETE", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(ContainersDestoryOptions{
		Force: params.Get("force"),
		V:     params.Get("v"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 404, 409, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}