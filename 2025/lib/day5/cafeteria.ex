defmodule Cafeteria do
  def amount_of_fresh_ingredients do
    part1("lib/day5/input/data")
    part2("lib/day5/input/data")
  end

  def part1(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> then(fn {ranges, numbers} ->
      count_numbers_in_ranges(numbers, ranges)
    end)
    |> IO.inspect(label: "Part 1 Result")
  end

  def part2(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
  end

  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n", trim: true)
    |> Enum.split_with(fn elem -> String.contains?(elem, "-") end)
  end

  def count_numbers_in_ranges(numbers, ranges) do
    Enum.count(numbers, fn number ->
      falls_in_range(String.to_integer(number), ranges)
    end)
  end

  def falls_in_range(number, ranges) do
    Enum.any?(ranges, fn range ->
      [min, max] = String.split(range, "-") |> Enum.map(&String.to_integer/1)
      number >= min and number <= max
    end)
  end
end
