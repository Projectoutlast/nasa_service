package config

type Clients struct {
	Nasa string `yaml:"nasa_addr" env:"NASA_ADDR"`
	Auth string `yaml:"auth_addr" env:"AUTH_ADDR"`
}