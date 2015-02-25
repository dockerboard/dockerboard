package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func CmdNotFound(c *cli.Context, command string) {
	log.Fatalf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
}