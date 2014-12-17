package controllers

import (
	"fmt"
	"io"
	"net/http"
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
	q.Query(ImagesOptions{
		All:     r.URL.Query().Get("all"),
		Filters: r.URL.Query().Get("filters"),
	})
	b, err := q.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.Copy(w, b)
}

func (ic *ImagesController) Show(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/images/%s/json", r.URL.Query().Get(":id"))
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
