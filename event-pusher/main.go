package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/alexandrevilain/iot-thermometer/event-pusher/internal/config"
	mqttutil "github.com/alexandrevilain/iot-thermometer/event-pusher/internal/mqtt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var metric = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "device_metric",
	Help: "Current temperature",
}, []string{"room", "type"})

func main() {
	cfg, err := config.GetFromEnv()
	if err != nil {
		log.Fatalf("Unable to get config: %v", err.Error())
	}

	mqttClient, err := mqttutil.CreateClient(cfg.Mqtt.ConnectionString, cfg.Mqtt.ClientID, cfg.Mqtt.UseCredentials)
	if err != nil {
		log.Fatalf("Unable to connect to mqtt: %v", err)
	}

	log.Println("Connected to the MQTT broker")

	topicRegex := regexp.MustCompile(`rooms\/(\w+)\/(\w+)`)

	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Topic : %s received new message\n", msg.Topic())

		result := topicRegex.FindStringSubmatch(msg.Topic())
		roomName := result[1]
		metricType := result[2]

		payload := string(msg.Payload())
		value, err := strconv.Atoi(payload)
		if err != nil {
			log.Printf("Unable to parse metric value: %v for metric type: %v", payload, metricType)
			return
		}
		metric.With(prometheus.Labels{"room": roomName, "type": metricType}).Set(float64(value))
	}

	go func() {
		if token := mqttClient.Subscribe("rooms/+/+", 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(cfg.HTTP.ListenAddr, nil)
}
