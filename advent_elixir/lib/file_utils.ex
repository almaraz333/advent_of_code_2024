defmodule FileUtils do
  def readFile(path) do
    # The path INSIDE the lib directory
    case File.read("./lib/#{path}") do
      {:ok, text} ->
        String.trim(text)

      {:error, reason} ->
        IO.inspect(reason, label: "Error reading file")
    end
  end
end
