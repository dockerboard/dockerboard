package app

import "github.com/justinas/alice"
import "github.com/bmizerany/pat"
import "net/http"

// App builder.
type App struct {
	mux   *pat.PatternServeMux
	chain alice.Chain
}

// New application.
func New() *App {
	return &App{
		mux:   pat.New(),
		chain: alice.New(),
	}
}

// Use the given middleware.
func (a *App) Use(mw ...alice.Constructor) {
	a.chain = a.chain.Append(mw...)
}

// Head will register a pattern with a handler for HEAD requests.
func (a *App) Head(path string, h interface{}) {
	a.mux.Head(path, handler(h))
}

// Get will register a pattern with a handler for GET requests.
// The handler given is also registered for HEAD requests.
func (a *App) Get(path string, h interface{}) {
	a.mux.Get(path, handler(h))
}

// Post will register a pattern with a handler for POST requests.
func (a *App) Post(path string, h interface{}) {
	a.mux.Post(path, handler(h))
}

// Put will register a pattern with a handler for PUT requests.
func (a *App) Put(path string, h interface{}) {
	a.mux.Put(path, handler(h))
}

// Del will register a pattern with a handler for DELETE requests.
func (a *App) Del(path string, h interface{}) {
	a.mux.Del(path, handler(h))
}

// Options will register a pattern with a handler for OPTIONS requests.
func (a *App) Options(path string, h interface{}) {
	a.mux.Options(path, handler(h))
}

// Listen on `addr`.
func (a *App) Listen(addr string) error {
	handler := a.chain.Then(a.mux)
	http.Handle("/", handler)
	return http.ListenAndServe(addr, nil)
}

// ServeHTTP implements http.Handler
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := a.chain.Then(a.mux)
	handler.ServeHTTP(w, r)
}

// coerce handler into an http.Handler.
func handler(h interface{}) http.Handler {
	switch h.(type) {
	case func(w http.ResponseWriter, r *http.Request):
		return http.HandlerFunc(h.(func(w http.ResponseWriter, r *http.Request)))
	case http.Handler:
		return h.(http.Handler)
	default:
		panic("invalid handler")
	}
}
