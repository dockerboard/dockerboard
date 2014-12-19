package controllers

import (
	"fmt"
	"io"
	"net/http"
)

type ContainersIndexOptions struct {
	All    string `url:"all"`
	Limit  string `url:"limit"`
	Size   string `url:"size"`
	Since  string `url:"since"`
	Before string `url:"before"`
}

type ContainersDestoryOptions struct {
	Force string `url:"force"`
	V     string `url:"v"`
}

type ContainersController struct{}

func (cc *ContainersController) Index(w http.ResponseWriter, r *http.Request) {
	q, err := NewRequest("GET", "/containers/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := r.URL.Query()
	q.Query(ContainersIndexOptions{
		All:    params.Get("all"),
		Limit:  params.Get("limit"),
		Size:   params.Get("size"),
		Since:  params.Get("since"),
		Before: params.Get("before"),
	})
	b, err := q.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

func (cc *ContainersController) Create(w http.ResponseWriter, r *http.Request) {
}

func (cc *ContainersController) Show(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/containers/%s/json", r.URL.Query().Get(":id"))
	q, err := NewRequest("GET", endpoint)
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

func (cc *ContainersController) Destroy(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/containers/%s", r.URL.Query().Get(":id"))
	q, err := NewRequest("DELETE", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := r.URL.Query()
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
