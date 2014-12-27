package app

import (
	"fmt"
	"net/http"
	"os"
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
	bluewhale := os.Getenv("BLUEWHALE_DIST")
	if bluewhale == "" {
		bluewhale = "/bluewhale/dist/"
	}
	dir := path.Join(path.Dir(bluewhale))

	// Controllers
	containersController := controllers.NewContainers()
	containerActionsController := controllers.NewContainerActions()
	imagesController := controllers.NewImages()
	imageActionsController := controllers.NewImageActions()
	systemController := controllers.NewSystem()

	app := app.New()
	app.Use(logger.New())
	app.Use(serve.New(dir))
	app.Use(methodoverride.New())
	app.Get("/api", APIIndex)

	app.Get("/api/containers", containersController.Index)
	app.Post("/api/containers", containersController.Create)
	app.Get("/api/containers/:id", containersController.Show)
	app.Del("/api/containers/:id", containersController.Destroy)
	app.Post("/api/containers/:id/start", containerActionsController.Start)
	app.Post("/api/containers/:id/stop", containerActionsController.Stop)
	app.Post("/api/containers/:id/restart", containerActionsController.Restart)
	app.Post("/api/containers/:id/pause", containerActionsController.Pause)
	app.Post("/api/containers/:id/unpause", containerActionsController.UnPause)

	app.Get("/api/images", imagesController.Index)
	app.Post("/api/images", imagesController.Create)
	app.Get("/api/images/:id", imagesController.Show)
	app.Del("/api/images/:id", imagesController.Destroy)
	app.Get("/api/images/:id/history", imageActionsController.History)

	app.Get("/api/system", systemController.Info)

	app.Get("/api/apps", controllers.NewApps().Index)

	app.Listen(":8001")
}
