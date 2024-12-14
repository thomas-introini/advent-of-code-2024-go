package main

import (
	"aoc24/utils"
	"fmt"
	"strconv"
)

var (
	directions = [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func main() {
	lines := utils.MustGetInput("day10")
	for _, line := range lines {
		fmt.Println(line)
	}
	part1(lines)
	part2(lines)
}

func part1(grid []string) {
	trailheads := make([][2]int, 0)
	for i, line := range grid {
		for j, c := range line {
			if c == '0' {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}
	scores := make(map[[2]int]int)
	for _, trailhead := range trailheads {
		visited := make(map[[2]int]bool)
		visiting := make([][2]int, 1)
		visiting[0] = trailhead
		for len(visiting) > 0 {
			for _, current := range visiting {
				if !visited[current] {
					if grid[current[0]][current[1]] == '9' {
						scores[trailhead] += 1
					}
					for _, direction := range directions {
						if next, valid := IsValidMove(grid, current, direction); valid {
							visiting = append(visiting, next)
						}
					}
				}
				visited[current] = true
			}
			newVisiting := make([][2]int, 0)
			for _, current := range visiting {
				if !visited[current] {
					newVisiting = append(newVisiting, current)
				}
			}
			visiting = newVisiting
		}
	}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Println("part1", sum)
}

func part2(grid []string) {
	trailheads := make([][2]int, 0)
	for i, line := range grid {
		for j, c := range line {
			if c == '0' {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}
	fmt.Println(trailheads)
	scores := make(map[[2]int]int)
	ratings := make(map[[2]int][][2]int)
	for _, trailhead := range trailheads {
		visited := make(map[[2]int]bool)
		visiting := make([][2]int, 1)
		visiting[0] = trailhead
		for len(visiting) > 0 {
			newVisiting := make([][2]int, 0)
			for _, current := range visiting {
				if grid[current[0]][current[1]] == '9' {
					scores[trailhead] += 1
					ratings[trailhead] = append(ratings[trailhead], current)
					visited[current] = true
				}
				for _, direction := range directions {
					if next, valid := IsValidMove(grid, current, direction); valid {
						newVisiting = append(newVisiting, next)
					}
				}
			}
			fmt.Println(visiting)
			visiting = newVisiting
		}
	}
	sum := 0
	for _, rating := range ratings {
		sum += len(rating)
	}
	fmt.Println("part2", sum)
}

func IsValidMove(grid []string, current [2]int, direction [2]int) ([2]int, bool) {
	if current[0]+direction[0] < 0 || current[0]+direction[0] >= len(grid) || current[1]+direction[1] < 0 || current[1]+direction[1] >= len(grid[0]) {
		return [2]int{}, false
	}

	newPos := [2]int{current[0] + direction[0], current[1] + direction[1]}
	currentInt, err := strconv.Atoi(string(grid[current[0]][current[1]]))
	if err != nil {
		panic(err)
	}
	nextInt, err := strconv.Atoi(string(grid[newPos[0]][newPos[1]]))
	if err != nil {
		panic(err)
	}

	return newPos, nextInt-currentInt == 1
}
