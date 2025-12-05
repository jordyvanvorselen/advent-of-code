defmodule CafeteriaTest do
  use ExUnit.Case

  describe "part1/1" do
    test "returns amount of fresh ids" do
      result = Cafeteria.part1("lib/day5/input/test_data")

      assert result == 3
    end
  end
end
