package docli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_WriteConfig(t *testing.T) {
	assert := assert.New(t)

	var b bytes.Buffer

	config := &Config{}
	err := config.Save(&b)
	assert.NoError(err)
	assert.NotEmpty(b.String())
}

func TestConfig_ReadConfig(t *testing.T) {
	assert := assert.New(t)

	j := `{"Auth":{"Token":"12345"}}`
	r := strings.NewReader(j)

	config, err := ConfigLoad(r)
	assert.NoError(err)

	assert.Equal("12345", config.Auth.Token)
}
