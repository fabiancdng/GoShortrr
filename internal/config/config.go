package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

// Holds data of parsed config
type Config struct {
	MySQL struct {
		Host     string `yaml:"host" env:"GOSHORTRR_MYSQL_HOST,notEmpty"`
		Port     int    `yaml:"port" env:"GOSHORTRR_MYSQL_PORT,notEmpty"`
		User     string `yaml:"user" env:"GOSHORTRR_MYSQL_USER,notEmpty"`
		Password string `yaml:"password" env:"GOSHORTRR_MYSQL_PASSWORD,notEmpty"`
		DB       string `yaml:"db" env:"GOSHORTRR_MYSQL_DB,notEmpty"`
	} `yaml:"mysql"`

	WebServer struct {
		AddressAndPort string `yaml:"address_and_port" env:"GOSHORTRR_ADDRESS_AND_PORT,notEmpty"`
		APIAccessToken string `yaml:"api_access_token" env:"GOSHORTRR_API_ACCESS_TOKEN"`
	} `yaml:"webserver"`
}

// Checks whether or not the config file exists
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("unable to find config.yml. (Path: '%s')", path)
	}

	return nil
}

// Parses the config file
func parseConfigFile(path string) (*Config, error) {
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

// Parses the environment variables as a config
// Also checks for required fields and empty fields
func parseEnvConfig() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Parses the config file
func ParseConfig(path string) (*Config, error) {
	// Config file doesn't exist
	if err := validateConfigPath(path); err != nil {
		// Try to parse environment variables
		config, err := parseEnvConfig()
		if err != nil {
			return nil, err
		}

		return config, nil
	}

	// Config file exists
	// Read and parse config file
	config, err := parseConfigFile(path)
	if err != nil {
		return nil, err
	}

	return config, nil
}
