package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input7.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}

	equations := strings.Split(string(file), "\n")
	//
	// Part 1:
	totalTrueCount := 0
	totalCalibrationResult := 0
	cachedCombinations := make(map[int][][]string)
	for _, equation := range equations {
		value, nums := getEquationParts(equation)
		if canNumsMakeValue(value, nums, cachedCombinations, []string{"+", "*"}) {
			totalCalibrationResult += value
			totalTrueCount++
		}
	}
	fmt.Printf("Part 1 Answer: %d (count: %d)\n", totalCalibrationResult, totalTrueCount)

	//
	// Part 2:
	totalTrueCount2 := 0
	totalCalibrationResult2 := 0
	cachedCombinations2 := make(map[int][][]string)
	for _, equation := range equations {
		value, nums := getEquationParts(equation)
		if canNumsMakeValue(value, nums, cachedCombinations2, []string{"+", "*", "||"}) {
			totalCalibrationResult2 += value
			totalTrueCount2++
		}
	}

	fmt.Printf("Part 2 Answer: %d (count: %d)\n", totalCalibrationResult2, totalTrueCount2)
}

func canNumsMakeValue(value int, nums []int, cachedCombinations map[int][][]string, operatorValues []string) bool {
	// Calculate all different permutations of "+" and "*" operators
	combinations := cachedCombinations[len(nums)-1]
	if len(combinations) == 0 {
		combinations = getCombinations(operatorValues, len(nums)-1)
		cachedCombinations[len(nums)-1] = combinations
	}

	for _, operators := range combinations {
		currentVal := 0
		for idx, operator := range operators {
			if idx == 0 {
				currentVal = nums[0]
			}

			switch operator {
			case "+":
				currentVal += nums[idx+1]
			case "*":
				currentVal = currentVal * nums[idx+1]
			case "||":
				currentVal = convertStringToInt(convertIntToString(currentVal) + convertIntToString(nums[idx+1]))
			default:
				log.Fatalf("operator not recognized: %s", operator)
			}
		}

		if currentVal == value {
			return true
		}
	}

	return false
}

func getCombinations(operators []string, length int) [][]string {
	var result [][]string
	var generate func(current []string, depth int)

	// Recursive function to generate combinations
	generate = func(current []string, depth int) {
		if depth == length {
			result = append(result, current)
			return
		}

		for _, operator := range operators {
			copy := slices.Clone(current)
			generate(append(copy, operator), depth+1)
		}
	}

	generate([]string{}, 0)
	return result
}

func getEquationParts(equation string) (int, []int) {
	parts := strings.Split(equation, ": ")
	value := convertStringToInt(parts[0])
	numStrs := strings.Split(parts[1], " ")
	nums := []int{}
	for _, numStr := range numStrs {
		nums = append(nums, convertStringToInt(numStr))
	}

	return value, nums
}

func convertStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Error converting string to int: %s\n", str)
	}

	return num
}

func convertIntToString(num int) string {
	str := strconv.Itoa(num)
	return str
}
