package docli

import "io"

// Init manages the initialization actions for go-cli
type Init struct {
}

// CreateConfig creates a new configuration for go-cli
func (i *Init) CreateConfig(token string, w io.Writer) {
	config := &Config{}
	config.Auth.Token = token
	config.Save(w)
}
