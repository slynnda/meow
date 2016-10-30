package main

import (
	"github.com/urfave/cli"
	commands "meow/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "meow"
	app.Usage = "make your development story easier... right meow."
	app.Commands = []cli.Command{
                commands.Init,
                commands.Config,
		commands.Docker,
                commands.Git,
	}
	app.Run(os.Args)
}
