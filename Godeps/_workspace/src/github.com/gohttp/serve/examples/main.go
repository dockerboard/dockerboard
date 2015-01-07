package main

import "github.com/gohttp/logger"
import "github.com/gohttp/serve"
import "github.com/gohttp/app"
import "net/http"

func main() {
	a := app.New()

	a.Use(logger.New())
	a.Use(serve.New("examples"))

	a.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))

	a.Listen(":3000")
}
