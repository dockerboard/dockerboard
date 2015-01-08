package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/docker/docker/pkg/parsers"
)

const DEFAULTTAG = "latest"

// Images Container.
type ImagesController struct{}

// Query Parameters for list images.
type indexOptions struct {
	All     string `url:"all"`
	Filters string `url:"filters"`
}

// Query Parameters for create an image.
type createOptions struct {
	FromImage string `url:"fromImage"`
	FromSrc   string `url:"fromSrc"`
	Repo      string `url:"repo"`
	Tag       string `url:"tag"`
	Registry  string `url:"registry"`
}

// Query Parameters for remove an images.
type destoryOptions struct {
	Force   string `url:"force"`
	NoPrune string `url:"noprune"`
}

// Query Parameters for Search images on Docker Hub.
type searchOptions struct {
	Term string `url:"term"`
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
	q.Query(indexOptions{
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

// Create an image
// POST /images
// /docker_remote_api_v1.16/#create-an-image
func (ic *ImagesController) Create(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	endpoint := fmt.Sprintf("/images/create")
	q, err := NewRequest("POST", endpoint, params.Get("host"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	auth := r.Header.Get("Authorization")
	if auth != "" {
		q.Set("X-Registry-Auth", auth)
	}

	var (
		image   = params.Get("fromImage")
		repo    = params.Get("repo")
		tag     = params.Get("tag")
		src     = params.Get("fromSrc")
		options = createOptions{Registry: params.Get("registry")}
	)

	if image != "" { // pull
		if tag == "" {
			image, tag = parsers.ParseRepositoryTag(image)
		}
		options.FromImage = image
		options.Tag = defaultTo(tag, DEFAULTTAG)
	} else { // import
		if tag == "" {
			repo, tag = parsers.ParseRepositoryTag(repo)
		}
		options.FromSrc = src
		options.Repo = repo
		options.Tag = defaultTo(tag, DEFAULTTAG)
	}

	q.Query(options)
	q.Timeout(0)
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	//w.Header().Set("Content-Type", "application/json;charset=utf-8")
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
	q.Query(destoryOptions{
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

// Search images
// GET /images/search
// Search for an image on Docker Hub https://hub.docker.com/.
func (ic *ImagesController) Search(w http.ResponseWriter, r *http.Request) {
	endpoint := "/images/search"
	q, err := NewRequest("GET", endpoint, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := r.URL.Query()
	q.Query(searchOptions{
		Term: params.Get("term"),
	})
	b, err := q.Do()
	if !q.ValidateStatusCode(200, 500) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(q.StatusCode)
	io.Copy(w, b)
}
