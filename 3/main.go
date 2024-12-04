package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input3.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}

	corruptedMem := string(file)

	// Part 1:
	part1Sum := findAndSumMuls(corruptedMem)
	fmt.Printf("Part 1 Answer: %d\n", part1Sum)

	// Part 2:
	var enabledStr strings.Builder
	// Get all strings in between the do() instructions
	splitOnDo := strings.Split(corruptedMem, "do()")
	for _, str := range splitOnDo {
		// The string will only contain don't() at this point, split on that instruction,
		// which means the 0 index of the resulting array would have been between a do() and a don't()/EOF
		splitOnDont := strings.Split(str, "don't()")
		enabledStr.WriteString(splitOnDont[0])
	}

	part2Sum := findAndSumMuls(enabledStr.String())
	fmt.Printf("Part 2 Answer: %d\n", part2Sum)
}

func findAndSumMuls(corruptedMem string) int {
	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	muls := mulRegex.FindAllStringSubmatch(corruptedMem, -1)
	sum := 0
	for _, mul := range muls {
		arg1, err1 := strconv.Atoi(mul[1])
		arg2, err2 := strconv.Atoi(mul[2])
		if err1 != nil || err2 != nil {
			log.Fatal("Error while converting strings to ints")
		}
		sum += arg1 * arg2
	}

	return sum
}
