package util

import (
	"log"
	"strconv"
)

// Converts the text to an int and appends it to the int slice
func ConvertAndAddToList(list []int, text string) []int {
	i, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	return append(list, i)
}

// Returns absolute value of an int
func AbsOfInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
