package config

type Auth struct {
	Port int `yaml:"auth_port" env:"AUTH_PORT"`
}
