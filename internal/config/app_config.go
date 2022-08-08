package config

type AppConfig struct {
	Port         string `mapstructure:"PORT"`
	ReadTimeout  string `mapstructure:"READ_TIMEOUT"`
	WriteTimeout string `mapstructure:"WRITE_TIMEOUT"`
}
