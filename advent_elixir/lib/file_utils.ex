defmodule FileUtils do
  def readFile(path) do
    # The path INSIDE the lib directory
    File.stream!("./lib/#{path}")
    |> Enum.to_list()
    |> Enum.map(&String.trim/1)
  end
end
