package main

import (
	"os"

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
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
