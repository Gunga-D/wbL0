package config

type NATStreamingConfig struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	ClusterID string `mapstructure:"clusterid"`
	ClientID  string `mapstructure:"clientid"`
}
