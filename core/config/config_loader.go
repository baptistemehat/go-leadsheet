package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Folder  string `yaml:"folder"`
	Storage string `yaml:"storage"`
	Script  string `yaml:"script"`
	Log     string `yaml:"log"`
}

func LoadConfiguration(configPath string) (*Configuration, error) {
	config := &Configuration{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
