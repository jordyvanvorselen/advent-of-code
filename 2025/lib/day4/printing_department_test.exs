defmodule PrintingDepartmentTest do
  use ExUnit.Case

  describe "part1/1" do
    test "returns how many paper rolls are accessible" do
      result = PrintingDepartment.part1("lib/day4/input/test_data")

      assert result == 13
    end
  end

  describe "part2/1" do
    test "returns how many paper rolls are accessible" do
      result = PrintingDepartment.part2("lib/day4/input/test_data")

      assert result == 43
    end
  end

  describe "get_positions_in_line/3" do
    test "returns positions in line including middle" do
      line = "ABCDEFGH"

      result = PrintingDepartment.get_positions_in_line(3, line, true)

      assert result == ["C", "D", "E"]
    end

    test "returns positions in line excluding middle" do
      line = "ABCDEFGH"

      result = PrintingDepartment.get_positions_in_line(3, line, false)

      assert result == ["C", "E"]
    end
  end

  describe "is_accessible?/3" do
    test "returns false if 4 or more paper rolls around" do
      chars_above = ["@", "@", "@"]
      chars = ["@", "."]
      chars_below = [".", ".", "."]

      result = PrintingDepartment.is_accessible?(chars_above, chars, chars_below)

      assert result == false
    end

    test "returns true if less than 4 paper rolls around" do
      chars_above = ["@", ".", "@"]
      chars = ["@", "."]
      chars_below = [".", ".", "."]

      result = PrintingDepartment.is_accessible?(chars_above, chars, chars_below)

      assert result == true
    end
  end
end
