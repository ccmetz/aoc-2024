package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Position struct {
	Row int
	Col int
}

type Guard struct {
	Pos       Position
	Direction string // up, down, left, right
}

func (g *Guard) turnRight(grid [][]string) {
	switch g.Direction {
	case "up":
		g.Direction = "right"
		grid[g.Pos.Row][g.Pos.Col] = ">"
	case "down":
		g.Direction = "left"
		grid[g.Pos.Row][g.Pos.Col] = "<"
	case "right":
		g.Direction = "down"
		grid[g.Pos.Row][g.Pos.Col] = "v"
	case "left":
		g.Direction = "up"
		grid[g.Pos.Row][g.Pos.Col] = "^"
	}
}

func (g *Guard) move(grid [][]string) {
	// Mark old position with an X
	grid[g.Pos.Row][g.Pos.Col] = "X"

	switch g.Direction {
	case "up":
		g.Pos.Row -= 1
	case "down":
		g.Pos.Row += 1
	case "left":
		g.Pos.Col -= 1
	case "right":
		g.Pos.Col += 1
	}

	// Mark current position
	grid[g.Pos.Row][g.Pos.Col] = getDirectionalChar(g.Direction)
}

func (g *Guard) exit(grid [][]string) {
	grid[g.Pos.Row][g.Pos.Col] = "X"
}

// Added a "visual" mode for fun with smaller inputs
// Run "go run main.go visual" to see the guard move around the grid within your terminal
func main() {
	visualMode := false
	if len(os.Args) > 1 && os.Args[1] == "visual" {
		visualMode = true
	}

	file, err := os.ReadFile("./input6.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}

	// Build 2D grid of the lab
	originalGrid := [][]string{}
	originalGuard := Guard{}
	for rowIdx, row := range strings.Split(string(file), "\n") {
		rowSlice := []string{}
		for colIdx, char := range row {
			if slices.Contains([]string{"^", ">", "<", "v"}, string(char)) {
				originalGuard.Pos.Row = rowIdx
				originalGuard.Pos.Col = colIdx
				originalGuard.Direction = getDirection(string(char))
			}
			rowSlice = append(rowSlice, string(char))
		}
		originalGrid = append(originalGrid, rowSlice)
	}

	//
	// Part 1:
	grid := copyGrid(originalGrid)
	guard := copyGuard(originalGuard)
	printGrid(grid, visualMode)

	distinctPositions := []Position{}
	for !isOnBorderAndFacingOutward(grid, guard) {
		if isFacingObstacle(grid, guard) {
			guard.turnRight(grid)
			continue
		}

		distinctPositions = addDistinctPosition(distinctPositions, guard.Pos)
		guard.move(grid)
		printGrid(grid, visualMode)
	}

	// Exit the lab
	distinctPositions = addDistinctPosition(distinctPositions, guard.Pos)
	guard.exit(grid)

	fmt.Printf("Part 1 Answer: %d\n", len(distinctPositions))

	//
	// Part 2:
	loopCount := 0
	for _, distinctPos := range distinctPositions {
		if originalGuard.Pos.Row == distinctPos.Row && originalGuard.Pos.Col == distinctPos.Col {
			// Skip place where guard already exists
			continue
		}

		resetGrid := copyGrid(originalGrid)
		resetGuard := copyGuard(originalGuard)
		if checkGuardPatrolForLoop(resetGrid, resetGuard, distinctPos) {
			loopCount++
		}
	}

	fmt.Printf("Part 2 Answer: %d\n", loopCount)
}

func checkGuardPatrolForLoop(grid [][]string, guard Guard, obstaclePos Position) bool {
	// Add the obstacle to the grid
	grid[obstaclePos.Row][obstaclePos.Col] = "#"
	guardHistory := []Guard{guard}
	for !isOnBorderAndFacingOutward(grid, guard) {
		if isFacingObstacle(grid, guard) {
			guard.turnRight(grid)
			continue
		}

		guard.move(grid)

		if isGuardInLoop(guard, guardHistory) {
			return true
		}
		guardHistory = append(guardHistory, copyGuard(guard))
	}

	return false
}

// Guard is in a loop if they end up in the exact same position, facing the same direction as
// a previous position they held in the past
func isGuardInLoop(guard Guard, guardHistory []Guard) bool {
	for _, history := range guardHistory {
		if history.Pos.Row == guard.Pos.Row && history.Pos.Col == guard.Pos.Col && history.Direction == guard.Direction {
			return true
		}
	}
	return false
}

func isFacingObstacle(grid [][]string, guard Guard) bool {
	if isOnBorderAndFacingOutward(grid, guard) {
		log.Fatal("Error: Cannot check for obstacle when guard is on border and facing outward")
	}

	nextPos := Position{Row: guard.Pos.Row, Col: guard.Pos.Col}
	switch guard.Direction {
	case "up":
		nextPos.Row -= 1
	case "down":
		nextPos.Row += 1
	case "left":
		nextPos.Col -= 1
	case "right":
		nextPos.Col += 1
	}

	return grid[nextPos.Row][nextPos.Col] == "#"
}

func isOnBorderAndFacingOutward(grid [][]string, guard Guard) bool {
	facingUp := guard.Pos.Row == 0 && guard.Direction == "up"
	facingDown := guard.Pos.Row == len(grid)-1 && guard.Direction == "down"
	facingLeft := guard.Pos.Col == 0 && guard.Direction == "left"
	facingRight := guard.Pos.Col == len(grid[0])-1 && guard.Direction == "right"
	return facingUp || facingDown || facingLeft || facingRight
}

func getDirection(directionalChar string) string {
	switch directionalChar {
	case "^":
		return "up"
	case ">":
		return "right"
	case "<":
		return "left"
	case "v":
		return "down"
	}

	return ""
}

func getDirectionalChar(direction string) string {
	switch direction {
	case "up":
		return "^"
	case "down":
		return "v"
	case "left":
		return "<"
	case "right":
		return ">"
	}

	return ""
}

func addDistinctPosition(distinctPositions []Position, pos Position) []Position {
	if !slices.Contains(distinctPositions, pos) {
		distinctPositions = append(distinctPositions, pos)
	}

	return distinctPositions
}

// Helper func for printing grid in visual mode
func printGrid(grid [][]string, enabled bool) {
	if !enabled {
		return
	}

	time.Sleep(time.Millisecond * 200)
	fmt.Printf("\033[0;0H")
	for _, row := range grid {
		fmt.Println(row)
	}
}

func copyGuard(guard Guard) Guard {
	return Guard{Pos: Position{Row: guard.Pos.Row, Col: guard.Pos.Col}, Direction: guard.Direction}
}

func copyGrid(grid [][]string) [][]string {
	newGrid := [][]string{}
	for _, row := range grid {
		newGrid = append(newGrid, slices.Clone(row))
	}
	return newGrid
}
