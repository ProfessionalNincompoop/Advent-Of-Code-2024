package main

// most of this is frankentined copy pasted code cuz i dont know algroritims :ThumpsUPsmiley:
import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// first google  summariery
func remove(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice

}

// google summariey checked i
func findSmallest(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Return 0 for an empty list, you can handle it differently if needed
	}

	smallest := numbers[0] // Assume the first element is the smallest initially
	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest

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

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	g1small := []int{}
	g2small := []int{}

	for { // go while true loops are just 'for' and nothing else (this was for my girlfriend cuz she doesnt know go)
		small1 := findSmallest(group1)
		group1 = remove(group1, small1)
		g1small = append(g1small, small1)

		small2 := findSmallest(group2)
		group2 = remove(group2, small2)
		g2small = append(g2small, small2)

		if len(group2) == 0 {
			break
		}

	}
	// final loop
	currentdistance := 0
	for index := range g1small {
		dist := 0
		if g1small[index] > g2small[index] {
			dist = g1small[index] - g2small[index]
		} else {
			dist = g2small[index] - g1small[index]
		}
		debug1 := g1small[index]
		debug2 := g2small[index]
		println(debug1)
		println(debug2)

		currentdistance = currentdistance + dist
	}
	println(currentdistance)
}
