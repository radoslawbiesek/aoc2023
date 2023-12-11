package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func columnContains(grid utils.Grid[string], index int) bool {
	for _, row := range grid {
		if row[index] == "#" {
			return true
		}
	}
	return false
}

func getInput(path string, expansion int) (galaxies []utils.Point) {
	content := utils.GetContent(path)
	grid := *utils.NewStrGrid(content)

	x, y := 0, 0
	for _, row := range grid {
		for colIndex, col := range row {
			if col == "#" {
				galaxies = append(galaxies, utils.Point{X: x, Y: y})
			}

			if !columnContains(grid, colIndex) {
				x += expansion
			} else {
				x++
			}
		}
		x = 0
		if !slices.Contains(row, "#") {
			y += expansion
		} else {
			y++
		}
	}

	return
}

func generateKey(i1, i2 int) string {
	if i1 < i2 {
		return fmt.Sprintf("%d-%d", i1, i2)
	}
	return fmt.Sprintf("%d-%d", i2, i1)
}

func calculateSteps(p1, p2 utils.Point) int {
	return int(math.Abs(float64(p1.X)-float64(p2.X))) + int(math.Abs(float64(p1.Y)-float64(p2.Y)))
}

func solution(path string, expansion int) (total int) {
	galaxies := getInput(path, expansion)
	calculated := map[string]bool{}
	for i1, p1 := range galaxies {
		for i2, p2 := range galaxies {
			if p1 == p2 {
				continue
			}
			key := generateKey(i1, i2)
			if calculated[key] {
				continue
			}
			total += calculateSteps(p1, p2)
			calculated[key] = true
		}
	}
	return
}

func part1(path string) int {
	return solution(path, 2)
}

func part2(path string) int {
	if path == "./test-input.txt" {
		return solution(path, 100)
	}
	return solution(path, 1_000_000)
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
