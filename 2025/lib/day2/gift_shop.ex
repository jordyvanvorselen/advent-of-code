defmodule GiftShop do
  def main, do: part1("lib/day2/input/data")

  def part1(path) do
    Path.expand(path)
    |> File.read!()
    |> parse_input()
    |> Enum.flat_map(&range_for/1)
    |> Enum.reject(&valid_identifier?/1)
    |> Enum.sum()
    |> IO.inspect(label: "Part 1 Result")
  end

  def parse_input(input) do
    input
    |> String.trim()
    |> String.split(",", trim: true)
  end

  def range_for(range_string) do
    [start_str, end_str] = String.split(range_string, "-")
    String.to_integer(start_str)..String.to_integer(end_str)
  end

  def valid_identifier?(number) do
    number
    |> Integer.to_string()
    |> then(&{&1, String.length(&1)})
    |> valid_by_halves?()
  end

  defp valid_by_halves?({_str, len}) when rem(len, 2) != 0, do: true

  defp valid_by_halves?({str, len}) do
    str
    |> String.split_at(div(len, 2))
    |> then(fn {first, second} -> first != second end)
  end
end
