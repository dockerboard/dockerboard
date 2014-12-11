package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-libs/quest"
)

type ContainersController struct{}

func (cc *ContainersController) Index(w http.ResponseWriter, r *http.Request) {
	q, err := NewRequest(quest.GET, "/containers/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, b)
}

func (cc *ContainersController) Show(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/containers/%s/json", r.URL.Query().Get(":id"))
	q, err := NewRequest(quest.GET, endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, b)
}
