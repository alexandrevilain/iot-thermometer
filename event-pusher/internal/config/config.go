package config

import (
	"github.com/kelseyhightower/envconfig"
)

// MqttConfig is the config stuct for mqtt client related config
type MqttConfig struct {
	ConnectionString string `default:"tcp://test.mosquitto.org:1883"`
	ClientID         string `default:"home-cloud-agent"`
}

// InfluxDBConfig is the config stuct for influxdb client related config
type InfluxDBConfig struct {
	ConnectionString string `default:"http://localhost:8086"`
	Username         string `default:"thermometer"`
	Password         string `default:"thermometer"`
}

// Config is the config struct of the forwarder agent
type Config struct {
	Mqtt     MqttConfig
	InfluxDB InfluxDBConfig
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
