defmodule SensorTest do
  use ExUnit.Case
  doctest Sensor

  test "greets the world" do
    assert Sensor.hello() == :world
  end
end
