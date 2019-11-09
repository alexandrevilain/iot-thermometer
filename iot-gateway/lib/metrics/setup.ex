defmodule Metrics.Setup do
  @moduledoc """
  Common area to set up metrics
  """

  def setup() do
    Metrics.TemperatureInstrumenter.setup()
    Metrics.HumidityInstrumenter.setup()

    IotGatewayWeb.PrometheusExporter.setup()
  end
end
