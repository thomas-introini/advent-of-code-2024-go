package main

import (
	"aoc24/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	input, err := utils.GetInputAsString("day3")
	if err != nil {
		panic(err)
	}
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		sum += first * second
	}
	fmt.Println("part1", sum)

}

func part2() {
	input, err := utils.GetInputAsString("day3")
	if err != nil {
		panic(err)
	}
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	fmt.Println(matches)
	sum := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		} else if match[0] == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			first, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			sum += first * second
		}
	}
	fmt.Println("part2", sum)

}
