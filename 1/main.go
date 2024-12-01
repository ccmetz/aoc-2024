package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Read the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close file when main() stops running
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		lineSlice := strings.Fields(scanner.Text())
		if len(lineSlice) == 2 {
			// Add values to slices
			list1 = convertAndAddToList(list1, lineSlice[0])
			list2 = convertAndAddToList(list2, lineSlice[1])
		}
	}

	// Sort them
	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Printf("List 1 size: %d (sorted? %t)\n", len(list1), slices.IsSorted(list1))
	fmt.Printf("List 2 size: %d (sorted? %t)\n", len(list2), slices.IsSorted(list2))

	// Calculate and sum up total distance between each element in the lists
	totalDistance := 0
	for index, val1 := range list1 {
		val2 := list2[index]
		totalDistance += absOfInt(val1 - val2)
	}

	// Part 1 Answer:
	fmt.Printf("Total Distance: %d\n", totalDistance)

	// Calculate similarity scores
	// Note: To be more efficient, I could do this within the same for-loop for calculating
	// totalDistance up above.
	similarityScore := 0
	for _, val1 := range list1 {
		occurrences := 0
		for _, val2 := range list2 {
			if val2 == val1 {
				occurrences += 1
			}
		}

		similarityScore += occurrences * val1
	}

	// Part 2 Answer:
	fmt.Printf("Similarity Score: %d\n", similarityScore)
}

// Converts the text to an int and appends it to the int slice
func convertAndAddToList(list []int, text string) []int {
	i, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	return append(list, i)
}

// Returns absolute value of an int
func absOfInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
