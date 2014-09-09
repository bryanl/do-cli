package docli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/digitaloceancloud/godo"
)

func ImageList(c *Config) error {
	list := []godo.Image{}

	client := c.Client()

	opt := &godo.ListOptions{}
	for {
		images, resp, err := client.Images.List(opt)

		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %s\n", err)
			return err
		}

		for _, i := range images {
			list = append(list, i)
		}

		if resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %s\n", err)
			return err
		}

		opt.Page = page + 1
	}

	b, _ := json.MarshalIndent(list, "", "    ")
	fmt.Println(string(b))
	fmt.Fprintf(os.Stderr, "%d\n", len(list))

	return nil
}
