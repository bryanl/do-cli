package docli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit_CreateConfig(t *testing.T) {
	assert := assert.New(t)

	i := &Init{}

	var b bytes.Buffer
	i.CreateConfig("12345", &b)

	assert.Contains(b.String(), `"Token": "12345"`)
}
