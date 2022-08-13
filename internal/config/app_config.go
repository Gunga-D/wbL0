package config

type AppConfig struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  string `mapstructure:"readtimeout"`
	WriteTimeout string `mapstructure:"writetimeout"`
}
