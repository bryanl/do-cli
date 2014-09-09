package docli

import (
	"encoding/json"
	"fmt"

	"github.com/digitaloceancloud/godo"
)

type DropletCreateConfig struct {
	Region            string
	Image             string
	Size              string
	SSHKey            int
	PrivateNetworking bool
	BackupsEnabled    bool
}

func NewDropletCreateConfig(c *Config) *DropletCreateConfig {
	return &DropletCreateConfig{
		Region:            c.Defaults.Region,
		Image:             c.Defaults.Image,
		Size:              c.Defaults.Size,
		SSHKey:            c.Defaults.SSHKey,
		PrivateNetworking: c.Defaults.PrivateNetworking,
		BackupsEnabled:    c.Defaults.BackupsEnabled,
	}
}

func DropletCreate(name string, dcr *DropletCreateConfig, c *Config) error {
	client := c.Client()
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: c.Defaults.Region,
		Size:   c.Defaults.Size,
		Image:  c.Defaults.Image,
	}

	droplet, _, err := client.Droplets.Create(createRequest)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(droplet, "", "    ")
	fmt.Println(string(b))
	return nil
}

func DropletList(c *Config) error {
	list := []godo.Droplet{}

	client := c.Client()

	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(opt)

		if err != nil {
			return err
		}

		for _, d := range droplets {
			list = append(list, d)
		}

		if resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}

		opt.Page = page + 1
	}

	b, _ := json.MarshalIndent(list, "", "    ")
	fmt.Println(string(b))

	return nil
}

func DropletDelete(id int, c *Config) error {
	client := c.Client()
	_, err := client.Droplets.Delete(id)
	return err
}
