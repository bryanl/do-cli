package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/bryanl/do-cli"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "do-cli"
	app.Usage = "Digital Ocean cli"

	app.Commands = []cli.Command{
		{
			Name:      "init",
			ShortName: "i",
			Usage:     "initialize configuration",
			Action:    ActionInit,
		},
	}

	app.Run(os.Args)
}

// ActionInit is the init handler
func ActionInit(c *cli.Context) {
	fmt.Println("initializing", c.Args())

	i := docli.Init{}

	homeDir, err := UserHomeDir()
	if err != nil {
		log.Println("can't find home directory")
		return
	}

	p := path.Join(homeDir, ".do-cli.json")
	f, err := os.Create(p)
	if err != nil {
		log.Println("can't create config file")
		return
	}

	i.CreateConfig(c.Args()[0], f)
}

// UserHomeDir returns the user's home directory
func UserHomeDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}
