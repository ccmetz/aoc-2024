package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/ccmetz/aoc-2024/util"
)

func main() {
	file, err := os.ReadFile("./input8.txt")
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
	//
	// Part 2:
	// I modified the code to instead calculate all possible antinodes on either side of
	// the in-line nodes
	//
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

			if len(otherNodes) == 0 {
				// Add the antenna itself as an antinode since it an antinode will technically take up the same location
				// as an actual antinode
				addAntennaAsAntinode(&antinodes, rowIdx, colIdx)
			}

			// Calculate distance between the node each otherNode add an antinode
			// on either side (if in bounds of the grid).
			// NOTE: Bits and pieces of this logic could probably abstracted into reusable functions to
			// decrease the number of nested if/elses and for-loops
			for _, otherNode := range otherNodes {
				diffRow := otherNode[0] - rowIdx // node row will always be lower or equal to otherNode
				diffCol := util.AbsOfInt(otherNode[1] - colIdx)

				var antinodeGroup1 [][]int
				var antinodeGroup2 [][]int
				if otherNode[1]-colIdx < 0 {
					//
					// Going along the right diagonal (/)
					for {
						var antinode []int
						if len(antinodeGroup1) == 0 {
							antinode = []int{rowIdx - diffRow, colIdx + diffCol}
						} else {
							lastAntinode := antinodeGroup1[len(antinodeGroup1)-1]
							antinode = []int{lastAntinode[0] - diffRow, lastAntinode[1] + diffCol}
						}

						// Out of bounds, stop finding antinodes in this direction
						if !isValidAntinode(grid, antinode) {
							break
						}
						antinodeGroup1 = append(antinodeGroup1, antinode)
					}

					for {
						var antinode []int
						if len(antinodeGroup2) == 0 {
							antinode = []int{otherNode[0] + diffRow, otherNode[1] - diffCol}
						} else {
							lastAntinode := antinodeGroup2[len(antinodeGroup2)-1]
							antinode = []int{lastAntinode[0] + diffRow, lastAntinode[1] - diffCol}
						}

						// Out of bounds, stop finding antinodes in this direction
						if !isValidAntinode(grid, antinode) {
							break
						}
						antinodeGroup2 = append(antinodeGroup2, antinode)
					}
				} else {
					//
					// Going along the left diagonal or straight up/down (\ or |)
					for {
						var antinode []int
						if len(antinodeGroup1) == 0 {
							antinode = []int{rowIdx - diffRow, colIdx - diffCol}
						} else {
							lastAntinode := antinodeGroup1[len(antinodeGroup1)-1]
							antinode = []int{lastAntinode[0] - diffRow, lastAntinode[1] - diffCol}
						}

						// Out of bounds, stop finding antinodes in this direction
						if !isValidAntinode(grid, antinode) {
							break
						}
						antinodeGroup1 = append(antinodeGroup1, antinode)
					}

					for {
						var antinode []int
						if len(antinodeGroup2) == 0 {
							antinode = []int{otherNode[0] + diffRow, otherNode[1] + diffCol}
						} else {
							lastAntinode := antinodeGroup2[len(antinodeGroup2)-1]
							antinode = []int{lastAntinode[0] + diffRow, lastAntinode[1] + diffCol}
						}

						// Out of bounds, stop finding antinodes in this direction
						if !isValidAntinode(grid, antinode) {
							break
						}
						antinodeGroup2 = append(antinodeGroup2, antinode)
					}
				}

				// Append only unique antinode positions to the slice keeping track of all antinodes
				antinodeGroups := append(antinodeGroup1, antinodeGroup2...)
				for _, antinode := range antinodeGroups {
					if isUniqueAntinode(antinodes, antinode) {
						antinodes = append(antinodes, antinode)
					}
				}

				// Add the antenna itself as an antinode since it an antinode will technically take up the same location
				// as an actual antinode
				addAntennaAsAntinode(&antinodes, rowIdx, colIdx)
			}
		}
	}

	fmt.Printf("Part 2 Answer: %d\n", len(antinodes))
}

// Trying out pointers here just for the hell of it. Could just return and
// reassign the antinodes slice instead
func addAntennaAsAntinode(antinodes *[][]int, rowIdx int, colIdx int) {
	antennaAntinode := []int{rowIdx, colIdx}
	if isUniqueAntinode(*antinodes, antennaAntinode) {
		*antinodes = append(*antinodes, antennaAntinode)
	}
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
