package commands

import (
  "github.com/urfave/cli"
  "bufio"
  "os"
  "fmt"
)

var Init = cli.Command {
  Name: "init",
  Usage: "initialize project",
  Action: func(c *cli.Context) error {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Welcome to the interactive meow project initializer! Follow the prompts to create your project!\n")
    fmt.Print("Project name: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
    return nil
  },
}
