defmodule CafeteriaTest do
  use ExUnit.Case

  describe "part1/1" do
    test "returns amount of fresh ids" do
      result = Cafeteria.part1("lib/day5/input/test_data")

      assert result == 3
    end
  end

  describe "part2/1" do
    test "returns total amount of numbers in ranges" do
      result = Cafeteria.part2("lib/day5/input/test_data")

      assert result == 14
    end
  end

  describe "count_total_amount_of_numbers_in_ranges/1" do
    test "returns total amount of numbers in ranges" do
      ranges = [
        "3-5",
        "10-14",
        "16-20",
        "12-18"
      ]

      result = Cafeteria.count_total_amount_of_numbers_in_ranges(ranges)

      assert result == 14
    end
  end
end
