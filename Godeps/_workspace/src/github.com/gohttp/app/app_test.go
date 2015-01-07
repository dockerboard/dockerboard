package app

import . "github.com/franela/go-supertest"
import "net/http/httptest"
import "net/http"
import "testing"
import "fmt"

// test GET
func TestGet(t *testing.T) {
	app := New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	s := httptest.NewServer(app)

	NewRequest(s.URL).
		Get("/").
		Expect(200, "hello")
}

// test HEAD
func TestHead(t *testing.T) {
	app := New()

	app.Head("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	s := httptest.NewServer(app)

	NewRequest(s.URL).
		Head("/").
		Expect(200)
}

// test HEAD for GET route
func TestHeadGet(t *testing.T) {
	app := New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	s := httptest.NewServer(app)

	NewRequest(s.URL).
		Head("/").
		Expect(200)
}

// test route precedence
func TestPrecedence(t *testing.T) {
	app := New()

	app.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	app.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})

	s := httptest.NewServer(app)

	NewRequest(s.URL).
		Get("/foo").
		Expect(200, "hello")
}

// test many routes
func TestMany(t *testing.T) {
	app := New()

	app.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	app.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})

	s := httptest.NewServer(app)

	NewRequest(s.URL).
		Get("/foo").
		Expect(200, "hello")

	NewRequest(s.URL).
		Get("/bar").
		Expect(200, "world")
}

// test params
func TestParams(t *testing.T) {
	app := New()

	app.Get("/user/:name/pet/:pet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get(":name")
		pet := r.URL.Query().Get(":pet")
		fmt.Fprint(w, "user %s's pet %s", name, pet)
	})

	s := httptest.NewServer(app)

	NewRequest(s.URL).
		Get("/user/tobi/pet/loki").
		Expect(200, "user tobi's pet loki")
}
