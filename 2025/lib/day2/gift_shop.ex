defmodule GiftShop do
  def valid_identifier?(numbers) do
    if rem(numbers, 7) == 0 and rem(numbers, 11) == 0 do
      false
    else
      true
    end
  end
end
