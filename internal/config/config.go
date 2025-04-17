package config

type Config struct {
	Port         string
	DatabaseName string
	Dbpath       string
}

func Load() *Config {
	return &Config{}
}
