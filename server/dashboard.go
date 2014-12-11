package dashboard

import (
	"fmt"
	"net/http"
	"path"
	"runtime"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"

	"github.com/dockerboard/dockerboard/server/controllers"
)

func APIIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Coming soon!\n")
}

func Serve() {
	_, filename, _, _ := runtime.Caller(0)
	// Note config dir
	dir := path.Join(path.Dir(filename), "../client/src")

	app := app.New()
	app.Use(logger.New())
	app.Use(serve.New(dir))
	app.Get("/api", APIIndex)
	app.Get("/api/containers", controllers.NewContainers().Index)
	app.Get("/api/containers/:id", controllers.NewContainers().Show)
	app.Get("/api/images", controllers.NewImages().Index)
	app.Get("/api/images/:id", controllers.NewImages().Show)
	app.Get("/api/apps", controllers.NewApps().Index)
	app.Listen(":3333")
}
