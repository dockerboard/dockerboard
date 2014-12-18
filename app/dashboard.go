package dashboard

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"

	"github.com/dockerboard/dockerboard/app/controllers"
	"github.com/goocean/methodoverride"
)

func APIIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Coming soon!\n")
}

func Serve() {
	// TODO @fundon, config dir or from ENV
	dir := path.Join(path.Dir("/bluewhale/dist/"))

	// Controllers
	containersController := controllers.NewContainers()
	imagesController := controllers.NewImages()

	app := app.New()
	app.Use(logger.New())
	app.Use(serve.New(dir))
	app.Use(methodoverride.New())
	app.Get("/api", APIIndex)
	app.Get("/api/containers", containersController.Index)
	app.Post("/api/containers", containersController.Create)
	app.Get("/api/containers/:id", containersController.Show)
	app.Del("/api/containers/:id", containersController.Destroy)
	app.Get("/api/images", imagesController.Index)
	app.Post("/api/images", imagesController.Create)
	app.Get("/api/images/:id", imagesController.Show)
	app.Del("/api/images/:id", imagesController.Destroy)
	app.Get("/api/apps", controllers.NewApps().Index)
	app.Listen(":8001")
}
