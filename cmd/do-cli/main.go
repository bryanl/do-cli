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

var (
	commands = []cli.Command{
		{
			Name:      "init",
			ShortName: "i",
			Usage:     "initialize configuration",
			Action:    ActionInit,
		},
		{
			Name:        "droplets",
			ShortName:   "d",
			Usage:       "manage droplets",
			Subcommands: dropletCommands,
		},
		{
			Name:        "keys",
			ShortName:   "k",
			Usage:       "manage keys",
			Subcommands: keyCommands,
		},
		{
			Name:        "domains",
			ShortName:   "do",
			Usage:       "manage domains",
			Subcommands: domainCommands,
		},
	}

	dropletCommands = []cli.Command{
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
	}

	keyCommands = []cli.Command{
		{
			Name:   "list",
			Usage:  "list keys",
			Action: ActionKeyList,
		},
		{
			Name:   "delete",
			Usage:  "deletes a key by id or fingerprint",
			Action: ActionKeyDelete,
		},
	}

	domainCommands = []cli.Command{
		{
			Name:   "list",
			Usage:  "list domains",
			Action: ActionDomainList,
		},
		{
			Name:   "get",
			Usage:  "get a domain by name",
			Action: ActionDomainGet,
		},
		{
			Name:   "create",
			Usage:  "create a domain",
			Action: ActionDomainCreate,
		},
		{
			Name:   "delete",
			Usage:  "delete a domain",
			Action: ActionDomainDelete,
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "do-cli"
	app.Usage = "Digital Ocean CLI"
	app.Commands = commands
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

func checkError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		os.Exit(1)
	}
}

// ActionDropletCreate is the droplet create handler
func ActionDropletCreate(c *cli.Context) {
	name := c.Args()[0]
	config := loadConfig()
	err := docli.DropletCreate(name, config)
	checkError("couldn't create droplet", err)
}

func ActionDropletList(c *cli.Context) {
	config := loadConfig()
	err := docli.DropletList(config)
	checkError("couldn't list droplets", err)
}

func ActionDropletDelete(c *cli.Context) {
	config := loadConfig()
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}
	id, err := strconv.Atoi(c.Args()[0])
	checkError("couldn't delete droplet", err)
	err = docli.DropletDelete(id, config)
	checkError("couldn't delete droplet", err)

	fmt.Println("deleted", id)
}

func ActionKeyList(c *cli.Context) {
	config := loadConfig()

	err := docli.KeyList(config)
	if err != nil {
	}
}

func ActionKeyDelete(c *cli.Context) {
	config := loadConfig()
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}

	id := c.Args()[0]
	err := docli.KeyDelete(id, config)
	checkError("couldn't delete key", err)
}

func ActionDomainList(c *cli.Context) {
	config := loadConfig()
	err := docli.DomainList(config)
	checkError("couldn'tlist domains", err)
}

func ActionDomainGet(c *cli.Context) {
	config := loadConfig()
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}

	name := c.Args()[0]
	err := docli.DomainGet(name, config)
	checkError("couldn't fetch domain", err)
}

func ActionDomainCreate(c *cli.Context) {
	config := loadConfig()
	if len(c.Args()) != 2 {
		cli.ShowSubcommandHelp(c)
		return
	}

	name := c.Args()[0]
	ip := c.Args()[1]

	err := docli.DomainCreate(name, ip, config)
	checkError("couldn't create domain", err)
}

func ActionDomainDelete(c *cli.Context) {
	config := loadConfig()
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}
	name := c.Args()[0]
	err := docli.DomainDelete(name, config)
	checkError("couldn't delete domain", err)

	fmt.Println("deleted", name)
}

func loadConfig() *docli.Config {
	p, err := configFile()
	checkError("couldn't load config", err)
	f, err := os.Open(p)
	checkError("couldn't load config", err)
	config, err := docli.ConfigLoad(f)
	checkError("couldn't load config", err)

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

func userHomeDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}
