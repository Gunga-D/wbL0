package config

import "github.com/spf13/viper"

func InitConfig(path, name string) (*viper.Viper, error) {
	instance := viper.New()

	instance.AddConfigPath(path)
	instance.SetConfigName(name)
	instance.SetConfigType("env")

	instance.AutomaticEnv()
	err := instance.ReadInConfig()

	return instance, err
}

func LoadAppConfig(path string) (*AppConfig, error) {
	instance, err := InitConfig(path, "app")
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{}
	err = instance.Unmarshal(&appConfig)
	return &appConfig, err
}

func LoadDBConfig(path string) (*DBConfig, error) {
	instance, err := InitConfig(path, "database")
	if err != nil {
		return nil, err
	}

	databaseConfig := DBConfig{}
	err = instance.Unmarshal(&databaseConfig)
	return &databaseConfig, err
}
