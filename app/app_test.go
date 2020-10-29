package app

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestApp(t *testing.T) {
	app := App()
	assert.Equal(t, app.Name, "Deployer@RinNe.Dev")
}
