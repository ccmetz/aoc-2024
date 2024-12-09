package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/ccmetz/aoc-2024/util"
)

func main() {
	file, err := os.ReadFile("./input8.test")
	if err != nil {
		log.Fatal("Error reading file")
	}

	// Build 2D grid
	grid := [][]string{}
	for _, row := range strings.Split(string(file), "\n") {
		rowSlice := []string{}
		for _, char := range row {
			rowSlice = append(rowSlice, string(char))
		}
		grid = append(grid, rowSlice)
	}
	printGrid(grid)

	// Part 1:
	// Crawl through the grid, once a node is identified, get the diff between
	// each node of the same type and add an antinote that is that distance away
	// on either side of the nodes
	antinodes := [][]int{}
	for rowIdx, row := range grid {
		for colIdx, col := range row {
			if col == "." {
				continue
			}

			// Node found - find all other matching nodes
			node := col
			startRow := rowIdx
			startCol := colIdx + 1
			if colIdx == len(grid[0])-1 {
				startRow = rowIdx + 1
				startCol = 0
			}
			otherNodes := findNodes(grid, startRow, startCol, node)
			fmt.Println(node, otherNodes)

			// Calculate distance between the node each otherNode add an antinode
			// on either side (if in bounds of the grid)
			for _, otherNode := range otherNodes {
				diffRow := otherNode[0] - rowIdx // node row will always be lower or equal to otherNode
				diffCol := util.AbsOfInt(otherNode[1] - colIdx)

				var antinode1 []int
				var antinode2 []int
				if otherNode[1]-colIdx < 0 {
					// Add antinode in same direction outside of node
					antinode1 = []int{rowIdx - diffRow, colIdx + diffCol}
					// Add antinode in same direction outside of otherNode
					antinode2 = []int{otherNode[0] + diffRow, otherNode[1] - diffCol}
				} else {
					// Add antinode in same direction outside of node
					antinode1 = []int{rowIdx - diffRow, colIdx - diffCol}
					// Add antinode in same direction outside of otherNode
					antinode2 = []int{otherNode[0] + diffRow, otherNode[1] + diffCol}
				}

				if isValidAntinode(grid, antinode1) && isUniqueAntinode(antinodes, antinode1) {
					antinodes = append(antinodes, antinode1)
				}

				if isValidAntinode(grid, antinode2) && isUniqueAntinode(antinodes, antinode2) {
					antinodes = append(antinodes, antinode2)
				}
			}
		}
	}

	fmt.Println("Antinodes", antinodes)
	fmt.Printf("Part 1 Answer: %d\n", len(antinodes))
}

func isValidAntinode(grid [][]string, antinode []int) bool {
	rowMax := len(grid) - 1
	colMax := len(grid[0]) - 1
	return antinode[0] >= 0 && antinode[0] <= rowMax && antinode[1] >= 0 && antinode[1] <= colMax
}

func isUniqueAntinode(antinodes [][]int, antinode []int) bool {
	for _, otherAntinode := range antinodes {
		if otherAntinode[0] == antinode[0] && otherAntinode[1] == antinode[1] {
			return false
		}
	}

	return true
}

func findNodes(grid [][]string, startRow int, startCol int, node string) [][]int {
	nodes := [][]int{}
	for i := startRow; i < len(grid); i++ {
		if i != startRow {
			startCol = 0
		}
		for j := startCol; j < len(grid[0]); j++ {
			if grid[i][j] == node {
				nodes = append(nodes, []int{i, j})
			}
		}
	}

	return nodes
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
