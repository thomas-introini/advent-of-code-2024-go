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
	lines := utils.MustGetInput("day4")
	count := 0
	for i, line := range lines {
		// fmt.Println(line)
		for j := 0; j < len(line); j++ {
			if line[j] == 'X' || line[j] == 'S' {
				/* if j >= 3 {
					str := line[j-3 : j+1]
					if str == "SAMX" {
						count++
					}
				} */
				if j < len(line)-3 {
					str := line[j : j+4]
					if str == "XMAS" || str == "SAMX" {
						count++
					}
				}

				if i < len(lines)-3 && j >= 3 {
					str := fmt.Sprintf(
						"%s%s%s%s",
						string(lines[i][j]),
						string(lines[i+1][j-1]),
						string(lines[i+2][j-2]),
						string(lines[i+3][j-3]),
					)
					if str == "XMAS" || str == "SAMX" {
						// fmt.Println(str)
						count++
					}
				}

				if i < len(lines)-3 && j < len(line)-3 {
					str := fmt.Sprintf(
						"%s%s%s%s",
						string(lines[i][j]),
						string(lines[i+1][j+1]),
						string(lines[i+2][j+2]),
						string(lines[i+3][j+3]),
					)
					if str == "XMAS" || str == "SAMX" {
						// fmt.Println(str)
						count++
					}
				}
				if i < len(lines)-3 {
					str := fmt.Sprintf(
						"%s%s%s%s",
						string(lines[i][j]),
						string(lines[i+1][j]),
						string(lines[i+2][j]),
						string(lines[i+3][j]),
					)
					if str == "XMAS" || str == "SAMX" {
						// fmt.Println(str)
						count++
					}
				}
			}
		}
	}
	fmt.Println("part1", count)
}

func part2() {
	lines := utils.MustGetInput("day4")
	count := 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if i < len(lines)-2 && j < len(line)-2 {
				left := fmt.Sprintf(
					"%s%s%s",
					string(lines[i][j]),
					string(lines[i+1][j+1]),
					string(lines[i+2][j+2]),
				)
				fmt.Println("left", left)
				if left != "MAS" && left != "SAM" {
					continue
				}
				right := fmt.Sprintf(
					"%s%s%s",
					string(lines[i][j+2]),
					string(lines[i+1][j+1]),
					string(lines[i+2][j]),
				)
				fmt.Println("right", left)
				if right != "MAS" && right != "SAM" {
					continue
				}
				count++
			}
		}
	}
	fmt.Println("part2", count)

}
