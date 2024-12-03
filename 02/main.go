package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// google ai generated the data loading the rest is me

	x := isPositive(12)
	_ = x
	y := isPositive(-20)
	_ = y
	reports := [][]int{}
	file, err := os.Open("levels.puzzleValues")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := convertLineToNumbers(line)
		// this is where the stuff i wrote begins
		reports = append(reports, numbers)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	totalSafe := 0
	for indexReports := range reports {
		safe := true
		lastPositive := true
		for index := range reports[indexReports] {
			length := len(reports[indexReports])
			lvl := reports[indexReports]
			x := index + 1
			_ = x // here for debugging reasons
			diff := 0

			if index != length-1 {
				diff = lvl[index] - lvl[index+1]

				if absolute(diff) > 3 || diff == 0 || index != 0 && isPositive(diff) != lastPositive {
					cur := isPositive(diff)
					_ = cur
					last := lastPositive
					_ = last
					safe = false

				}
				lastPositive = isPositive(diff)
			}
		}
		if safe {
			totalSafe += 1

		}
	}
	println("total: ", totalSafe)
}

func isPositive(num int) bool {
	if num > 0 {
		return true
	} else {
		return false
	}
}

func absolute(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func convertLineToNumbers(line string) []int {
	var numbers []int
	for _, strNum := range strings.Split(line, " ") {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue // Skip if not a number
		}
		numbers = append(numbers, num)
	}
	return numbers
}
