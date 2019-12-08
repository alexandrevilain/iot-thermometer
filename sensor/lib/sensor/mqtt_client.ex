defmodule Sensor.MqttClient do
  use Application

  def get_config do
    [
      client_id: "device_chambre",
      server: {Tortoise.Transport.Tcp, host: 'test.mosquitto.org', port: 1883},
      handler: {Tortoise.Handler.Logger, []}
    ]
  end

  def get_device_name() do
  end

  def publish_temperature(temp) do
    Tortoise.publish("device_chambre", "rooms/chambre/temparure", temp, qos: 0)
  end
end
