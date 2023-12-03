package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func isEmpty(char string) bool {
	return char == "."
}

func isInt(char string) bool {
	_, err := strconv.Atoi(char)

	return err == nil
}

func isGear(char string) bool {
	return char == "*"
}

func isAdjacentToSymbol(grid utils.Grid[string], lineIdx, idxStart, idxEnd int) bool {
	for i := idxStart; i <= idxEnd; i++ {
		currPoint := utils.Point{X: i, Y: lineIdx}
		for _, neighbor := range grid.Get8Neighbors(currPoint) {
			value := grid.GetValue(neighbor)
			if !isEmpty(value) && !isInt(value) {
				return true
			}
		}
	}

	return false
}

func findAdjacentGearSymbol(grid utils.Grid[string], lineIdx, idxStart, idxEnd int) *utils.Point {
	for i := idxStart; i <= idxEnd; i++ {
		currPoint := utils.Point{X: i, Y: lineIdx}
		for _, neighbor := range grid.Get8Neighbors(currPoint) {
			value := grid.GetValue(neighbor)
			if isGear(value) {
				return &neighbor
			}
		}
	}

	return nil
}

func getNumber(line []string, idxStart, idxEnd int) int {
	return utils.ParseInt(strings.Join(line[idxStart:idxEnd+1], ""))

}

func part1(path string) (total int) {
	content := utils.GetContent(path)
	grid := *utils.NewStrGrid(content)

	for lineIdx, line := range grid {
		idxStart := -1
		idxEnd := -1

		for charIdx, char := range line {
			if isInt(char) {
				if idxStart == -1 {
					idxStart = charIdx
				}
				idxEnd = charIdx
			}

			if idxStart != -1 && (!isInt(char) || charIdx == len(line)-1) {
				if isAdjacentToSymbol(grid, lineIdx, idxStart, idxEnd) {
					num := getNumber(line, idxStart, idxEnd)
					total += num
				}
				idxStart = -1
				idxEnd = -1
			}
		}
	}

	return
}

func part2(path string) (total int) {
	content := utils.GetContent(path)
	grid := *utils.NewStrGrid(content)
	gearMap := map[utils.Point][]int{}

	for lineIdx, line := range grid {
		idxStart := -1
		idxEnd := -1
		for charIdx, char := range line {
			if isInt(char) {
				if idxStart == -1 {
					idxStart = charIdx
				}
				idxEnd = charIdx
			}

			if idxStart != -1 && (!isInt(char) || charIdx == len(line)-1) {
				gearPoint := findAdjacentGearSymbol(grid, lineIdx, idxStart, idxEnd)
				if gearPoint != nil {
					num := getNumber(line, idxStart, idxEnd)
					gearPoints, ok := gearMap[*gearPoint]
					if ok {
						gearMap[*gearPoint] = append(gearPoints, num)
					} else {
						gearMap[*gearPoint] = []int{num}
					}
				}
				idxStart = -1
				idxEnd = -1
			}
		}
	}

	for _, numbers := range gearMap {
		if len(numbers) == 2 {
			total += numbers[0] * numbers[1]
		}
	}

	return
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./test-input.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./input.txt"))
}
