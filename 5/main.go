package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type OrderRule struct {
	// Value that this rule is applied to
	Value string
	// all numbers that must be before this value
	Before []string
	// all numbers that must be after this value
	After []string
}

type OrderRuleMap map[string]OrderRule

func (m OrderRuleMap) AddBeforeNum(key string, val string) {
	rule := m[key]
	rule.Before = append(rule.Before, val)
	m[key] = rule
}

func (m OrderRuleMap) AddAfterNum(key string, val string) {
	rule := m[key]
	rule.After = append(rule.After, val)
	m[key] = rule
}

func main() {
	file, err := os.ReadFile("./input5.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}

	fileParts := strings.Split(string(file), "\n\n")
	orderRules := strings.Split(fileParts[0], "\n")
	pageUpdates := strings.Split(fileParts[1], "\n")

	//
	// Part 1:
	orderMap := OrderRuleMap{}
	for _, rule := range orderRules {
		ruleNums := strings.Split(rule, "|")
		str1 := ruleNums[0]
		str2 := ruleNums[1]

		if orderMap[str1].Value == "" {
			orderMap[str1] = OrderRule{Value: str1, Before: []string{}, After: []string{str2}}
		} else {
			orderMap.AddAfterNum(str1, str2)
		}

		if orderMap[str2].Value == "" {
			orderMap[str2] = OrderRule{Value: str2, Before: []string{str1}, After: []string{}}
		} else {
			orderMap.AddBeforeNum(str2, str1)
		}
	}

	correctPageUpdates := [][]string{}
	incorrectPageUpdates := [][]string{}
	for _, updates := range pageUpdates {
		correct := true
		pages := strings.Split(updates, ",")
		for pageIdx, page := range pages {
			for comparedIdx, comparedPage := range pages {
				if page == comparedPage {
					continue
				}

				rule := orderMap[page]
				if rule.Value == "" {
					// Not in map -- skip comparisons against this page
					break
				}

				if slices.Contains(rule.Before, comparedPage) && pageIdx < comparedIdx {
					// Breaks the before rule, page is before the comparedPage
					correct = false
					break
				}

				if slices.Contains(rule.After, comparedPage) && pageIdx > comparedIdx {
					// Breaks the after rule, page is after the comparedPage
					correct = false
					break
				}
			}

			if !correct {
				// If one page is incorrect, don't test the rest of them
				break
			}
		}

		if correct {
			// If this row of pages is determined to be correct, then add it to the list of correct page updates
			correctPageUpdates = append(correctPageUpdates, pages)
		} else {
			incorrectPageUpdates = append(incorrectPageUpdates, pages)
		}
	}

	sum := getSumOfMiddleNums(correctPageUpdates)
	fmt.Printf("Part 1 Answer: %d\n", sum)

	//
	// Part 2:
	orderedPageUpdates := [][]string{}
	for _, incorrectUpdates := range incorrectPageUpdates {
		orderedUpdates := []string{}
		for index, page := range incorrectUpdates {
			if index == 0 {
				orderedUpdates = append(orderedUpdates, page)
				continue
			}

			rule := orderMap[page]
			if rule.Value == "" {
				orderedUpdates = append(orderedUpdates, page)
				continue
			}

			beforeIndex := -1
			for orderedIndex, orderedPage := range orderedUpdates {
				if slices.Contains(rule.After, orderedPage) {
					beforeIndex = orderedIndex
					break
				}
			}

			if beforeIndex != -1 {
				orderedUpdates = insertBefore(orderedUpdates, beforeIndex, page)
			} else {
				orderedUpdates = insertAfter(orderedUpdates, len(orderedUpdates)-1, page)
			}
		}

		orderedPageUpdates = append(orderedPageUpdates, orderedUpdates)
	}

	part2Sum := getSumOfMiddleNums(orderedPageUpdates)
	fmt.Printf("Part 2 Answer: %d\n", part2Sum)
}

func getSumOfMiddleNums(correctPageUpdates [][]string) int {
	sum := 0
	for _, pages := range correctPageUpdates {
		nums := []int{}
		for _, page := range pages {
			nums = append(nums, stringToInt(page))
		}

		middleNum := nums[len(nums)/2]
		sum += middleNum
	}

	return sum
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error converting string to int")
	}

	return num
}

func insertBefore(slice []string, index int, val string) []string {
	beforeSlice := slice[:index]
	afterSlice := slice[index:]
	return append(beforeSlice, append([]string{val}, afterSlice...)...)
}

func insertAfter(slice []string, index int, val string) []string {
	beforeSlice := slice[:index+1]
	afterSlice := slice[index+1:]
	return append(beforeSlice, append([]string{val}, afterSlice...)...)
}
