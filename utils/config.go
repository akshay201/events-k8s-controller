package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	EnableLeaderElection bool   `mapstructure:"LEADER_ELECT"`
	MetricsAddr          string `mapstructure:"METRICS_BIND_ADDRESS"`
	ProbeAddr            string `mapstructure:"HEALTH_PROBE_BIND_ADDRESS"`
	LokiGatewayUrl       string `mapstructure:"LOKI_GATEWAY_URL"`
	LogLevel             string `mapstructure:"LOG_LEVEL"`
	LokiLabels           string `mapstructure:"LOKI_LABELS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	log.Print(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
