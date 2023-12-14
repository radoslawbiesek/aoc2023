package main

import (
	"fmt"
	"slices"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func moveVertical(grid *utils.Grid[string], direction int) {
	_, height := grid.GetDimensions()
	points := grid.GetAllPoints()
	if direction > 0 {
		slices.Reverse(points)
	}
	for _, point := range points {
		value := grid.GetValue(point)
		if value != "O" {
			continue
		}

		i := 0
		for {
			curr := utils.Point{X: point.X, Y: point.Y + i*direction}
			next := utils.Point{X: point.X, Y: curr.Y + direction}
			if next.Y < 0 || next.Y >= height || grid.GetValue(next) != "." {
				break
			}
			grid.SetValue(curr, ".")
			grid.SetValue(next, "O")
			i++
		}
	}
}

func moveHorizontal(grid *utils.Grid[string], direction int) {
	width, _ := grid.GetDimensions()
	points := grid.GetAllPoints()
	if direction > 0 {
		slices.Reverse(points)
	}
	for _, point := range points {
		value := grid.GetValue(point)
		if value != "O" {
			continue
		}

		i := 0
		for {
			curr := utils.Point{X: point.X + i*direction, Y: point.Y}
			next := utils.Point{X: curr.X + direction, Y: point.Y}
			if next.X < 0 || next.X >= width || grid.GetValue(next) != "." {
				break
			}
			grid.SetValue(curr, ".")
			grid.SetValue(next, "O")
			i++
		}
	}
}

func calculateTotal(grid *utils.Grid[string]) (total int) {
	_, height := grid.GetDimensions()
	for _, point := range grid.GetAllPoints() {
		if grid.GetValue(point) == "O" {
			total += height - point.Y
		}
	}
	return
}

func part1(path string) int {
	content := utils.GetContent(path)
	grid := utils.NewStrGrid(content)

	moveVertical(grid, -1)

	return calculateTotal(grid)
}

const LOOPS = 1_000_000_000

func part2(path string) int {
	content := utils.GetContent(path)
	grid := utils.NewStrGrid(content)
	gridMap := map[string][]int{}

	cycleDetected := false
	for i := 0; i < LOOPS; i++ {
		moveVertical(grid, -1)   // N
		moveHorizontal(grid, -1) // W
		moveVertical(grid, +1)   // S
		moveHorizontal(grid, 1)  // E

		if cycleDetected {
			continue
		}

		key := fmt.Sprintf("%v", grid)
		indexes := gridMap[fmt.Sprintf("%v", grid)]
		indexes = append(indexes, i)
		gridMap[key] = indexes

		if len(indexes) >= 3 && indexes[2]-indexes[1] == indexes[1]-indexes[0] {
			cycleDetected = true
		}

		if cycleDetected {
			cycle := indexes[1] - indexes[0]
			loopsLeft := LOOPS - i
			i = LOOPS - loopsLeft%cycle
		}
	}

	return calculateTotal(grid)
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
