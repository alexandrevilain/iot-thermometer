defmodule Sensor.Worker do
  use GenServer
  require Logger

  def start_link(_opts) do
    GenServer.start_link(__MODULE__, %{}, [])
  end

  def init(state) do
    schedule_read_dht()
    {:ok, state}
  end

  def handle_info(:read_dht, state) do
    # Do the work you desire here
    {:ok, humidity, temperature} = NervesDHT.device_read(:dht_sensor)
    Logger.info("Actual humidity: #{humidity}")
    Logger.info("Actual temperature: #{temperature}")
    send_to_websocket(temperature, humidity)

    # Reschedule once more
    schedule_read_dht()
    {:noreply, state}
  end

  defp send_to_websocket(temperature, humidity) do
    pid = GenServer.whereis(Sensor.SocketClient)
    Sensor.SocketClient.send_temperature_value(pid, temperature)
    Sensor.SocketClient.send_humidity_value(pid, humidity)
  end

  defp schedule_read_dht() do
    # In 1min
    Process.send_after(self(), :read_dht, 60 * 1000)
  end
end
