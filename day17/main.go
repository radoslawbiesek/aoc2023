package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func last3(slice []utils.Direction, nextDirection utils.Direction) []utils.Direction {
	if len(slice) < 3 {
		return append(slice, nextDirection)
	}
	return append(slice[1:], nextDirection)
}

func checkLast3(slice []utils.Direction, dir utils.Direction) bool {
	if len(slice) < 3 {
		return true
	}

	return slices.ContainsFunc(slice, func(el utils.Direction) bool {
		return el != dir
	})
}

func copySeen(seen map[utils.Point]bool) map[utils.Point]bool {
	newSeen := map[utils.Point]bool{}
	for k, v := range seen {
		newSeen[k] = v
	}
	return newSeen
}

func walk(grid *utils.Grid[int], seen map[utils.Point]bool, currPoint, endPoint utils.Point, directions []utils.Direction, currDistance int, minDistance *int) {
	newSeen := copySeen(seen)
	newSeen[currPoint] = true

	newDistance := currDistance + grid.GetValue(currPoint)
	width, height := grid.GetDimensions()

	if newDistance >= *minDistance {
		return
	}
	if currPoint == endPoint {
		if newDistance < *minDistance {
			*minDistance = newDistance
		}
		return
	}

	for _, nextDirection := range utils.Directions4 {
		nextPoint := utils.Point{X: currPoint.X + nextDirection.X, Y: currPoint.Y + nextDirection.Y}

		if newSeen[nextPoint] {
			continue
		}

		if nextPoint.X < 0 || nextPoint.X >= width || nextPoint.Y < 0 || nextPoint.Y >= height {
			continue
		}

		if !checkLast3(directions, nextDirection) {
			continue
		}

		walk(grid, newSeen, nextPoint, endPoint, last3(directions, nextDirection), newDistance, minDistance)
	}

}

func part1(path string) (minDistance int) {
	minDistance = math.MaxInt
	content := utils.GetContent(path)
	grid := utils.NewIntGrid(content)
	width, height := grid.GetDimensions()

	startPoint := utils.Point{X: 0, Y: 0}
	endPoint := utils.Point{X: width - 1, Y: height - 1}

	walk(grid, make(map[utils.Point]bool), startPoint, endPoint, []utils.Direction{}, -grid.GetValue(startPoint), &minDistance)

	return
}

func part2(path string) (total int) {
	return
}

func main() {
	// fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	// fmt.Printf("Part 2: %d\n", part2("./test-input.txt"))
	// fmt.Println("")
	// fmt.Println("Input: ")
	// fmt.Printf("Part 1: %d\n", part1("./input.txt"))
	// fmt.Printf("Part 2: %d\n", part2("./input.txt"))
}
