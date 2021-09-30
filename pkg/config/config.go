package config

type Config struct {
	Name           string
	Port           string
	WithImprobable bool
	WithTLS        bool
}

var Conf = &Config{}

func GetConfig() *Config {
	return Conf
}
