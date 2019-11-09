defmodule Metrics.TemperatureInstrumenter do
  use Prometheus.Metric

  def setup() do
    Gauge.declare(
      name: :temperature,
      help: "Current temperature",
      labels: [:device]
    )
  end

  def set_temperature(device, value) do
    Gauge.set([name: :temperature, labels: [String.to_atom(device)]], value)
  end
end
