defmodule PrintingDepartment do
  def total_accessible_paper_rolls do
    part1("lib/day4/input/data")
  end

  def part1(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> count_accessible_paper_rolls()
    |> IO.inspect(label: "Part 1 Result")
  end

  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n", trim: true)
  end

  def count_accessible_paper_rolls(lines_in_grid) do
    lines_in_grid
    |> Enum.with_index()
    |> Enum.reduce(0, fn {line, row_index}, acc ->
      line_above = if row_index > 0, do: Enum.at(lines_in_grid, row_index - 1), else: nil

      line_below =
        if row_index < length(lines_in_grid) - 1,
          do: Enum.at(lines_in_grid, row_index + 1),
          else: nil

      acc + count_accessible_in_line(line, line_above, line_below)
    end)
  end

  def count_accessible_in_line(line, line_above, line_below) do
    line
    |> String.graphemes()
    |> Enum.with_index()
    |> Enum.reduce(0, fn {char, index}, acc ->
      chars_above = if line_above, do: get_positions_in_line(index, line_above), else: []
      chars = get_positions_in_line(index, line, false)
      chars_below = if line_below, do: get_positions_in_line(index, line_below), else: []

      case is_paper?(char) do
        true -> if is_accessible?(chars_above, chars, chars_below), do: acc + 1, else: acc
        false -> acc
      end
    end)
  end

  def is_paper?(char) do
    char == "@"
  end

  def is_accessible?(chars_above, chars, chars_below) do
    amount_of_paper_rolls =
      [chars_above, chars, chars_below]
      |> List.flatten()
      |> Enum.filter(&is_paper?/1)
      |> length()

    amount_of_paper_rolls < 4
  end

  def get_positions_in_line(index, line, include_middle \\ true) do
    case include_middle do
      true -> [index - 1, index, index + 1]
      false -> [index - 1, index + 1]
    end
    |> Enum.map(fn i ->
      if i >= 0 and i < String.length(line), do: String.at(line, i), else: nil
    end)
  end
end
