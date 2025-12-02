defmodule GiftShopTest do
  use ExUnit.Case

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
