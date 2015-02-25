package controllers

import (
	"fmt"
	"io"
	"net/http"
)

// Container Actions Controller
type ContainerActionsController struct{}

// Query Parameters for start/stop/restart a container.
type stopOptions struct {
	T string `url:"t"`
}

// Query Parameters for kill a container.
type killOptions struct {
	Signal string `url:"signal"`
}

// Query Parameters for get container logs.
type logsOptions struct {
	Follow     string `url:"follow"`
	Stdout     string `url:"stdout"`
	Stderr     string `url:"stderr"`
	Timestamps string `url:"timestamps"`
	Tail       string `url:"tail"`
}

// Query Parameters For List processes running inside a container.
type topOptions struct {
	PS_Args string `url:"ps_args"`
}

// Query Parameters For rename a container.
type renameOptions struct {
	Name string `url:"name"`
}

// Start a container.
// POST /containers/:id/start
func (ca *ContainerActionsController) Start(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/start", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("host"))
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

// Stop a container.
// POST /containers/:id/stop
func (ca *ContainerActionsController) Stop(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/stop", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("host"))
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

// Restart a container.
// POST /containers/:id/restart
func (ca *ContainerActionsController) Restart(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/restart", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("host"))
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

// Pause a container.
// POST /containers/:id/pause
func (ca *ContainerActionsController) Pause(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/pause", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("hosts"))
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

// Unpause a container.
// POST /containers/:id/unpause
func (ca *ContainerActionsController) UnPause(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/unpause", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("host"))
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

// Get container logs.
// GET /containers/:id/logs
func (ca *ContainerActionsController) Logs(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/logs", params.Get(":id"))
	q, err := NewRequest("GET", endpoint, params.Get("host"))
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

// Kill a container.
// POST /containers/:id/kill
func (ca *ContainerActionsController) Kill(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/kill", params.Get(":id"))
	q, err := NewRequest("POST", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(killOptions{
		Signal: params.Get("signal"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(204, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}

// List processes running inside a container.
// GET /containers/:id/top
func (ca *ContainerActionsController) Top(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/top", params.Get(":id"))
	q, err := NewRequest("GET", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(topOptions{
		PS_Args: params.Get("ps_args"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 404, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}

// Rename the container id to a new_name.
// POST /containers/:id/rename
// v1.17
func (ca *ContainerActionsController) Rename(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/stats", params.Get(":id"))
	q, err := NewRequest("GET", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q.Query(renameOptions{
		Name: params.Get("name"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(204, 404, 409, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}

// Get container stats based on resource usage.
// GET /containers/:id/stats
// v1.17
func (ca *ContainerActionsController) Stats(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/containers/%s/stats", params.Get(":id"))
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
	io.Copy(w, b)
}
