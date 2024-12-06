package main

import (
	"aoc24/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func part1() {
	input := utils.MustGetInputAsString("day5")
	split := strings.Split(input, "\n\n")
	rulesStr := split[0]
	updatesStr := split[1]

	ruleMap := make(map[string][]string)
	ruleMapInverted := make(map[string][]string)
	for _, rule := range strings.Split(rulesStr, "\n") {
		components := strings.Split(rule, "|")
		ruleMap[components[0]] = append(ruleMap[components[0]], components[1])
		ruleMapInverted[components[1]] = append(ruleMap[components[1]], components[0])
	}

	sum := 0
	sortedSum := 0
	lessFn := func(a, b string) int {
		if Contains(ruleMap[a], b) {
			return -1
		} else if Contains(ruleMapInverted[b], a) {
			return 1
		} else {
			return 0
		}
	}
	for _, update := range strings.Split(updatesStr, "\n") {
		if update == "" {
			continue
		}
		updateSlice := strings.Split(update, ",")
		isSorted := slices.IsSortedFunc(updateSlice, lessFn)
		if isSorted {
			// fmt.Println(updateSlice)
			middle := updateSlice[len(updateSlice)/2]
			n, err := strconv.Atoi(middle)
			if err != nil {
				panic(err)
			}
			sum += n
		} else {
			slices.SortFunc(updateSlice, lessFn)
			middle := updateSlice[len(updateSlice)/2]
			n, err := strconv.Atoi(middle)
			if err != nil {
				panic(err)
			}
			sortedSum += n
		}
	}

	/* fmt.Println(rulesStr)
	fmt.Println(updatesStr) */
	fmt.Println("part1", sum)
	fmt.Println("part2", sortedSum)
}
