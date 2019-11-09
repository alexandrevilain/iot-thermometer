defmodule IotGatewayWeb.Router do
  use IotGatewayWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", IotGatewayWeb do
    pipe_through :api
  end
end
