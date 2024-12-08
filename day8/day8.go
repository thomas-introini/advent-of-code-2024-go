package main

import (
	"aoc24/utils"
	"fmt"
)

func main() {
	grid := utils.MustGetInput("day8")
	antennas := make(map[byte][][2]int)
	for i := 0; i < len(grid); i++ {
		// fmt.Println(i, grid[i])
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], [2]int{i, j})
			}
		}
	}
	part1(grid, antennas)
	part2(grid, antennas)
}

func part1(grid []string, antennas map[byte][][2]int) {
	keys := utils.Keys(antennas)
	uniqueAntinodes := make(map[[2]int]bool)
	for _, k := range keys {
		a := antennas[k]
		// fmt.Println("Antenna", string(k), a)
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				a1, a2 := a[i], a[j]
				antinodes := GetAntinodes(grid, a1, a2)
				// fmt.Println("Antinodes", a1, a2, antinodes)
				for _, an := range antinodes {
					uniqueAntinodes[an] = true
				}
			}
		}
	}
	/* for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				fmt.Print(i, " ")
			}
			if grid[i][j] == '.' && uniqueAntinodes[[2]int{i, j}] {
				fmt.Print("#")
			} else {
				fmt.Print(string(grid[i][j]))
			}
		}
		fmt.Println()
	} */
	fmt.Println("part1", len(uniqueAntinodes))

}

func part2(grid []string, antennas map[byte][][2]int) {
	keys := utils.Keys(antennas)
	uniqueAntinodes := make(map[[2]int]bool)
	for _, k := range keys {
		a := antennas[k]
		// fmt.Println("Antenna", string(k), a)
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				a1, a2 := a[i], a[j]

				uniqueAntinodesForAntenna := make(map[[2]int]bool)
				antinodes := GetAntinodesHarmonic(grid, a1, a2, uniqueAntinodesForAntenna)
				uniqueAntinodes[a1] = true
				uniqueAntinodes[a2] = true
				// fmt.Println("Antinodes", a1, a2, antinodes)
				for _, an := range antinodes {
					uniqueAntinodes[an] = true
				}
			}
		}
	}
	/* for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				fmt.Print(i, " ")
			}
			if grid[i][j] == '.' && uniqueAntinodes[[2]int{i, j}] {
				fmt.Print("#")
			} else {
				fmt.Print(string(grid[i][j]))
			}
		}
		fmt.Println()
	} */
	fmt.Println("part2", len(uniqueAntinodes))

}

func GetAntinodes(grid []string, a1, a2 [2]int) (antinodes [][2]int) {
	antinode := GetAntinode(grid, a1, a2)
	if antinode != [2]int{-1, -1} {
		antinodes = append(antinodes, antinode)
	}
	antinode = GetAntinode(grid, a2, a1)
	if antinode != [2]int{-1, -1} {
		antinodes = append(antinodes, antinode)
	}
	return
}

func GetAntinodesHarmonic(grid []string, a1, a2 [2]int, uniqueAntinodes map[[2]int]bool) (antinodes [][2]int) {
	antinode := GetAntinode(grid, a1, a2)
	if antinode != [2]int{-1, -1} && !uniqueAntinodes[antinode] {
		antinodes = append(antinodes, antinode)
		uniqueAntinodes[antinode] = true
		harmonicAntinodes := GetAntinodesHarmonic(grid, a1, antinode, uniqueAntinodes)
		// fmt.Println("HAntinodes", a1, antinode, harmonicAntinodes)
		for _, h := range harmonicAntinodes {
			antinodes = append(antinodes, h)
		}
	}
	antinode = GetAntinode(grid, a2, a1)
	if antinode != [2]int{-1, -1} && !uniqueAntinodes[antinode] {
		antinodes = append(antinodes, antinode)
		uniqueAntinodes[antinode] = true
		harmonicAntinodes := GetAntinodesHarmonic(grid, a2, antinode, uniqueAntinodes)
		// fmt.Println("HAntinodes", a2, antinode, harmonicAntinodes)
		for _, h := range harmonicAntinodes {
			antinodes = append(antinodes, h)
		}
	}
	return
}

func GetAntinode(grid []string, a1, a2 [2]int) (antinode [2]int) {
	antinode = [2]int{-1, -1}
	deltaX := a1[0] - a2[0]
	deltaY := a1[1] - a2[1]

	possibleAntinode := [2]int{a1[0] + deltaX, a1[1] + deltaY}

	if possibleAntinode[0] >= 0 && possibleAntinode[0] < len(grid[0]) && possibleAntinode[1] >= 0 && possibleAntinode[1] < len(grid) {
		antinode = possibleAntinode
	}
	return
}
