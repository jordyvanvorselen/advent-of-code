defmodule LobbyTest do
  use ExUnit.Case

  describe "part1/1" do
    test "returns total output joltage" do
      result = Lobby.part1("lib/day3/input/test_data")

      assert result == 357
    end
  end

  describe "part2/1" do
    test "returns total output joltage" do
      result = Lobby.part2("lib/day3/input/test_data")

      assert result == 3_121_910_778_619
    end
  end

  describe "find_largest_numbers_within_battery/1" do
    for {value, expected} <- [
          {~w(9 8 7 6 5 4 3 2 1 1 1 1 1 1 1), [{"9", 0}, {"8", 1}]},
          {~w(8 1 1 1 1 1 1 1 1 1 1 1 1 1 9), [{"8", 0}, {"9", 14}]},
          {~w(2 3 4 2 3 4 2 3 4 2 3 4 2 7 8), [{"7", 13}, {"8", 14}]},
          {~w(8 1 8 1 8 1 9 1 1 1 1 2 1 1 1), [{"9", 6}, {"2", 11}]}
        ] do
      test "#{inspect(value)} returns #{inspect(expected)}" do
        assert Lobby.find_largest_numbers_within_battery(unquote(value)) == unquote(expected)
      end
    end
  end
end
