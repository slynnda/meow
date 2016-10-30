package commands

import (
  "github.com/urfave/cli"
)

var Git = cli.Command {
  Name: "git",
  Usage: "make git do stuff in the context of your develop environment",
  Action: func(c *cli.Context) error {
   return nil
  },
}
