defmodule Metrics.HumidityInstrumenter do
  use Prometheus.Metric

  def setup() do
    Gauge.declare(
      name: :humidity,
      help: "Current humidity",
      labels: [:device]
    )
  end

  def set_humidity(device, value) do
    Gauge.set([name: :humidity, labels: [String.to_atom(device)]], value)
  end
end
