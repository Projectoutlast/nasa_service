package config

import "time"

type ServerConfig struct {
	Host           string        `yaml:"http_host" env:"HTTP_HOST" default:"localhost"`
	Port           int           `yaml:"http_port" env:"HTTP_PORT" default:"8080"`
	ReadTimeout    time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT"`
	WriteTimeout   time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT"`
	MaxHeaderBytes int           `yaml:"max_header_bytes" env:"MAX_HEADER_BYTES"`
}
