package config

type DBConfig struct {
	Database string `mapstructure:"POSTGRES_DB"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
}
