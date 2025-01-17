package main

import (
	"fmt"
	"strings"

	fileUtils "github.com/almaraz333/advent_of_code_2024"
)

func main() {
	lines, _ := fileUtils.ReadFileToLines("./input.txt")
	lettersToCheck := []string{"X", "M", "A", "S"}

	var grid [][]string

	for _, line := range lines {
		var currArr []string

		splitLine := strings.Split(line, "")

		currArr = append(currArr, splitLine...)

		grid = append(grid, currArr)
	}

	res := 0
	resTwo := 0

	for i, row := range grid {
		for j, letter := range row {
			if (letter == lettersToCheck[0]) {
				if checkLeft(grid, i, j - 1, 1, lettersToCheck) {
					res++
				}
				if checkRight(grid, i, j + 1, 1, lettersToCheck) {
					res++
				}
				if checkTop(grid, i - 1, j, 1, lettersToCheck) {
					res++
				}
				if checkBottom(grid, i + 1, j, 1, lettersToCheck) {
					res++
				}
				if topRight(grid, i - 1, j + 1, 1, lettersToCheck) {
					res++
				}
				if bottomRight(grid, i + 1, j + 1, 1, lettersToCheck) {
					res++
				}
				if bottomLeft(grid, i + 1, j - 1, 1, lettersToCheck) {
					res++
				}
				if topLeft(grid, i - 1, j - 1, 1, lettersToCheck) {
					res++
				}	
			}
		}
	}

	for i, row := range grid {
		for j, letter := range row {
			if (letter == "A") {
				if topRight(grid, i - 1, j + 1, 0, []string{"S"}) && bottomRight(grid, i + 1, j + 1, 0, []string{"S"}) && bottomLeft(grid, i + 1, j - 1, 0, []string{"M"}) && topLeft(grid, i - 1, j - 1, 0, []string{"M"}) {
					resTwo++
				} else if topRight(grid, i - 1, j + 1, 0, []string{"S"}) && bottomRight(grid, i + 1, j + 1, 0, []string{"M"}) && bottomLeft(grid, i + 1, j - 1, 0, []string{"M"}) && topLeft(grid, i - 1, j - 1, 0, []string{"S"}) {
					resTwo++
				} else if topRight(grid, i - 1, j + 1, 0, []string{"M"}) && bottomRight(grid, i + 1, j + 1, 0, []string{"S"}) && bottomLeft(grid, i + 1, j - 1, 0, []string{"S"}) && topLeft(grid, i - 1, j - 1, 0, []string{"M"}) {
					resTwo++
				} else if topRight(grid, i - 1, j + 1, 0, []string{"M"}) && bottomRight(grid, i + 1, j + 1, 0, []string{"M"}) && bottomLeft(grid, i + 1, j - 1, 0, []string{"S"}) && topLeft(grid, i - 1, j - 1, 0, []string{"S"}) {
					resTwo++
				} 
			}
		}
	}

	fmt.Println("Part One: ",res,"\n", "Part Two: ", resTwo)
}

func checkLeft(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (columnIndex == -1) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return checkLeft(grid, rowIndex, columnIndex - 1, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func checkRight(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (columnIndex == len(grid[0])) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return checkRight(grid, rowIndex, columnIndex + 1, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func checkTop(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (rowIndex == -1) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return checkTop(grid, rowIndex - 1, columnIndex, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func checkBottom(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (rowIndex == len(grid)) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return checkBottom(grid, rowIndex + 1, columnIndex, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func topRight(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (rowIndex == -1 || columnIndex == len(grid[0])) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return topRight(grid, rowIndex - 1, columnIndex + 1, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func bottomRight(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (rowIndex == len(grid) || columnIndex == len(grid[0])) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return bottomRight(grid, rowIndex + 1, columnIndex + 1, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func bottomLeft(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (rowIndex == len(grid) || columnIndex == -1) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return bottomLeft(grid, rowIndex + 1, columnIndex - 1, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}

func topLeft(grid [][]string, rowIndex int, columnIndex int, letterToCheckIndex int, lettersToCheck []string) bool {
	if (rowIndex == -1 || columnIndex == -1) {
		return false
	}

	currLetter := grid[rowIndex][columnIndex]

	if letterToCheckIndex == len(lettersToCheck) - 1 && currLetter == lettersToCheck[letterToCheckIndex] {
		return true
	}

	if (currLetter == lettersToCheck[letterToCheckIndex]) {
		return topLeft(grid, rowIndex - 1, columnIndex - 1, letterToCheckIndex + 1, lettersToCheck)
	} else {
		return false
	}
}