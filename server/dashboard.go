package dashboard

import (
	"fmt"
	"net/http"
	"path"
	"runtime"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"
)

func ApiIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Coming soon!\n")
}

func Serve() {
	_, filename, _, _ := runtime.Caller(0)
	// Note config dir
	dir := path.Join(path.Dir(filename), "../client/src")

	app := app.New()
	app.Use(logger.New())
	app.Use(serve.New(dir))
	app.Get("/api/v1", ApiIndex)
	app.Listen(":3333")
}
