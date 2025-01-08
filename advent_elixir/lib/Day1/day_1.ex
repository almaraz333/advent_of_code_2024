defmodule Day1 do
  require Integer
  def part_one do
    FileUtils.readFile("Day1/input.txt")
    |> Enum.reduce({[],[]}, fn val, {left_acc, right_acc} ->
      [left, right] = String.split(val, "   ")

      {left_int, _} = Integer.parse(left)
      {right_int, _} = Integer.parse(right)

      {left_acc ++ [left_int], right_acc ++ [right_int]}
    end)
    |> (fn {left, right} ->
    Enum.zip(Enum.sort(left), Enum.sort(right))
    end).()
    |>
    Enum.map(fn {left, right} ->
      abs(left - right)
    end)
    |> Enum.sum()
  end

  def part_two do
    {left, right} = FileUtils.readFile("Day1/input.txt")
    |> Enum.reduce({[],[]}, fn val, {left_acc, right_acc} ->
      [left, right] = String.split(val, "   ")

      {left_int, _} = Integer.parse(left)
      {right_int, _} = Integer.parse(right)

      {left_acc ++ [left_int], right_acc ++ [right_int]}
    end)

    count_map = Enum.reduce(right, [], fn val, acc ->
      Keyword.update(acc, Integer.to_string(val) |> String.to_atom, 1, fn x -> x + 1 end)
    end)

    Enum.map(left, fn val ->
      res = Keyword.get(count_map, Integer.to_string(val) |> String.to_atom, 0)
      res * val
    end)
    |> Enum.sum
  end
end
