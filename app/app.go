package app

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"
	"github.com/goocean/methodoverride"

	"github.com/dockerboard/dockerboard/app/controllers"
)

func APIIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Coming soon!\n")
}

// Start Dockerboard application.
func Serve() {

	// Set bluewhale dir form ENV BLUEWHALE_DIST or /bluewhale/dist.
	bluewhale := os.Getenv("BLUEWHALE_DIST")
	if bluewhale == "" {
		bluewhale = "/bluewhale/dist/"
	}
	dir := path.Join(path.Dir(bluewhale))

	// The Instances of Controllers.
	containersController := controllers.NewContainers()
	containerActionsController := controllers.NewContainerActions()
	imagesController := controllers.NewImages()
	imageActionsController := controllers.NewImageActions()
	systemController := controllers.NewSystem()
	hostsController := controllers.NewHosts()

	// Create app.
	app := app.New()
	app.Use(logger.New())
	app.Use(serve.New(dir))
	app.Use(methodoverride.New())
	app.Get("/api", APIIndex)

	// Controllers CURD APIs etc.
	app.Get("/api/containers", containersController.Index)
	app.Post("/api/containers", containersController.Create)
	app.Get("/api/containers/:id", containersController.Show)
	app.Del("/api/containers/:id", containersController.Destroy)
	app.Post("/api/containers/:id/start", containerActionsController.Start)
	app.Post("/api/containers/:id/stop", containerActionsController.Stop)
	app.Post("/api/containers/:id/restart", containerActionsController.Restart)
	app.Post("/api/containers/:id/pause", containerActionsController.Pause)
	app.Post("/api/containers/:id/unpause", containerActionsController.UnPause)
	app.Post("/api/containers/:id/kill", containerActionsController.Kill)
	app.Get("/api/containers/:id/logs", containerActionsController.Logs)

	// Images CURD APIs etc.
	app.Get("/api/images", imagesController.Index)
	app.Post("/api/images", imagesController.Create)
	app.Get("/api/images/:id", imagesController.Show)
	app.Del("/api/images/:id", imagesController.Destroy)
	app.Get("/api/images/:id/history", imageActionsController.History)

	// Hosts CURD APIs etc.
	app.Get("/api/hosts", hostsController.Index)
	app.Post("/api/hosts", hostsController.Create)
	app.Del("/api/hosts/:id", hostsController.Destroy)
	//app.Del("/api/hosts/:id/ping", hostActionsController.Destroy)

	app.Get("/api/system", systemController.Info)
	app.Get("/api/apps", controllers.NewApps().Index)

	// Listen Port.
	app.Listen(":8001")
}
