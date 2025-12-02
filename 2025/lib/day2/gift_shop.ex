defmodule GiftShop do
  def range_for(range_string) do
    [start_str, end_str] = String.split(range_string, "-")
    String.to_integer(start_str)..String.to_integer(end_str)
  end

  def valid_identifier?(numbers) do
    if rem(numbers, 7) == 0 and rem(numbers, 11) == 0 do
      false
    else
      true
    end
  end
end
