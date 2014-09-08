package docli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/digitaloceancloud/godo"
)

func KeyList(c *Config) error {
	list := []godo.Key{}

	client := c.Client()

	opt := &godo.ListOptions{}
	for {
		keys, resp, err := client.Keys.List(opt)

		if err != nil {
			return err
		}

		for _, k := range keys {
			list = append(list, k)
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

func KeyDelete(id interface{}, c *Config) error {
	client := c.Client()
	idStr := id.(string)
	if idInt, err := strconv.Atoi(idStr); err == nil {
		_, err := client.Keys.DeleteByID(idInt)
		return err
	}
	_, err := client.Keys.DeleteByFingerprint(idStr)
	return err
}
