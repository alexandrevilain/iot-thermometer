defmodule IotGatewayWeb.DeviceChannel do
  use Phoenix.Channel
  use Prometheus.Metric

  def join("device:" <> device_id, _message, socket) do
    IO.puts(device_id)
    {:ok, assign(socket, :device_id, device_id)}
  end

  def handle_in("temperature", %{"value" => value}, socket) do
    Metrics.TemperatureInstrumenter.set_temperature(socket.assigns.device_id, value)
    {:noreply, socket}
  end

  def handle_in("humidity", %{"value" => value}, socket) do
    Metrics.HumidityInstrumenter.set_humidity(socket.assigns.device_id, value)
    {:noreply, socket}
  end
end
