package main

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/alexandrevilain/iot-thermometer/event-pusher/internal/config"
	mqttutil "github.com/alexandrevilain/iot-thermometer/event-pusher/internal/mqtt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb "github.com/influxdata/influxdb-client-go"
)

func createMessageHandler(influxClient *influxdb.Client) func(mqtt.Client, mqtt.Message) {
	re := regexp.MustCompile(`rooms\/(\w+)\/(\w+)`)
	return func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Topic : %s received new message\n", msg.Topic())

		result := re.FindStringSubmatch(msg.Topic())
		roomName := result[1]
		metricType := result[2]

		payload := string(msg.Payload())
		value, err := strconv.Atoi(payload)
		if err != nil {
			log.Printf("Unable to parse metric value: %v for metric type: %v", payload, metricType)
			return
		}

		metric := []influxdb.Metric{
			influxdb.NewRowMetric(
				map[string]interface{}{metricType: value},
				"system-metrics",
				map[string]string{"room": roomName},
				time.Now()),
		}

		_, err = influxClient.Write(context.Background(), "my-awesome-bucket", "my-very-awesome-org", metric...)
		if err != nil {
			log.Printf("Unable to send metric to influxdb: %v", err.Error())
		}
	}
}

func main() {
	cfg, err := config.GetFromEnv()
	if err != nil {
		log.Fatalf("Unable to get config: %v", err.Error())
	}

	influxClient, err := influxdb.New(cfg.InfluxDB.ConnectionString, "", influxdb.WithUserAndPass(cfg.InfluxDB.Username, cfg.InfluxDB.Password))
	if err != nil {
		log.Fatalf("Unable to connect to influxdb: %v", err)
	}

	mqttClient, err := mqttutil.CreateClient(cfg.Mqtt.ConnectionString, cfg.Mqtt.ClientID)
	if err != nil {
		log.Fatalf("Unable to connect to mqtt: %v", err)
	}

	log.Println("Connected to the MQTT broker")

	messageHandler := createMessageHandler(influxClient)

	go func() {
		if token := mqttClient.Subscribe("rooms/+/temp", 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}
	}()

	// Todo: use a done chan to handle CTRL+C
	select {}
}
