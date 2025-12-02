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

        number_of_times_passed_zero =
          case direction do
            :R -> div(current_position + distance, 100)
            :L -> div(distance - current_position + 99, 100)
          end

        times_ended_on_zero =
          case new_position do
            0 -> zero_count + 1
            _ -> zero_count
          end

        {new_position, times_ended_on_zero + number_of_times_passed_zero}
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
