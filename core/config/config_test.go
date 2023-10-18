package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {

	// TODO write config file here
	// and defer deletion

	// TODO test that configuration is properly loaded with the proper values

	path := "testResources/config.yaml"
	config, err := LoadConfiguration(path)

	assert.NoError(t, err, "should not return error")

	fmt.Println(config)
}

func TestLoadConfiguration_UnknownFile(t *testing.T) {

	path := "illegal.path"
	config, err := LoadConfiguration(path)

	assert.Error(t, err, "should return an error")
	assert.Nil(t, config, "should be nil")

}
