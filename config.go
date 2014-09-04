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
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}
