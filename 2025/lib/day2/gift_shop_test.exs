defmodule GiftShopTest do
  use ExUnit.Case

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
