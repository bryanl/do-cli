package docli

import (
	"fmt"

	"github.com/digitaloceancloud/godo"
)

func DropletCreate(name string, c *Config) error {
	client := c.Client()
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: c.Defaults.Region,
		Size:   c.Defaults.Size,
		Image:  c.Defaults.Image,
	}

	_, _, err := client.Droplet.Create(createRequest)

	return err
}

func DropletList(c *Config) error {
	client := c.Client()
	droplets, _, err := client.Droplet.List()
	if err != nil {
		return err
	}

	// (ip: 107.170.118.88, status: active, region: 4, id: 1400861)
	for _, d := range droplets {
		fmt.Printf(
			"%s (ip: %s, status: %s, region: %s, id: %d)\n",
			d.Name,
			d.Networks.V4[0].IPAddress,
			d.Status,
			d.Region.Slug,
			d.ID,
		)

	}

	return nil
}
