package mqtt

import (
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createMqttOptions(connectionString, clientID string, useCredentials bool) (*mqtt.ClientOptions, error) {
	opts := mqtt.NewClientOptions()
	uri, err := url.Parse(connectionString)
	if err != nil {
		return opts, err
	}
	opts.AddBroker(fmt.Sprintf("tls://%s", uri.Host))
	if useCredentials {
		opts.SetUsername(uri.User.Username())
		password, _ := uri.User.Password()
		opts.SetPassword(password)
	}
	opts.SetClientID(clientID)
	return opts, nil
}

// CreateClient instanciate a connection to an MQTT broker using a connection string and a clientID
func CreateClient(connectionString, clientID string, useCredentials bool) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(connectionString)
	// opts, err := createMqttOptions(connectionString, clientID, useCredentials)
	// if err != nil {
	// 	return nil, err
	// }
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}
