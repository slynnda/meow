package commands

import (
  "github.com/urfave/cli"
)

var Config = cli.Command {
  Name: "config",
  Usage: "configure project",
  Action: func(c *cli.Context) error {
    return nil
  },
}
