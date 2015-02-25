package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/dockerboard/dockerboard/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "DockerBoard"
	app.Usage = "Simple dashboards, visualizations, managements for your dockers."
	app.Author = ""
	app.Email = ""
	app.Version = cmd.VERSION
	app.Commands = []cli.Command{
		cmd.CmdServer,
	}
	app.CommandNotFound = cmd.CmdNotFound
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "Enable debug mode",
			EnvVar: "DEBUG",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			os.Setenv("DEBUG", "1")
			initLogging(log.DebugLevel)
			log.Info("Enable debugging")
		}
		return nil
	}
	app.Run(os.Args)
}

func initLogging(lvl log.Level) {
	log.SetOutput(os.Stderr)
	log.SetLevel(lvl)
}