package docli

import (
	"encoding/json"
	"io"
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
		SSHKey            string
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
