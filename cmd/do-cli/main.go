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

var ()

func main() {
	app := cli.NewApp()
	app.Name = "do-cli"
	app.Usage = "Digital Ocean CLI"
	app.Version = "0.1.0"
	app.Commands = buildCommands()
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
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// ActionDropletCreate is the droplet create handler
func ActionDropletCreate(c *cli.Context) {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}
	name := c.Args()[0]

	config, err := loadConfig()
	checkError("foo", err)
	dcr := docli.NewDropletCreateConfig(config)

	if f := c.String("region"); f != "" {
		dcr.Region = f
	}

	if f := c.String("image"); f != "" {
		dcr.Image = f
	}

	if f := c.String("size"); f != "" {
		dcr.Size = f
	}

	if f := c.Int("ssh-key"); f != 0 {
		dcr.SSHKey = f
	}

	err = docli.DropletCreate(name, dcr, config)
	checkError("couldn't create droplet", err)
}

func ActionDropletList(c *cli.Context) {
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.DropletList(config)
	checkError("couldn't list droplets", err)
}

func ActionDropletDelete(c *cli.Context) {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}
	id, err := strconv.Atoi(c.Args()[0])
	checkError("couldn't delete droplet", err)
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.DropletDelete(id, config)
	checkError("couldn't delete droplet", err)

	fmt.Println("deleted", id)
}

func ActionKeyList(c *cli.Context) {
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.KeyList(config)
	if err != nil {
	}
}

func ActionKeyDelete(c *cli.Context) {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}

	id := c.Args()[0]
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.KeyDelete(id, config)
	checkError("couldn't delete key", err)
}

func actionImageList(c *cli.Context) {
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.ImageList(config)
	if err != nil {
	}
}

func ActionDomainList(c *cli.Context) {
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.DomainList(config)
	checkError("couldn'tlist domains", err)
}

func ActionDomainGet(c *cli.Context) {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}

	name := c.Args()[0]
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.DomainGet(name, config)
	checkError("couldn't fetch domain", err)
}

func ActionDomainCreate(c *cli.Context) {
	if len(c.Args()) != 2 {
		cli.ShowSubcommandHelp(c)
		return
	}

	name := c.Args()[0]
	ip := c.Args()[1]

	config, err := loadConfig()
	checkError("foo", err)
	err = docli.DomainCreate(name, ip, config)
	checkError("couldn't create domain", err)
}

func ActionDomainDelete(c *cli.Context) {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return
	}
	name := c.Args()[0]
	config, err := loadConfig()
	checkError("foo", err)
	err = docli.DomainDelete(name, config)
	checkError("couldn't delete domain", err)

	fmt.Println("deleted", name)
}

func loadConfig() (*docli.Config, error) {
	p, err := configFile()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	config, err := docli.ConfigLoad(f)
	if err != nil {
		return nil, err
	}

	return config, nil
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

func buildCommands() []cli.Command {
	return []cli.Command{
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
			Subcommands: buildDropletCommands(),
		},
		{
			Name:        "keys",
			ShortName:   "k",
			Usage:       "manage keys",
			Subcommands: buildKeyCommands(),
		},
		{
			Name:        "images",
			ShortName:   "i",
			Usage:       "manage imagesj",
			Subcommands: buildImageCommands(),
		},
		{
			Name:        "domains",
			ShortName:   "do",
			Usage:       "manage domains",
			Subcommands: buildDomainCommands(),
		},
	}

}

func buildKeyCommands() []cli.Command {
	return []cli.Command{
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
}

func buildImageCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "list images",
			Action: actionImageList,
		},
	}
}

func buildDomainCommands() []cli.Command {
	return []cli.Command{
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
}

func buildDropletCommands() []cli.Command {
	config, err := loadConfig()
	if err != nil {
		return []cli.Command{}
	}
	dropletCreateFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "region",
			Value: config.Defaults.Region,
			Usage: "Region",
		},
		cli.StringFlag{
			Name:  "image",
			Value: config.Defaults.Image,
			Usage: "Image",
		},
		cli.StringFlag{
			Name:  "size",
			Value: config.Defaults.Size,
			Usage: "Size",
		},
		cli.IntFlag{
			Name:  "ssh-key",
			Value: config.Defaults.SSHKey,
			Usage: "SSH Key",
		},
	}

	if config.Defaults.PrivateNetworking {
		f := cli.BoolTFlag{
			Name:  "private-networking",
			Usage: "Private networking enabled",
		}
		dropletCreateFlags = append(dropletCreateFlags, f)
	} else {
		f := cli.BoolTFlag{
			Name:  "private-networking",
			Usage: "Private networking disabled",
		}
		dropletCreateFlags = append(dropletCreateFlags, f)
	}

	if config.Defaults.BackupsEnabled {
		f := cli.BoolTFlag{
			Name:  "backups",
			Usage: "Backups enabled",
		}
		dropletCreateFlags = append(dropletCreateFlags, f)
	} else {
		f := cli.BoolTFlag{
			Name:  "backups",
			Usage: "Backups disabled",
		}
		dropletCreateFlags = append(dropletCreateFlags, f)
	}

	return []cli.Command{
		{
			Name:        "create",
			Usage:       "create a new droplet given a name",
			Description: "create a new droplet given a name",
			Action:      ActionDropletCreate,
			Flags:       dropletCreateFlags,
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
}
