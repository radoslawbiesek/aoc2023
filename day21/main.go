package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

type infiniteGrid [][]string

func newInfiniteGrid(path string) infiniteGrid {
	grid := infiniteGrid{}
	for _, line := range utils.GetLines(path, "\n") {
		row := []string{}
		for _, char := range strings.Split(line, "") {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	return grid
}

func (g *infiniteGrid) getDimensions() (width, height int) {
	width, height = len((*g)[0]), len(*g)
	return
}

func findIdx(curr, dim int) int {
	idx := curr
	for idx < 0 {
		idx += dim
	}
	return idx
}

func (g *infiniteGrid) getValue(p utils.Point) string {
	width, height := g.getDimensions()
	var x, y int
	if p.X < 0 {
		x = findIdx(p.X, width)
	} else {
		x = p.X % width
	}
	if p.Y < 0 {
		y = findIdx(p.Y, height)
	} else {
		y = p.Y % height
	}
	return (*g)[y][x]
}

func (g *infiniteGrid) findValue(value string) utils.Point {
	for y, row := range *g {
		for x, char := range row {
			if char == value {
				return utils.Point{X: x, Y: y}
			}
		}
	}
	panic("not found")
}

func (g *infiniteGrid) get4Neighbors(p utils.Point) (points []utils.Point) {
	for _, dir := range utils.Directions4 {
		points = append(points, utils.Point{X: p.X + dir.X, Y: p.Y + dir.Y})
	}
	return
}

func solution(path string, steps int) int {
	grid := newInfiniteGrid(path)
	startPoint := grid.findValue("S")

	currentSteps := &map[utils.Point]bool{}
	nextSteps := &map[utils.Point]bool{}
	(*currentSteps)[startPoint] = true

	for i := 0; i < steps; i++ {
		for curr := range *currentSteps {
			for _, neighbor := range grid.get4Neighbors(curr) {
				if grid.getValue(neighbor) == "#" {
					continue
				}
				(*nextSteps)[neighbor] = true
			}
		}

		currentSteps = nextSteps
		nextSteps = &map[utils.Point]bool{}
	}

	return len(*currentSteps)
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", solution("./test-input.txt", 6))
	// fmt.Printf("Part 2: %d\n", solution("./test-input.txt", 5000))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", solution("./input.txt", 64))
	// fmt.Printf("Part 2: %d\n", solution("./input.txt", 26501365))
}
