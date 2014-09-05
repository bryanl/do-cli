package docli

import (
	"encoding/json"
	"fmt"
)

func KeyList(c *Config) error {
	client := c.Client()
	keys, _, err := client.Keys.List()
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(keys, "", "    ")
	fmt.Println(string(b))

	return nil

}
