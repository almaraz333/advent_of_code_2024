defmodule DayThree do
  def part_one do
    data =
      FileUtils.readFile("DayThree/input.txt")

    Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)/, data, capture: :all_but_first)
    |> Enum.map(fn [first, second] ->
      String.to_integer(first) * String.to_integer(second)
    end)
    |> Enum.sum()
  end
end
