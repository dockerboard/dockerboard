package dashboard

import (
	"fmt"
	"net/http"
	"path"
	"runtime"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"

	"github.com/dockerboard/dockerboard/app/controllers"
)

func APIIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Coming soon!\n")
}

func Serve() {
	_, filename, _, _ := runtime.Caller(0)
	// Note config dir
	dir := path.Join(path.Dir(filename), "../static")

	// Controllers
	containersController := controllers.NewContainers()
	imagesController := controllers.NewImages()

	app := app.New()
	app.Use(logger.New())
	app.Use(serve.New(dir))
	app.Get("/api", APIIndex)
	app.Get("/api/containers", containersController.Index)
	app.Get("/api/containers/:id", containersController.Show)
	app.Get("/api/images", imagesController.Index)
	app.Get("/api/images/:id", imagesController.Show)
	app.Get("/api/apps", controllers.NewApps().Index)
	app.Listen(":8001")
}
