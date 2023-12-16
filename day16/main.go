package main

import (
	"fmt"
	"slices"

	"github.com/radoslawbiesek/aoc2023/utils"
)

type beam struct {
	point     utils.Point
	direction utils.Direction
}

func move(point utils.Point, direction utils.Direction) utils.Point {
	return utils.Point{X: point.X + direction.X, Y: point.Y + direction.Y}
}

func calculateEnergized(grid *utils.Grid[string], startPoint utils.Point, startDirection utils.Direction) int {
	width, height := grid.GetDimensions()
	start := beam{point: startPoint, direction: startDirection}

	seen := map[utils.Point][]utils.Direction{}
	queue := utils.Queue[beam]{}
	queue.Enqueue(start)

	for queue.Len > 0 {
		curr, _ := queue.Dequeue()

		if slices.Contains(seen[curr.point], curr.direction) {
			continue
		}
		seen[curr.point] = append(seen[curr.point], curr.direction)

		next := []beam{}
		symbol := grid.GetValue(curr.point)
		switch symbol {
		case ".":
			next = append(next, beam{point: move(curr.point, curr.direction), direction: curr.direction})
		case "-":
			switch curr.direction {
			case utils.DIRECTION_LEFT, utils.DIRECTION_RIGHT:
				next = append(next, beam{point: move(curr.point, curr.direction), direction: curr.direction})
			case utils.DIRECTION_DOWN, utils.DIRECTION_UP:
				left := beam{point: move(curr.point, utils.DIRECTION_LEFT), direction: utils.DIRECTION_LEFT}
				right := beam{point: move(curr.point, utils.DIRECTION_RIGHT), direction: utils.DIRECTION_RIGHT}
				next = append(next, left, right)
			}
		case "|":
			switch curr.direction {
			case utils.DIRECTION_UP, utils.DIRECTION_DOWN:
				next = append(next, beam{point: move(curr.point, curr.direction), direction: curr.direction})
			case utils.DIRECTION_LEFT, utils.DIRECTION_RIGHT:
				up := beam{point: move(curr.point, utils.DIRECTION_UP), direction: utils.DIRECTION_UP}
				down := beam{point: move(curr.point, utils.DIRECTION_DOWN), direction: utils.DIRECTION_DOWN}
				next = append(next, up, down)
			}
		case "/":
			newDirection := utils.Direction{Y: curr.direction.X * -1, X: curr.direction.Y * -1}
			next = append(next, beam{point: move(curr.point, newDirection), direction: newDirection})
		case "\\":
			newDirection := utils.Direction{Y: curr.direction.X, X: curr.direction.Y}
			next = append(next, beam{point: move(curr.point, newDirection), direction: newDirection})
		}

		for _, beam := range next {
			if beam.point.X < 0 || beam.point.Y < 0 || beam.point.X >= width || beam.point.Y >= height {
				continue
			}
			queue.Enqueue(beam)
		}
	}
	return len(seen)
}

func part1(path string) int {
	content := utils.GetContent(path)
	grid := utils.NewStrGrid(content)

	startPoint := utils.Point{X: 0, Y: 0}
	startDirection := utils.DIRECTION_RIGHT

	return calculateEnergized(grid, startPoint, startDirection)
}

func part2(path string) (total int) {
	content := utils.GetContent(path)
	grid := utils.NewStrGrid(content)
	width, height := grid.GetDimensions()

	for x := 0; x < width; x++ {
		topRow := calculateEnergized(grid, utils.Point{X: x, Y: 0}, utils.DIRECTION_DOWN)
		if topRow > total {
			total = topRow
		}
		bottomRow := calculateEnergized(grid, utils.Point{X: x, Y: height - 1}, utils.DIRECTION_UP)
		if bottomRow > total {
			total = bottomRow
		}
	}
	for y := 0; y < height; y++ {
		leftColumn := calculateEnergized(grid, utils.Point{X: 0, Y: y}, utils.DIRECTION_RIGHT)
		if leftColumn > total {
			total = leftColumn
		}
		rightColumn := calculateEnergized(grid, utils.Point{X: width - 1, Y: y}, utils.DIRECTION_LEFT)
		if rightColumn > total {
			total = rightColumn
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
