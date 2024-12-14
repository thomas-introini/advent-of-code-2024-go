package main

import (
	"aoc24/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexpA     *regexp.Regexp
	regexpB     *regexp.Regexp
	regexpPrize *regexp.Regexp
)

func main() {
	lines := utils.MustGetInputAsString("day13")
	var err error
	regexpA, err = regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)`)
	if err != nil {
		panic(err)
	}
	regexpB, err = regexp.Compile(`Button B: X\+(\d+), Y\+(\d+)`)
	if err != nil {
		panic(err)
	}
	regexpPrize, err = regexp.Compile(`Prize: X=(\d+), Y=(\d+)`)
	if err != nil {
		panic(err)
	}

	part1(strings.Split(lines, "\n\n"))
	part2(strings.Split(lines, "\n\n"))

}

func part1(lines []string) {
	sum := 0
	for i, line := range lines {
		split := strings.Split(line, "\n")
		aMatches := regexpA.FindAllStringSubmatch(split[0], -1)
		bMatches := regexpB.FindAllStringSubmatch(split[1], -1)
		prizeMatches := regexpPrize.FindAllStringSubmatch(split[2], -1)
		aX, _ := strconv.Atoi(aMatches[0][1])
		aY, _ := strconv.Atoi(aMatches[0][2])
		bX, _ := strconv.Atoi(bMatches[0][1])
		bY, _ := strconv.Atoi(bMatches[0][2])
		prizeX, _ := strconv.Atoi(prizeMatches[0][1])
		prizeY, _ := strconv.Atoi(prizeMatches[0][2])

		a, b, err := SolveLinearEquations(aX, bX, prizeX, aY, bY, prizeY)
		if err == nil && a >= 0 && b >= 0 && a <= 100 && b <= 100 {
			fmt.Println("Claw", i+1, a, b)
			sum += (3 * a) + b
		} else {
			fmt.Println("No solution", i+1, a, b, err)
		}

	}
	fmt.Println("part1", sum)

}

func part2(lines []string) {
	sum := 0
	for i, line := range lines {
		split := strings.Split(line, "\n")
		aMatches := regexpA.FindAllStringSubmatch(split[0], -1)
		bMatches := regexpB.FindAllStringSubmatch(split[1], -1)
		prizeMatches := regexpPrize.FindAllStringSubmatch(split[2], -1)
		aX, _ := strconv.Atoi(aMatches[0][1])
		aY, _ := strconv.Atoi(aMatches[0][2])
		bX, _ := strconv.Atoi(bMatches[0][1])
		bY, _ := strconv.Atoi(bMatches[0][2])
		prizeX, _ := strconv.Atoi(prizeMatches[0][1])
		prizeY, _ := strconv.Atoi(prizeMatches[0][2])

		a, b, err := SolveLinearEquations(aX, bX, prizeX+10000000000000, aY, bY, prizeY+10000000000000)
		if err == nil && a >= 0 && b >= 0 {
			fmt.Println("Claw", i+1, a, b)
			sum += (3 * a) + b
		} else {
			fmt.Println("No solution", i+1, a, b, err)
		}

	}
	fmt.Println("part2", sum)

}

func SolveLinearEquations(aX, bX, prizeX, aY, bY, prizeY int) (int, int, error) {
	det := aX*bY - bX*aY
	if det == 0 {
		return 0, 0, fmt.Errorf("no solution")
	}
	detA := prizeX*bY - bX*prizeY
	detB := aX*prizeY - prizeX*aY

	if detA%det != 0 || detB%det != 0 {
		return 0, 0, fmt.Errorf("no integer solution")
	}

	a := detA / det
	b := detB / det

	return a, b, nil
}
