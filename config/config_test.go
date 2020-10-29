package config

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// TestConfigFile 测试配置文件
const TestConfigFile = "../config.yaml"

func TestLoad(t *testing.T) {
	err := Load(TestConfigFile)
	assert.Equal(t, err, nil)
}

func TestGet(t *testing.T) {
	res := Get("tasks.test.type")
	assert.Equal(t, res.(string), "github")
}

func TestGetString(t *testing.T) {
	res := GetString("tasks.test.type")
	assert.Equal(t, res, "github")
}

func TestGetStringSlice(t *testing.T) {
	res := GetStringSlice("tasks.test.templates")
	assert.Equal(t, len(res), 1)
}
