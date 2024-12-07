package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input4.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}

	fileStr := string(file)

	//
	// For part 1, I gather up a slice for each possible direction of the grid
	// (Rows, columns, diagonals, etc.) and search them XMAS and SAMX
	//

	// Slice of rows
	rows := strings.Split(fileStr, "\n")

	// Slice of columns
	cols := []string{}
	processCols(&cols, rows)

	// Slice of right-leaning diagonals (only need diags w/ at least 4 chars)
	rightDiagonal := []string{}
	processDiagonal(&rightDiagonal, rows, cols)

	// Slice of left-leaning diagonals (only need diags w/ at least 4 chars)
	leftDiagonal := []string{}
	reverseRows := []string{}
	for _, row := range rows {
		reversedStr := ""
		chars := []rune(row)
		for i := len(row) - 1; i >= 0; i-- {
			reversedStr += string(chars[i])
		}
		reverseRows = append(reverseRows, reversedStr)
	}
	flippedCols := slices.Clone(cols)
	slices.Reverse(flippedCols)
	processDiagonal(&leftDiagonal, reverseRows, flippedCols)

	// Find "XMAS" in all directions
	allCombinations := combine(rows, cols, rightDiagonal, leftDiagonal)
	count := 0
	for _, str := range allCombinations {
		count += countOverlappingStrings(str, "XMAS")
		count += countOverlappingStrings(str, "SAMX")
	}

	// Part 1 Answer:
	fmt.Printf("Part 1 COUNT: %d\n", count)

	//
	// For part 2, I convert the grid to a 2D slice, find each "A" that isn't on
	// a boundary, and check it for 2 crossing MAS's
	//
	grid := [][]string{}
	for _, rowStr := range rows {
		row := []string{}
		for _, char := range rowStr {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	part2Count := 0
	// Search between top and bottom row
	for i := 1; i < len(grid)-1; i++ {
		for j, char := range grid[i] {
			// Skip left and right most columns
			if j == 0 || j == len(grid[i])-1 {
				continue
			}

			if char == "A" && hasCrossingMAS(grid, i, j) {
				part2Count++
			}
		}
	}

	fmt.Printf("Part 2 Answer: %d\n", part2Count)
}

func processCols(cols *[]string, rows []string) {
	for rowIndex, rowStr := range rows {
		row := strings.Split(rowStr, "")
		// Add each char to the front of each col
		if rowIndex == 0 {
			*cols = append(*cols, row...)
			continue
		}

		// Add on chars from subsequent rows to each col
		for colIndex, char := range row {
			(*cols)[colIndex] = (*cols)[colIndex] + char
		}
	}
}

func processDiagonal(diagonals *[]string, rows []string, cols []string) {
	//
	// Build out the right leaning diagonals => /////
	for index, colChar := range strings.Split(rows[0], "") {
		*diagonals = append(*diagonals, colChar)
		appendDownRightDiag(diagonals, rows, index, 1, index-1)
	}

	for index, rowChar := range strings.Split(cols[len(cols)-1], "") {
		if index == 0 {
			// Skip since this diagonal should've been taken care of in the previous for loop
			continue
		}

		*diagonals = append(*diagonals, rowChar)
		appendDownRightDiag(diagonals, rows, index+len(rows)-1, index+1, len(rows)-2)
	}
}

func appendDownRightDiag(rightDiagonal *[]string, rows []string, diagIndex int, curRowIndex int, curColIndex int) {
	if curColIndex < 0 || curRowIndex >= len(rows) {
		// Can't go any further to the left
		return
	}

	(*rightDiagonal)[diagIndex] = (*rightDiagonal)[diagIndex] + strings.Split(rows[curRowIndex], "")[curColIndex]
	appendDownRightDiag(rightDiagonal, rows, diagIndex, curRowIndex+1, curColIndex-1)
}

func combine(rows []string, cols []string, rightDiag []string, leftDiag []string) []string {
	allCombinations := append(rows, cols...)
	allCombinations = append(allCombinations, rightDiag...)
	allCombinations = append(allCombinations, leftDiag...)
	return allCombinations
}

func countOverlappingStrings(haystack string, needle string) int {
	var c int
	for i := range haystack {
		if strings.HasPrefix(haystack[i:], needle) {
			c++
		}
	}
	return c
}

func hasCrossingMAS(grid [][]string, rowIndex int, colIndex int) bool {
	leftDiagStr := grid[rowIndex-1][colIndex-1] + grid[rowIndex][colIndex] + grid[rowIndex+1][colIndex+1]
	rightDiagStr := grid[rowIndex-1][colIndex+1] + grid[rowIndex][colIndex] + grid[rowIndex+1][colIndex-1]
	return (leftDiagStr == "MAS" || leftDiagStr == "SAM") && (rightDiagStr == "MAS" || rightDiagStr == "SAM")
}
