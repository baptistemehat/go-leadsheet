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

func SaveConfiguration(config *Configuration, configPath string) error {

	// Create config file
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)

	if err := encoder.Encode(&config); err != nil {
		return err
	}

	return nil
}

func DefaultConfiguration() (*Configuration, error) {
	config := &Configuration{
		Folder:  "",
		Storage: "",
		Script:  "",
		Log:     "",
	}

	return config, nil
}
