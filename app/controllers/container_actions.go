package controllers

import (
	"fmt"
	"io"
	"net/http"
)

type ContainerActionsController struct{}

type stopOptions struct {
	T string `url:"t"`
}

type logsOptions struct {
	Follow     string `url:"follow"`
	Stdout     string `url:"stdout"`
	Stderr     string `url:"stderr"`
	Timestamps string `url:"timestamps"`
	Tail       string `url:"tail"`
}

func (ca *ContainerActionsController) Start(w http.ResponseWriter, r *http.Request) {
	endpoint := fmt.Sprintf("/containers/%s/start", r.URL.Query().Get(":id"))
	q, err := NewRequest("POST", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := q.Do()
	if !q.ValidateStatusCode(204, 304, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}

func (ca *ContainerActionsController) Stop(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/stop", params.Get(":id"))
	q, err := NewRequest("POST", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(stopOptions{
		T: params.Get("t"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 304, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}

func (ca *ContainerActionsController) Restart(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/restart", params.Get(":id"))
	q, err := NewRequest("POST", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(stopOptions{
		T: params.Get("t"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 304, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}

func (ca *ContainerActionsController) Pause(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/pause", params.Get(":id"))
	q, err := NewRequest("POST", endpoint)
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

func (ca *ContainerActionsController) UnPause(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/unpause", params.Get(":id"))
	q, err := NewRequest("POST", endpoint)
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

func (ca *ContainerActionsController) Logs(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/logs", params.Get(":id"))
	q, err := NewRequest("GET", endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(logsOptions{
		Follow:     params.Get("follow"),
		Stdout:     params.Get("stdout"),
		Stderr:     params.Get("stderr"),
		Timestamps: params.Get("timestamps"),
		Tail:       params.Get("tail"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(204, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	io.Copy(w, b)
}
