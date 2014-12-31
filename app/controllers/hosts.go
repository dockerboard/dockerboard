package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/dockerboard/dockerboard/app/models"
	"github.com/gohttp/response"
)

type HostsController struct{}

type hostForm struct {
	Name, Host string
}

func (h *HostsController) Index(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, hosts)
}

func (h *HostsController) Create(w http.ResponseWriter, r *http.Request) {
	var hf hostForm
	err := json.NewDecoder(r.Body).Decode(&hf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addr := hf.Host
	name := hf.Name
	host, _, u, err := GetHost(addr)
	if err != nil && err.Error() != "Not Found Host." {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if host != nil {
		host.Name = name
		host.URL = u
	} else {
		host = &models.Host{Name: name, URL: u}
		hosts = append(hosts, host)
	}
	response.JSON(w, host, 201)
}

func (h *HostsController) Destroy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, _ := url.QueryUnescape(params.Get(":id"))
	_, i, _, err := GetHost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hosts = append(hosts[:i], hosts[i+1:]...)
	response.JSON(w, id)
}
