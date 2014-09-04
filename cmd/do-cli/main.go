package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"

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
					Usage:  "create a new droplet given a name",
					Action: ActionDropletCreate,
				},
				{
					Name:   "list",
					Usage:  "list droplets",
					Action: ActionDropletList,
				},
				{
					Name:   "delete",
					Usage:  "deletes a droplet by id",
					Action: ActionDropletDelete,
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

func actionErr(msg string) {
	log.Println(msg)
	os.Exit(1)
}

// ActionDropletCreate is the droplet create handler
func ActionDropletCreate(c *cli.Context) {
	name := c.Args()[0]
	config := loadConfig()
	err := docli.DropletCreate(name, config)
	if err != nil {
		actionErr(fmt.Sprintf("couldn't create droplet: %s", err))
	}
}

func ActionDropletList(c *cli.Context) {
	config := loadConfig()
	err := docli.DropletList(config)
	if err != nil {
		actionErr(fmt.Sprintf("couldn't list droplets: %s", err))
	}
}

func ActionDropletDelete(c *cli.Context) {
	config := loadConfig()
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}
	id, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		actionErr(fmt.Sprintf("couldn't delete droplet: %s", err))
	}

	err = docli.DropletDelete(id, config)
	if err != nil {
		fmt.Printf("couldn't delete: %s", err)
		return
	}

	fmt.Println("deleted", id)
}

func loadConfig() *docli.Config {
	p, err := configFile()
	if err != nil {
		actionErr(fmt.Sprintf("couldn't load config", err))
	}
	f, err := os.Open(p)
	if err != nil {
		actionErr(fmt.Sprintf("couldn't load config", err))
	}

	config, err := docli.ConfigLoad(f)
	if err != nil {
		actionErr(fmt.Sprintf("couldn't load config", err))
	}

	return config
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
