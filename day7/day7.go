package main

import (
	"aoc24/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.MustGetInput("day7")
	fmt.Println(lines)
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ":")
		result, expr := split[0], split[1]
		expr = strings.TrimLeft(expr, " ")
		resultInt, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		ints, err := utils.AtoiSplit(expr, " ")
		if err != nil {
			panic(err)
		}
		slices.Reverse(ints)
		evaluations := evalutate(ints)
		// fmt.Println(resultInt, ints, evaluations)
		if Contains(evaluations, resultInt) {
			sum += resultInt
			// fmt.Println(sum)
		}
	}
	fmt.Println(sum)
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func evalutate(ints []int) []int {
	if len(ints) == 1 {
		return ints[0:1]
	} else {
		rest := evalutate(ints[1:])
		ret := make([]int, 0)
		for _, v := range rest {
			ret = append(ret, ints[0]+v)
			ret = append(ret, ints[0]*v)

			// For Part 1 comment, the lines below
			int, err := strconv.Atoi(fmt.Sprintf("%d%d", v, ints[0]))
			if err != nil {
				panic(err)
			}
			ret = append(ret, int)
		}
		return ret
	}
}
