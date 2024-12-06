package main

import (
	"aoc24/utils"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines, err := utils.GetInput("day2")
	if err != nil {
		panic(err)
	}
	safeCount := 0
	for _, line := range lines {
		levels, err := utils.AtoiSplit(line, " ")
		if err != nil {
			panic(err)
		}
		last := levels[0]
		increasing := true
		safe := true
		for i := 1; i < len(levels); i++ {
			curr := levels[i]
			if i == 1 {
				increasing = curr > last
			}
			if curr == last || (increasing && curr < last) || (!increasing && curr > last) {
				safe = false
				break
			}
			if utils.IntAbs(curr-last) > 3 {
				safe = false
				break
			}
			last = curr
		}
		if safe {
			safeCount++
		}
	}
	fmt.Println("part1", safeCount)
}

func part2() {
	lines, err := utils.GetInput("day2")
	if err != nil {
		panic(err)
	}
	safeCount := 0
	for _, line := range lines {
		levels, err := utils.AtoiSplit(line, " ")
		if err != nil {
			panic(err)
		}
		last := levels[0]
		increasing := true
		warn := false
		safe := true
		for i := 1; i < len(levels); i++ {
			curr := levels[i]
			if i == 1 {
				increasing = curr > last
			}
			if curr == last || (increasing && curr < last) || (!increasing && curr > last) {
				if !warn {
					warn = true
				} else {
					safe = false
					break
				}
			}
			if utils.IntAbs(curr-last) > 3 {
				if !warn {
					warn = true
				} else {
					safe = false
					break
				}
			}
			last = curr
		}
		if safe {
			safeCount++
		}
	}
	fmt.Println("part2", safeCount)
}
