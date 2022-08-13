package config

import "github.com/spf13/viper"

func InitConfig(path, name string) (*viper.Viper, error) {
	instance := viper.New()

	instance.AddConfigPath(path)
	instance.SetConfigName(name)
	instance.SetConfigType("yaml")

	instance.AutomaticEnv()
	err := instance.ReadInConfig()

	return instance, err
}

func LoadAppConfig(path string, name string) (*AppConfig, error) {
	instance, err := InitConfig(path, name)
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{}
	err = instance.Unmarshal(&appConfig)
	return &appConfig, err
}

func LoadDBConfig(path string, name string) (*DBConfig, error) {
	instance, err := InitConfig(path, name)
	if err != nil {
		return nil, err
	}

	databaseConfig := DBConfig{}
	err = instance.Unmarshal(&databaseConfig)
	return &databaseConfig, err
}

func LoadNATStreamingConfig(path string, name string) (*NATStreamingConfig, error) {
	instance, err := InitConfig(path, name)
	if err != nil {
		return nil, err
	}

	natStreamingConfig := NATStreamingConfig{}
	err = instance.Unmarshal(&natStreamingConfig)
	return &natStreamingConfig, err
}
