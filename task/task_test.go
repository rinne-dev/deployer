package task

import (
	"github.com/magiconair/properties/assert"
	"rinne.dev/deployer/config"
	"testing"
)

func init() {
	config.Load("../config.yaml")
}

func TestExist(t *testing.T) {
	err := Exist("test")
	assert.Equal(t, err, nil)
}

func TestType(t *testing.T) {
	s := Type("test")
	assert.Equal(t, s, "github")
}

func TestSecret(t *testing.T) {
	s := Secret("test")
	assert.Equal(t, s, "xxxxxxxx")
}

func TestWorkdir(t *testing.T) {
	s := Workdir("test")
	assert.Equal(t, s, "./")
}

func TestTemplates(t *testing.T) {
	l := Templates("test")
	assert.Equal(t, len(l), 1)
}

func TestCommands(t *testing.T) {
	l := Commands("test")
	assert.Equal(t, len(l), 1)
}
