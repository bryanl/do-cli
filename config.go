package docli

import (
	"encoding/json"
	"io"

	"github.com/digitaloceancloud/godo"

	"code.google.com/p/goauth2/oauth"
)

// Config holds configuration directives
type Config struct {
	Auth struct {
		Token string
	}

	Defaults struct {
		Region            string
		Image             string
		Size              string
		SSHKey            int
		PrivateNetworking bool
		BackupsEnabled    bool
	}
}

func ConfigLoad(r io.Reader) (*Config, error) {
	var config Config
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

func (c *Config) Save(w io.Writer) error {
	b, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func (c *Config) Client() *godo.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: c.Auth.Token},
	}

	return godo.NewClient(t.Client())
}
