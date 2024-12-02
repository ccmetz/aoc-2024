package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/ccmetz/aoc-2024/util"
)

func main() {
	// Read the input file
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Close file when main() stops running
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	safeReportCount := 0
	safeReportCountWithDampener := 0
	for scanner.Scan() {
		lineSlice := strings.Fields(scanner.Text())
		report := []int{}
		for _, val := range lineSlice {
			report = util.ConvertAndAddToList(report, val)
		}

		isSafe := isReportSafe(report)

		if isSafe {
			safeReportCount++
			safeReportCountWithDampener++
		} else {
			// Remove one element at a time and try again (applying the problem dampener per challenge)
			for unsafeLevel := 0; unsafeLevel < len(report); unsafeLevel++ {
				dampenedReport := removeFromSlice(slices.Clone(report), unsafeLevel)
				isDampenedSafe := isReportSafe(dampenedReport)
				if isDampenedSafe {
					safeReportCountWithDampener++
					break
				}
			}
		}
	}

	// Part 1 Answer:
	log.Printf("There are %d safe reports\n", safeReportCount)
	// Part 2 Answer:
	log.Printf("There are %d safe reports with the problem dampener\n", safeReportCountWithDampener)
}

// Checks if the report is safe. Returns true or false. If false, it will
// also return an int slice representing the indexes of the unsafe levels in the report.
func isReportSafe(report []int) bool {
	direction := "increasing"
	lastVal := 0

	for index, val := range report {
		if index == 0 {
			lastVal = val
			continue
		}

		diff := val - lastVal
		if diff == 0 {
			// No increase, report not safe
			return false
		}

		// Check what direction the report is going
		if index == 1 && diff < 0 {
			direction = "decreasing"
		} else if index == 1 && diff > 0 {
			direction = "increasing"
		}

		// If directions changed, the report is considered unsafe
		if index > 1 && (diff < 0 && direction == "increasing") ||
			(diff > 0 && direction == "decreasing") {
			return false
		}

		// Check that the report is gradually increasing/decreasing
		if direction == "increasing" && (diff > 3 || diff < 1) {
			return false
		}
		if direction == "decreasing" && (diff > -1 || diff < -3) {
			return false
		}

		// Track this val for the next iteration
		lastVal = val
	}

	return true
}

// Removes the unsafeLevel from the slice
func removeFromSlice(slice []int, unsafeLevel int) []int {
	return append(slice[:unsafeLevel], slice[unsafeLevel+1:]...)
}
