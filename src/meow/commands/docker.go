package commands

import (
	"github.com/urfave/cli"
)

var Docker = cli.Command{
	Name:    "docker",
	Usage:   "make docker do stuff in the context of your development environment",
	Action: func(c *cli.Context) error {
		return nil
	},
}
