package docli

import "github.com/digitaloceancloud/godo"

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
