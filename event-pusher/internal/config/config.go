package config

import (
	"github.com/kelseyhightower/envconfig"
)

// MqttConfig is the config stuct for mqtt client related config
type MqttConfig struct {
	ConnectionString string `default:"tcp://test.mosquitto.org:1883"`
	ClientID         string `default:"home-cloud-agent"`
	UseCredentials   bool   `default:"false"`
}

// HTTP is the config stuct for the http server
type HTTP struct {
	ListenAddr string `default:":4000"`
}

// Config is the config struct of the forwarder agent
type Config struct {
	Mqtt MqttConfig
	HTTP HTTP
}

// GetFromEnv return the Config struct populated with env variables values
func GetFromEnv() (Config, error) {
	var cfg Config
	err := envconfig.Process("agent", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
