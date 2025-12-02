defmodule GiftShopTest do
  use ExUnit.Case

  describe "part1/1" do
    test "calculates the sum of invalid identifiers from the input file" do
      result = GiftShop.part1("lib/day2/input/test_data")

      assert result == 1_227_775_554
    end
  end

  describe "parse_input/1" do
    test "parses comma-separated values into a list" do
      input = "5-10,15-20,25-30"
      expected_output = ["5-10", "15-20", "25-30"]
      assert GiftShop.parse_input(input) == expected_output
    end
  end

  describe "range_for/1" do
    test "returns valid range for two numbers" do
      assert GiftShop.range_for("5-10") == 5..10
      assert GiftShop.range_for("1-3") == 1..3
    end
  end

  describe "valid_identifier?/1" do
    for {value, expected} <- [
          {101, true},
          {101_101, false}
        ] do
      test "#{value} returns #{expected}" do
        assert GiftShop.valid_identifier?(unquote(value)) == unquote(expected)
      end
    end
  end
end
