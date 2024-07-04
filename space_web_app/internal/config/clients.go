package config

type Clients struct {
	Nasa string `yaml:"nasa" env:"NASA_ADDR"`
	Auth string `yaml:"auth" env:"AUTH_ADDR"`
}