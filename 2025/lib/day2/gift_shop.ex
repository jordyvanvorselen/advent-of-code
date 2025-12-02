defmodule GiftShop do
  def sum_of_wrong_identifiers do
    part1("lib/day2/input/data")
    part2("lib/day2/input/data")
  end

  def part1(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> Enum.flat_map(&range_for/1)
    |> Enum.reject(&valid_identifier?/1)
    |> Enum.sum()
    |> IO.inspect(label: "Part 1 Result")
  end

  def part2(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> Enum.flat_map(&range_for/1)
    |> Enum.reject(&valid_identifier_v2?/1)
    |> Enum.sum()
    |> IO.inspect(label: "Part 2 Result")
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

  def valid_identifier_v2?(number) do
    number
    |> Integer.to_string()
    |> is_repeating_pattern?()
    |> Kernel.not()
  end

  defp is_repeating_pattern?(str) do
    str
    |> String.length()
    |> divisors()
    |> Enum.any?(fn chunk_size -> repeats_with_size?(str, chunk_size) end)
  end

  defp divisors(1), do: []

  defp divisors(len) do
    1..div(len, 2)
    |> Enum.filter(fn d -> rem(len, d) == 0 end)
  end

  defp repeats_with_size?(str, chunk_size) do
    str
    |> String.graphemes()
    |> Enum.chunk_every(chunk_size)
    |> Enum.uniq()
    |> length()
    |> Kernel.==(1)
  end
end
