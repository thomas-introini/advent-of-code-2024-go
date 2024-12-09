package main

import (
	"aoc24/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.MustGetInputAsString("day9")
	input = strings.TrimRight(input, "\n")
	part1(input)
	part2(input)
}

func part1(input string) {
	sum := 0
	for i := 0; i < len(input); i++ {
		if string(input[i]) != "\n" {
			int, err := strconv.Atoi(string(input[i]))
			if err != nil {
				panic(err)
			}
			sum += int
		}
	}
	fs := make([]int, sum)
	empty := false
	fileId := 0
	fsI := 0
	for i := 0; i < len(input); i++ {
		int, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}
		for j := fsI; j < fsI+int; j++ {
			if empty {
				fs[j] = -1
			} else {
				fs[j] = fileId
			}
		}
		if !empty {
			fileId++
		}
		fsI += int
		empty = !empty
	}
	emptyLeft, emptyRight := 0, len(fs)-1
	for emptyLeft < emptyRight {
		if fs[emptyLeft] != -1 {
			emptyLeft++
		}
		if fs[emptyRight] == -1 {
			emptyRight--
		}
		if fs[emptyLeft] == -1 && fs[emptyRight] != -1 {
			fs[emptyLeft] = fs[emptyRight]
			fs[emptyRight] = -1
		}
	}
	fmt.Println("part1", Checksum(fs))
}

func part2(input string) {
	sum := 0
	for i := 0; i < len(input); i++ {
		if string(input[i]) != "\n" {
			int, err := strconv.Atoi(string(input[i]))
			if err != nil {
				panic(err)
			}
			sum += int
		}
	}
	fs := make([]int, sum)
	empty := false
	fileId := 0
	fsI := 0
	fileMap := make(map[int][2]int)
	for i := 0; i < len(input); i++ {
		block, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}
		for j := fsI; j < fsI+block; j++ {
			if empty {
				fs[j] = -1
			} else {
				fs[j] = fileId
			}
		}
		if !empty {
			fileMap[fileId] = [2]int{fsI, fsI + block - 1}
			fileId++
		}
		fsI += block
		empty = !empty
	}
	for fileId := 9999; fileId >= 0; fileId-- {
		filePos, ok := fileMap[fileId]
		if !ok {
			continue
		}
		startIdx, endIdx := filePos[0], filePos[1]
		fileSize := endIdx - startIdx
		emptySize := 0
		for i := 0; i <= startIdx; i++ {
			if fs[i] == -1 {
				emptySize++
			} else if fs[i] != -1 && emptySize > 0 {
				if emptySize >= fileSize+1 {
					startEmpty := (i - emptySize)
					for idx := startEmpty; idx <= startEmpty+fileSize; idx++ {
						fs[idx] = fileId
					}
					for idx := startIdx; idx <= endIdx; idx++ {
						fs[idx] = -1
					}
					break
				}
				emptySize = 0
			}
		}
	}

	fmt.Println("part2", Checksum(fs))
}

func Checksum(fs []int) (checksum int) {
	for i := 0; i < len(fs); i++ {
		if fs[i] == -1 {
			continue
		}
		checksum += fs[i] * i
	}
	return
}

func PrettyPrintFs(fs []int) {
	for i := 0; i < len(fs); i++ {
		if fs[i] == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(fs[i])
		}
	}
	fmt.Println()
}
