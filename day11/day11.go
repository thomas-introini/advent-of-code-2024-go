package main

import (
	"aoc24/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.MustGetInputAsString("day11")
	part1(strings.TrimRight(input, "\n"))
	part2(strings.TrimRight(input, "\n"))
}

func part1(stones string) {
	fmt.Println(stones)
	arrangement := strings.Split(stones, " ")
	stoneMap := make(map[string]int)
	for _, stone := range arrangement {
		stoneMap[stone] += 1
	}
	for i := 0; i < 25; i++ {
		stoneMap = blink(stoneMap)
		fmt.Println("blink", i+1)
	}
	sum := 0
	for _, v := range stoneMap {
		sum += v
	}
	fmt.Println("part1", sum)
}

func part2(stones string) {
	fmt.Println(stones)
	arrangement := strings.Split(stones, " ")
	stoneMap := make(map[string]int)
	for _, stone := range arrangement {
		stoneMap[stone] += 1
	}
	for i := 0; i < 75; i++ {
		stoneMap = blink(stoneMap)
		fmt.Println("blink", i+1)
	}
	sum := 0
	for _, v := range stoneMap {
		sum += v
	}
	fmt.Println("part1", sum)
}

func blink(stones map[string]int) (newArrangement map[string]int) {
	newArrangement = make(map[string]int)
	for stone, count := range stones {
		if stone == "0" {
			newArrangement["1"] += count
		} else if len(stone)%2 == 0 {
			left := stone[:len(stone)/2]
			right, err := strconv.Atoi(stone[len(stone)/2:])
			if err != nil {
				panic(err)
			}
			newArrangement[left] += count
			newArrangement[fmt.Sprintf("%d", right)] += count
		} else {
			stone, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			newArrangement[fmt.Sprintf("%d", stone*2024)] += count
		}
	}
	return
}
