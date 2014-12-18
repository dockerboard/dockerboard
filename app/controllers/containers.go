package controllers

import (
	"fmt"
	"io"
	"net/http"
)

type ContainersOptions struct {
	All    string `url:"all"`
	Limit  string `url:"limit"`
	Size   string `url:"size"`
	Since  string `url:"since"`
	Before string `url:"before"`
}

type ContainersController struct{}

func (cc *ContainersController) Index(w http.ResponseWriter, r *http.Request) {
	q, err := NewRequest("GET", "/containers/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := r.URL.Query()
	q.Query(ContainersOptions{
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

func (cc *ContainersController) Show(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/containers/%s/json", r.URL.Query().Get(":id"))
	q, err := NewRequest("GET", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}
