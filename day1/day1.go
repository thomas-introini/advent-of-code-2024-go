package main

import (
	"aoc24/utils"
	"fmt"
	"math"
	"slices"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines, err := utils.GetInput("day1")
	if err != nil {
		panic(err)
	}
	lefts, rights := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		ints, err := utils.AtoiSplit(line, "   ")
		if err != nil {
			fmt.Errorf("invalid input")
		}
		lefts[i], rights[i] = ints[0], ints[1]
	}
	slices.Sort(lefts)
	slices.Sort(rights)
	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += int(math.Abs(float64(lefts[i]) - float64(rights[i])))
	}
	fmt.Println("part1", sum)
}

func part2() {
	lines, err := utils.GetInput("day1")
	if err != nil {
		panic(err)
	}
	lefts := make([]int, len(lines))
	occ := make(map[int]int)
	for i, line := range lines {
		ints, err := utils.AtoiSplit(line, "   ")
		if err != nil {
			fmt.Errorf("invalid input")
		}
		lefts[i] = ints[0]
		occ[ints[1]]++
	}
	sum := 0
	for i := range lefts {
		sum += lefts[i] * occ[lefts[i]]
	}
	fmt.Println("part2", sum)
}
