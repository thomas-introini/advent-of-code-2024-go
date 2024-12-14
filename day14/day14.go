package main

import (
	"aoc24/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines := utils.MustGetInput("day14")
	columns := 101
	rows := 103
	times := 100
	part1(times, columns, rows, lines)
	part2(times, columns, rows, lines)
}

type Robot struct {
	Pos utils.Point
	Vel utils.Point
}

func part1(times, columns, rows int, input []string) {
	inputRegex, err := regexp.Compile(`p\=(\d+),(\d+) v\=(-?\d+),(-?\d+)`)
	if err != nil {
		panic(err)
	}
	fmt.Println("rows", rows, "columns", columns)
	robots := make([]Robot, len(input))
	for i, line := range input {
		matches := inputRegex.FindAllStringSubmatch(line, -1)
		posRow, _ := strconv.Atoi(matches[0][2])
		posCol, _ := strconv.Atoi(matches[0][1])
		velRow, _ := strconv.Atoi(matches[0][4])
		velCol, _ := strconv.Atoi(matches[0][3])
		robots[i] = Robot{
			Pos: utils.Point{Row: posRow, Col: posCol},
			Vel: utils.Point{Row: velRow, Col: velCol},
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	middleRow := (rows - 1) / 2
	middleCol := (columns - 1) / 2
	for _, robot := range robots {
		finalPosCol := mod((robot.Vel.Col*times)+robot.Pos.Col, columns)
		finalPosRow := mod((robot.Vel.Row*times)+robot.Pos.Row, rows)
		// fmt.Println(robot.Pos, robot.Vel, finalPosCol, finalPosRow)
		if finalPosRow < middleRow {
			if finalPosCol < middleCol {
				q1++
			} else if finalPosCol > middleCol {
				q2++
			}
		} else if finalPosRow > middleRow {
			if finalPosCol < middleCol {
				q3++
			} else if finalPosCol > middleCol {
				q4++
			}
		}
	}
	fmt.Println("part1", q1*q2*q3*q4)
}

func part2(times, columns, rows int, input []string) {
	inputRegex, err := regexp.Compile(`p\=(\d+),(\d+) v\=(-?\d+),(-?\d+)`)
	if err != nil {
		panic(err)
	}
	fmt.Println("rows", rows, "columns", columns)
	robots := make([]Robot, len(input))
	for i, line := range input {
		matches := inputRegex.FindAllStringSubmatch(line, -1)
		posRow, _ := strconv.Atoi(matches[0][2])
		posCol, _ := strconv.Atoi(matches[0][1])
		velRow, _ := strconv.Atoi(matches[0][4])
		velCol, _ := strconv.Atoi(matches[0][3])
		robots[i] = Robot{
			Pos: utils.Point{Row: posRow, Col: posCol},
			Vel: utils.Point{Row: velRow, Col: velCol},
		}
	}

	start := 7412
	for i := start; i <= start; i++ {
		fmt.Print("\033[H\033[2J")

		fmt.Print(strings.Repeat("#", columns))
		fmt.Println()
		positions := make(map[utils.Point]bool)
		for _, robot := range robots {
			finalPosCol := mod((robot.Vel.Col*i)+robot.Pos.Col, columns)
			finalPosRow := mod((robot.Vel.Row*i)+robot.Pos.Row, rows)
			positions[utils.Point{Row: finalPosRow, Col: finalPosCol}] = true
			// fmt.Println(robot.Pos, robot.Vel, finalPosCol, finalPosRow)
		}
		for r := 0; r < rows; r++ {
			for c := 0; c < columns; c++ {
				if _, ok := positions[utils.Point{Row: r, Col: c}]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println(i)
		time.Sleep(time.Millisecond * 400)
		fmt.Println()
	}
	fmt.Println("part2", start)
}

func mod(a, b int) int {
	return (a%b + b) % b
}
