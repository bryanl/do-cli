package main

import (
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
		{
			Name:      "droplets",
			ShortName: "d",
			Usage:     "manage droplets",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "create a new droplet",
					Action: ActionDropletCreate,
				},
			},
		},
	}

	app.Run(os.Args)
}

// ActionInit is the init handler
func ActionInit(c *cli.Context) {
	i := docli.Init{}

	p, err := configFile()
	if err != nil {
		log.Printf("can't create config file: %s", err)
		return
	}
	f, err := os.Create(p)
	if err != nil {
		log.Printf("can't create config file: %s", err)
		return
	}

	i.CreateConfig(c.Args()[0], f)
}

// ActionDropletCreate is the droplet create handler
func ActionDropletCreate(c *cli.Context) {
	p, err := configFile()
	if err != nil {
		log.Printf("can't create config file: %s", err)
		return
	}
	f, err := os.Open(p)
	if err != nil {
		log.Printf("can't open config file: %s", err)
		return
	}

	name := c.Args()[0]
	config, err := docli.ConfigLoad(f)
	if err != nil {
		log.Printf("can't read config: %s", err)
		return
	}

	err = docli.DropletCreate(name, config)
	if err != nil {
		log.Printf("couldn't create droplet: %s", err)
	}
}

func configFile() (string, error) {
	homeDir, err := userHomeDir()
	if err != nil {
		log.Println("can't find home directory")
		return "", err
	}

	return path.Join(homeDir, ".do-cli.json"), nil
}

// UserHomeDir returns the user's home directory
func userHomeDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}
