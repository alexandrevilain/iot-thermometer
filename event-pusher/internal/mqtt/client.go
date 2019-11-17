package mqtt

import (
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createMqttOptions(connectionString, clientID string) (*mqtt.ClientOptions, error) {
	opts := mqtt.NewClientOptions()
	uri, err := url.Parse(connectionString)
	if err != nil {
		return opts, err
	}
	opts.AddBroker(fmt.Sprintf("tls://%s", uri.Host))
	// opts.SetUsername(uri.User.Username())
	// password, _ := uri.User.Password()
	// opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts, nil
}

// CreateClient instanciate a connection to an MQTT broker using a connection string and a clientID
func CreateClient(connectionString, clientID string) (mqtt.Client, error) {
	// opts, err := createMqttOptions(connectionString, clientID)
	// if err != nil {
	// 	return nil, err
	// }

	opts := mqtt.NewClientOptions().AddBroker(connectionString)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}
