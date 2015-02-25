package app

import (
	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"
	"github.com/goocean/methodoverride"

	"github.com/dockerboard/dockerboard/app/controllers"
)

// Start Dockerboard application.
func Run(static, port string) {

	// The Instances of Controllers.
	containersController := controllers.NewContainers()
	containerActionsController := controllers.NewContainerActions()
	imagesController := controllers.NewImages()
	imageActionsController := controllers.NewImageActions()
	systemController := controllers.NewSystem()
	hostsController := controllers.NewHosts()
	hostActionsController := controllers.NewHostActions()

	// Create app.
	app := app.New()
	app.Use(WSHandler("/ws"))
	app.Use(logger.New())
	app.Use(serve.New(static))
	app.Use(methodoverride.New())
	app.Get("/api", APIHandler)

	// Controllers CRUD APIs etc.
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
	app.Get("/api/containers/:id/top", containerActionsController.Top)
	app.Post("/api/containers/:id/rename", containerActionsController.Rename)
	app.Get("/api/containers/:id/stats", containerActionsController.Stats)

	// Images CRUD APIs etc.
	app.Get("/api/images", imagesController.Index)
	app.Post("/api/images", imagesController.Create)
	app.Get("/api/images/search", imagesController.Search)
	app.Get("/api/images/:id", imagesController.Show)
	app.Del("/api/images/:id", imagesController.Destroy)
	app.Get("/api/images/:id/history", imageActionsController.History)
	app.Post("/api/images/:id/tag", imageActionsController.Tag)
	app.Post("/api/images/:name/push", imageActionsController.Push)

	// Hosts CRUD APIs etc.
	app.Get("/api/hosts", hostsController.Index)
	app.Post("/api/hosts", hostsController.Create)
	app.Del("/api/hosts/:id", hostsController.Destroy)
	app.Get("/api/hosts/:id/ping", hostActionsController.Ping)
	app.Get("/api/hosts/:id/version", hostActionsController.Version)

	app.Get("/api/system", systemController.Info)
	app.Get("/api/apps", controllers.NewApps().Index)

	// Listen Port.
	app.Listen(":" + port)
}
