package config

type grpc struct {
	Port int `yaml:"port" env:"GRPC_PORT"`
}
