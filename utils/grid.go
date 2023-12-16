package utils

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y int
}

type Direction struct {
	X, Y int
}

var DIRECTION_UP = Direction{Y: -1, X: 0}
var DIRECTION_RIGHT = Direction{Y: 0, X: 1}
var DIRECTION_DOWN = Direction{Y: 1, X: 0}
var DIRECTION_LEFT = Direction{Y: 0, X: -1}

var Directions4 = []Direction{DIRECTION_DOWN, DIRECTION_LEFT, DIRECTION_UP, DIRECTION_RIGHT}

var DIRECTION_NW = Direction{Y: -1, X: -1}
var DIRECTION_NE = Direction{Y: -1, X: 1}
var DIRECTION_SW = Direction{Y: 1, X: -1}
var DIRECTION_SE = Direction{Y: 1, X: 1}

var Directions8 = []Direction{DIRECTION_DOWN, DIRECTION_LEFT, DIRECTION_UP, DIRECTION_RIGHT, DIRECTION_NE, DIRECTION_NW, DIRECTION_SE, DIRECTION_SW}

type Grid[T any] [][]T

func (g *Grid[T]) String() (str string) {
	for _, row := range *g {
		for _, col := range row {
			str += fmt.Sprintf("%v", col)
		}
		str += "\n"
	}
	return
}

func (g Grid[T]) GetDimensions() (width, height int) {
	width = len(g[0])
	height = len(g)

	return
}

func (g Grid[T]) GetAllPoints() (points []Point) {
	width, height := g.GetDimensions()
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			points = append(points, Point{X: col, Y: row})
		}
	}

	return
}

func (g Grid[T]) GetValue(p Point) T {
	return g[p.Y][p.X]
}

func (g *Grid[T]) SetValue(p Point, value T) {
	(*g)[p.Y][p.X] = value
}

func (g Grid[T]) getNeighbors(curr Point, directions []Direction) (points []Point) {
	width, height := g.GetDimensions()
	for _, dir := range directions {
		next := Point{X: curr.X + dir.X, Y: curr.Y + dir.Y}
		if next.X >= 0 && next.X < width && next.Y >= 0 && next.Y < height {
			points = append(points, next)
		}
	}

	return
}

func (g Grid[T]) Get4Neighbors(curr Point) []Point {
	return g.getNeighbors(curr, Directions4)
}

func (g Grid[T]) Get8Neighbors(curr Point) []Point {
	return g.getNeighbors(curr, Directions8)
}

func NewStrGrid(str string) *Grid[string] {
	grid := Grid[string]{}
	lines := strings.Split(str, "\n")
	for _, lineStr := range lines {
		line := []string{}
		for _, char := range strings.Split(lineStr, "") {
			line = append(line, char)
		}
		grid = append(grid, line)
	}

	return &grid
}
