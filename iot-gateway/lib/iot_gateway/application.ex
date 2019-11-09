defmodule IotGateway.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    # List all child processes to be supervised
    children = [
      # Start the endpoint when the application starts
      IotGatewayWeb.Endpoint
      # Starts a worker by calling: IotGateway.Worker.start_link(arg)
      # {IotGateway.Worker, arg},
    ]

    Metrics.Setup.setup()

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: IotGateway.Supervisor]
    Supervisor.start_link(children, opts)
  end

  def application do
    [applications: [:prometheus_ex]]
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  def config_change(changed, _new, removed) do
    IotGatewayWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
