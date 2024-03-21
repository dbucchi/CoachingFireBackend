package utility

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Configuration Config

type Config struct {
	Database      DatabaseConfig `yaml:"database"`
	Server        ServerConfig   `yaml:"server"`
	OtherSettings OtherSettings  `yaml:"otherSettings"`
}

// DatabaseConfig contiene le informazioni del database
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// ServerConfig contiene le impostazioni del server
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// OtherSettings contiene altre impostazioni
type OtherSettings struct {
	DebugMode bool   `yaml:"debugMode"`
	LogLevel  string `yaml:"logLevel"`
}

func NewConfiguration() *Config {
	config := &Config{}
	config.read()
	return config
}

func (config *Config) read() {
	rawConfig, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Errore nella lettura del file di configurazione: %v", err)
	}

	// Parsa il file di configurazione YAML
	err = yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		log.Fatalf("Errore nel parsing del file YAML: %v", err)
	}
}
