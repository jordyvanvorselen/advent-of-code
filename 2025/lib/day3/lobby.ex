defmodule Lobby do
  def total_output_joltage do
    part1("lib/day3/input/data")
  end

  def part1(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> find_two_largest_numbers_for_each_battery()
    |> Enum.map(fn [{num1, _}, {num2, _}] -> num1 <> num2 end)
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum()
    |> IO.inspect(label: "Part 1 - Total Output Joltage")
  end

  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n", trim: true)
  end

  def find_two_largest_numbers_for_each_battery(numbers) do
    numbers
    |> Enum.map(&String.graphemes/1)
    |> Enum.map(&find_largest_numbers_within_battery/1)
  end

  def find_largest_numbers_within_battery(battery) do
    indexed = Enum.with_index(battery)

    for {num1, idx1} <- indexed,
        {num2, idx2} <- indexed,
        idx1 < idx2 do
      [{num1, idx1}, {num2, idx2}]
    end
    |> Enum.max_by(fn [{num1, _}, {num2, _}] ->
      String.to_integer(num1 <> num2)
    end)
  end
end
