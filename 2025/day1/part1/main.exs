defmodule Moves do
  @type direction :: :R | :L
  @type move :: {direction(), non_neg_integer()}

  def calculate_password() do
    moves = parse_input!()

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

  defp parse_input!() do
    Path.expand("./input/data")
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(&parse_line!/1)
  end

  defp parse_line!(line) do
    <<dir::binary-size(1), dist::binary>> = line

    direction =
      case dir do
        "R" -> :R
        "L" -> :L
      end

    {direction, String.to_integer(dist)}
  end
end

IO.inspect(Moves.calculate_password())
