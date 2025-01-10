package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	fileUtils "github.com/almaraz333/advent_of_code_2024"
)

func checkSafety(report []string) bool {
	currStatus := true

	for index, item := range report {
		if index == 0 || index == len(report) - 1{
			continue
		}

		curr, _ := strconv.Atoi(item)
		prev, _ := strconv.Atoi(report[index - 1])
		next, _ := strconv.Atoi(report[index + 1])

		if (prev < curr && next < curr) || (prev > curr && next > curr) {
			currStatus = false
			continue
		}

		if (prev == curr || next == curr) {
			currStatus = false
			continue
		}

		floatPrev := float64(prev)
		floatNext := float64(next)
		floatCurr := float64(curr)

		if math.Abs(floatPrev - floatCurr) > 3 || math.Abs(floatNext - floatCurr) > 3 {
			currStatus = false
		}
	 }

	 return currStatus
}

func main() {
	lines, _ := fileUtils.ReadFileToLines("./Day2/input.txt")

	var safe_reports int

	for _, line := range lines { 
		line_slice := strings.Split(line, " ")

		status := true
		mainRes := checkSafety(line_slice)

		if !mainRes {
			onePassing := false

			for index := range line_slice {
				newSlice := make([]string, len(line_slice))
				copy(newSlice, line_slice)

				newReport := append(newSlice[:index], newSlice[index+1:]...)
				
				currRes := checkSafety(newReport)

				if currRes {
					onePassing = true
					break
				}
			}

			if !onePassing {
				status = false
			}
		}

		 if status {
			safe_reports += 1
		}
    } 

	fmt.Println(safe_reports)

}