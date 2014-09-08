package docli

import (
	"encoding/json"
	"fmt"

	"github.com/digitaloceancloud/godo"
)

func DomainGet(name string, c *Config) error {
	client := c.Client()
	domain, _, err := client.Domains.Get(name)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(domain, "", "    ")
	fmt.Println(string(b))

	return nil
}

func DomainDelete(name string, c *Config) error {
	client := c.Client()
	_, err := client.Domains.Delete(name)
	if err != nil {
		return err
	}

	return nil
}

func DomainCreate(name, ip string, c *Config) error {
	client := c.Client()
	cr := &godo.DomainCreateRequest{
		Name:      name,
		IPAddress: ip,
	}
	domain, _, err := client.Domains.Create(cr)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(domain, "", "    ")
	fmt.Println(string(b))

	return nil
}

func DomainList(c *Config) error {
	list := []godo.Domain{}

	client := c.Client()

	opt := &godo.ListOptions{}
	for {
		domains, resp, err := client.Domains.List(opt)

		if err != nil {
			return err
		}

		for _, d := range domains {
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
