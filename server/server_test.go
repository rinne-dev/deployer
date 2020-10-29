package server

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCliConfig(t *testing.T) {
	cmd := CliConfig()
	assert.Equal(t, cmd.Name, "serve")
}
