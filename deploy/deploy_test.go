package deploy

import (
	"github.com/magiconair/properties/assert"
	"os"
	"rinne.dev/deployer/config"
	"testing"
)

func loadConfig() {
	config.Load("../config.yaml")
}

func TestCliConfig(t *testing.T) {
	cmd := CliConfig()
	assert.Equal(t, cmd.Name, "deploy")
}

func TestDeploy(t *testing.T) {
	loadConfig()
	err := Deploy("test")
	assert.Equal(t, err, nil)
	_, err = os.Stat("./deploy_test")
	assert.Equal(t, err, nil)
	_ = os.Remove("./deploy_test")
}
