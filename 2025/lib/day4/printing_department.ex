defmodule PrintingDepartment do
  def total_accessible_paper_rolls do
    part1("lib/day4/input/data")
    part2("lib/day4/input/data")
  end

  def part1(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> get_accessible_paper_rolls()
    |> length()
    |> IO.inspect(label: "Part 1 Result")
  end

  def part2(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> get_accessible_paper_rolls_part_2()
    |> length()
    |> IO.inspect(label: "Part 2 Result")
  end

  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n", trim: true)
  end

  def get_accessible_paper_rolls_part_2(lines_in_grid) do
    accessible_rolls = lines_in_grid |> get_accessible_paper_rolls()
    new_grid = accessible_rolls |> remove_paper_rolls_that_were_accessible(lines_in_grid)

    if length(accessible_rolls) > 0,
      do: accessible_rolls ++ get_accessible_paper_rolls_part_2(new_grid),
      else: []
  end

  def remove_paper_rolls_that_were_accessible(accessible_rolls, lines_in_grid) do
    lines_in_grid
    |> Enum.with_index()
    |> Enum.map(fn {line, row_index} ->
      accessible_in_line =
        accessible_rolls
        |> Enum.filter(fn {_char, r_index, _c_index} -> r_index == row_index end)

      remove_accessible_roles_from_line(line, accessible_in_line)
    end)
  end

  def remove_accessible_roles_from_line(line, accessible_rolls) do
    line
    |> String.graphemes()
    |> Enum.with_index()
    |> Enum.map(fn {char, index} ->
      case Enum.any?(accessible_rolls, fn {_c, _r_index, c_index} -> c_index == index end) do
        true -> "."
        false -> char
      end
    end)
    |> Enum.join()
  end

  def get_accessible_paper_rolls(lines_in_grid) do
    lines_in_grid
    |> Enum.with_index()
    |> Enum.reduce([], fn {line, row_index}, acc ->
      line_above = if row_index > 0, do: Enum.at(lines_in_grid, row_index - 1), else: nil

      line_below =
        if row_index < length(lines_in_grid) - 1,
          do: Enum.at(lines_in_grid, row_index + 1),
          else: nil

      [get_accessible_in_line(line, row_index, line_above, line_below) | acc]
    end)
    |> List.flatten()
  end

  def get_accessible_in_line(line, row_index, line_above, line_below) do
    line
    |> String.graphemes()
    |> Enum.with_index()
    |> Enum.filter(fn {char, index} ->
      chars_above = if line_above, do: get_positions_in_line(index, line_above), else: []
      chars = get_positions_in_line(index, line, false)
      chars_below = if line_below, do: get_positions_in_line(index, line_below), else: []

      case is_paper?(char) do
        true -> is_accessible?(chars_above, chars, chars_below)
        false -> false
      end
    end)
    |> Enum.map(fn {char, index} -> {char, row_index, index} end)
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
