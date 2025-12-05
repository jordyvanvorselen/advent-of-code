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
    |> then(fn {ranges, _numbers} ->
      count_total_amount_of_numbers_in_ranges(ranges)
    end)
    |> IO.inspect(label: "Part 2 Result")
  end

  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n", trim: true)
    |> Enum.split_with(fn elem -> String.contains?(elem, "-") end)
  end

  def count_total_amount_of_numbers_in_ranges(ranges) do
    ranges
    |> Enum.map(fn range ->
      [min, max] = String.split(range, "-") |> Enum.map(&String.to_integer/1)
      {min, max}
    end)
    |> Enum.sort()
    |> Enum.reduce([], fn
      {min, max}, [] -> [{min, max}]
      {min, max}, [{prev_min, prev_max} | rest] when min <= prev_max + 1 ->
        [{prev_min, max(max, prev_max)} | rest]
      range, acc -> [range | acc]
    end)
    |> Enum.reduce(0, fn {min, max}, acc -> acc + (max - min + 1) end)
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
