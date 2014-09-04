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

	droplet, _, err := client.Droplet.Create(createRequest)
	if err != nil {
		return err
	}

	fmt.Println(droplet.Droplet.ID)
	return nil
}

func DropletList(c *Config) error {
	client := c.Client()
	droplets, _, err := client.Droplet.List()
	if err != nil {
		return err
	}

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

func DropletDelete(id int, c *Config) error {
	client := c.Client()
	_, err := client.Droplet.Delete(id)
	return err
}
