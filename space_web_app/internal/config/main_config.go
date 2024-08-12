package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment    string       `yaml:"environment" env:"ENVIRONMENT" default:"development"`
	PubKeyPath     string       `yaml:"pubkey_path" env:"PUBKEY_PATH"`
	Server         ServerConfig `yaml:"http_server_config"`
	ClientsAddress Clients      `yaml:"clients_address"`
	CertFile       string       `yaml:"cert_file" env:"CERT_FILE"`
	KeyFile        string       `yaml:"key_file" env:"KEY_FILE"`
}

func MustLoad() (*Config, error) {
	configPath := fetchConfigPath()

	if configPath == "" {
		cfg, err := getConfigDataFromEnvVar()
		if err != nil {
			return nil, err
		}

		return cfg, nil
	}

	cfg, err := getConfigDataFromFile(&configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}

func fetchConfigPath() string {
	res := flag.String("config", "", "путь к файлу конфигурации")
	
	flag.Parse()

	return *res
}

func getConfigDataFromEnvVar() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigDataFromFile(configPath *string) (*Config, error) {
	var cfg *Config

	file, err := os.Open(*configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
