package main

import (
	"aoc24/utils"
	"fmt"
)

var (
	directions = [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.MustGetInput("day6")
	currDir := directions[0]
	var currPos [2]int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '^' {
				currPos = [2]int{i, j}
				break
			}
		}
	}
	visited := make(map[[2]int][2]int)
	visited[currPos] = currDir
	for i := 0; ; {
		nextPos := [2]int{currPos[0] + currDir[0], currPos[1] + currDir[1]}
		if nextPos[0] < 0 || nextPos[0] >= len(lines) || nextPos[1] < 0 || nextPos[1] >= len(lines[0]) {
			break
		}
		if lines[nextPos[0]][nextPos[1]] == '#' {
			i++
			currDir = directions[i%4]
			nextPos = [2]int{currPos[0] + currDir[0], currPos[1] + currDir[1]}
			visited[nextPos] = currDir
			continue
		}
		visited[nextPos] = currDir
		currPos = nextPos
	}
	fmt.Println("part1", len(visited))
}

func part2() {
	lines := utils.MustGetInput("day6")
	var currPos [2]int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '^' {
				currPos = [2]int{i, j}
				break
			}
		}
	}
	obstacles := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '#' {
				if simulateGuardWithObstacle(lines, currPos, directions[0], [2]int{i, j}) {
					obstacles++
				}
			}
		}
	}
	fmt.Println("part2", obstacles)
}

func simulateGuardWithObstacle(lines []string, currPos [2]int, currDir [2]int, obstaclePos [2]int) bool {
	loopDir := currDir
	currLoopPos := currPos
	visitedLoop := make(map[string]bool)
	visitedLoop[fmt.Sprintf("%s,%s,%s,%s", currLoopPos[0], currLoopPos[1], loopDir[0], loopDir[1])] = true
	for i := 0; ; {
		nextLoopPos := [2]int{currLoopPos[0] + loopDir[0], currLoopPos[1] + loopDir[1]}
		if nextLoopPos[0] < 0 || nextLoopPos[0] >= len(lines) || nextLoopPos[1] < 0 || nextLoopPos[1] >= len(lines[0]) {
			return false
		}
		if lines[nextLoopPos[0]][nextLoopPos[1]] == '#' || (obstaclePos[0] == nextLoopPos[0] && obstaclePos[1] == nextLoopPos[1]) {
			i++
			loopDir = directions[i%4]
			continue
		}
		if visitedLoop[fmt.Sprintf("%s,%s,%s,%s", nextLoopPos[0], nextLoopPos[1], loopDir[0], loopDir[1])] {
			return true
		}
		visitedLoop[fmt.Sprintf("%s,%s,%s,%s", nextLoopPos[0], nextLoopPos[1], loopDir[0], loopDir[1])] = true
		currLoopPos = nextLoopPos
	}
}
