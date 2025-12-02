defmodule GiftShopTest do
  use ExUnit.Case

  describe "part1/1" do
    test "calculates the sum of invalid identifiers from the input file" do
      result = GiftShop.part1("lib/day2/input/test_data")

      assert result == 1_227_775_554
    end
  end

  describe "part2/1" do
    test "placeholder for part2 tests" do
      result = GiftShop.part2("lib/day2/input/test_data")

      assert result == 4_174_379_265
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

  describe "valid_identifier_v2?/1" do
    for {value, expected} <- [
          {1_111_111, false},
          {11, false},
          {22, false},
          {99, false},
          {111, false},
          {999, false},
          {1010, false},
          {1_188_511_885, false},
          {222_222, false},
          {446_446, false},
          {38_593_859, false},
          {565_656, false},
          {824_824_824, false},
          {2_121_212_121, false}
        ] do
      test "#{value} returns #{expected}" do
        assert GiftShop.valid_identifier_v2?(unquote(value)) == unquote(expected)
      end
    end
  end
end
