package config

type nasa struct {
	ApiKey     string `yaml:"api_key" env:"NASA_API_KEY"`
	BaseURL    string `yaml:"base_url" env:"NASA_BASE_URL"`
	MaxRetries int    `yaml:"max_retries" env:"NASA_MAX_RETRIES"`
}
