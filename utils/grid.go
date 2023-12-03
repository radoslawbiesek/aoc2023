package utils

import "strings"

type Point struct {
	X, Y int
}

type direction struct {
	x, y int
}

var directions4 = []direction{
	{y: -1, x: 0}, // up
	{y: 0, x: 1},  // right
	{y: 1, x: 0},  // down
	{y: 0, x: -1}, // left
}

var directions8 = []direction{
	{y: -1, x: -1}, // NW
	{y: -1, x: 0},  // N
	{y: -1, x: 1},  // NE
	{y: 0, x: 1},   // E
	{y: 0, x: -1},  // W
	{y: 1, x: 0},   // S
	{y: 1, x: -1},  // SW
	{y: 1, x: 1},   // SE
}

type Grid[T any] [][]T

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

func (g Grid[T]) getNeighbors(curr Point, directions []direction) (points []Point) {
	width, height := g.GetDimensions()
	for _, dir := range directions {
		next := Point{X: curr.X + dir.x, Y: curr.Y + dir.y}
		if next.X >= 0 && next.X < width && next.Y >= 0 && next.Y < height {
			points = append(points, next)
		}
	}

	return
}

func (g Grid[T]) Get4Neighbors(curr Point) []Point {
	return g.getNeighbors(curr, directions4)
}

func (g Grid[T]) Get8Neighbors(curr Point) []Point {
	return g.getNeighbors(curr, directions8)
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
