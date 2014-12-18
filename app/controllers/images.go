package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ImagesOptions struct {
	All     string `url:"all"`
	Filters string `url:"filters"`
}

type ImagesController struct{}

func (ic *ImagesController) Index(w http.ResponseWriter, r *http.Request) {
	q, err := NewRequest("GET", "/images/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := r.URL.Query()
	q.Query(ImagesOptions{
		All:     params.Get("all"),
		Filters: params.Get("filters"),
	})
	b, err := q.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

func (ic *ImagesController) Create(w http.ResponseWriter, r *http.Request) {
}

func (ic *ImagesController) Show(w http.ResponseWriter, r *http.Request) {
	id, _ := url.QueryUnescape(r.URL.Query().Get(":id"))
	endpoint := fmt.Sprintf("/images/%s/json", id)
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

func (ic *ImagesController) Destroy(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/images/%s", r.URL.Query().Get(":id"))
	q, err := NewRequest("DELETE", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	io.Copy(w, b)
}
