defmodule Sensor.SocketClient do
  use GenServer
  require Logger

  # Client

  def send_temperature_value(pid, value) do
    GenServer.cast(pid, {:send_temperature_value, value})
  end

  def send_humidity_value(pid, value) do
    GenServer.cast(pid, {:send_humidity_value, value})
  end

  # Server

  def start_link(opts) do
    GenServer.start_link(__MODULE__, opts, name: __MODULE__)
  end

  def init(opts) do
    {:ok, socket} = PhoenixClient.Socket.start_link(opts)
    {:ok, %{socket: socket, channel: nil}, {:continue, :join_channel}}
  end

  def handle_continue(:join_channel, state) do
    case PhoenixClient.Socket.connected?(state.socket) do
      true ->
        {:ok, _response, channel} = PhoenixClient.Channel.join(state.socket, "device:test")
        {:noreply, %{state | channel: channel}}

      _ ->
        {:noreply, state, {:continue, :join_channel}}
    end
  end

  def handle_cast({:send_temperature_value, value}, state) do
    if state.channel == nil do
      Logger.info("Channel isn't connected, skipping.")
      {:noreply, state}
    end

    PhoenixClient.Channel.push_async(state.channel, "temperature", %{value: value})
    {:noreply, state}
  end

  def handle_cast({:send_humidity_value, value}, state) do
    if state.channel == nil do
      Logger.info("Channel isn't connected, skipping.")
      {:noreply, state}
    end

    PhoenixClient.Channel.push_async(state.channel, "humidity", %{value: value})
    {:noreply, state}
  end
end
