package main

// most of this is frankentined copy pasted code cuz i dont know algroritims :ThumpsUPsmiley:
import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findOccurance(numbers []int, value int) int {
	count := 0
	for index := range numbers {
		if numbers[index] == value {
			count += 1
		}
	}
	return count
}

func main() {

	// google first result frankenstine
	file, _ := os.Open("numbers.puzzleValues")
	defer file.Close()
	group1 := []int{}
	group2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into numbers
		numbers := strings.Split(line, "   ") // Assuming space-separated numbers

		// Convert numbers to integers and add to lists
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue // Skip if not a valid number
			}

			// Alternate adding numbers to lists
			if len(group1) <= len(group2) {
				group1 = append(group1, num)
			} else {
				group2 = append(group2, num)
			}
		}
	}

	score := 0
	for index := range group1 {
		score += group1[index] * findOccurance(group2, group1[index])
	}
	println(score)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
