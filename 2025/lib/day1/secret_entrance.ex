defmodule SecretEntrance do
  @type direction :: :R | :L
  @type move :: {direction(), non_neg_integer()}

  def calculate_passwords do
    part1("lib/day1/input/data")
    part2("lib/day1/input/data")
  end

  def part1(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> calculate_password_part1()
    |> IO.inspect(label: "Part 1 Result")
  end

  def part2(path) do
    path
    |> Path.expand()
    |> File.read!()
    |> parse_input()
    |> calculate_password_part2()
    |> IO.inspect(label: "Part 2 Result")
  end

  def parse_input(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(&parse_line!/1)
  end

  def parse_line!(line) do
    <<dir::binary-size(1), dist::binary>> = line

    direction =
      case dir do
        "R" -> :R
        "L" -> :L
      end

    {direction, String.to_integer(dist)}
  end

  def calculate_password_part1(moves) do
    {_final_position, zero_count} =
      Enum.reduce(moves, {50, 0}, fn {direction, distance}, {current_position, zero_count} ->
        new_position =
          case direction do
            :R -> rem(current_position + distance, 100)
            :L -> rem(current_position - distance + 100, 100)
          end

        new_zero_count =
          if new_position == 0 do
            zero_count + 1
          else
            zero_count
          end

        {new_position, new_zero_count}
      end)

    zero_count
  end

  def calculate_password_part2(moves) do
    {_final_position, zero_count} =
      Enum.reduce(moves, {50, 0}, fn {direction, distance}, {current_position, zero_count} ->
        new_position =
          case direction do
            :R -> rem(current_position + distance, 100)
            :L -> Integer.mod(current_position - distance, 100)
          end

        times_crossed_zero =
          case direction do
            :R ->
              div(current_position + distance, 100)

            :L when distance >= current_position and current_position > 0 ->
              div(distance - current_position, 100) + 1

            :L when current_position == 0 ->
              div(distance, 100)

            :L ->
              0
          end

        {new_position, zero_count + times_crossed_zero}
      end)

    zero_count
  end
end
