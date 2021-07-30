package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Parsed config
type Config struct {
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"mysql"`

	WebServer struct {
		AddressAndPort string `yaml:"address_and_port"`
		APIAccessToken string `yaml:"api_access_token"`
	} `yaml:"webserver"`
}

// Checks whether or not the config file exists
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("Unable to find config.yml. (Path: '%s')", path)
	}
	return nil
}

// Parses the config file
func ParseConfig(path string) (*Config, error) {
	if err := validateConfigPath(path); err != nil {
		return nil, err
	}

	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)

	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	log.Println(">> Config has been parsed successfully!")
	return config, nil
}
