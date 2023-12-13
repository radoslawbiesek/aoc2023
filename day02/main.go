package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

var MAX_BAG = Cubes{
	red:   12,
	green: 13,
	blue:  14,
}

func parseLine(line string) (cubes Cubes) {
	game := strings.TrimSpace(strings.Split(line, ":")[1])
	roundStrs := strings.Split(game, ";")
	for _, roundStr := range roundStrs {
		roundStr = strings.TrimSpace(roundStr)
		cubeStrs := strings.Split(roundStr, ",")
		for _, cubeStr := range cubeStrs {
			cubeStr = strings.TrimSpace(cubeStr)
			splitted := strings.Fields(cubeStr)
			num := utils.ParseInt(splitted[0])
			color := splitted[1]

			if color == "blue" && num >= cubes.blue {
				cubes.blue = num
			} else if color == "red" && num >= cubes.red {
				cubes.red = num
			} else if color == "green" && num >= cubes.green {
				cubes.green = num
			}
		}
	}

	return
}

func part1(path string) (sum int) {
	lines := utils.GetLines(path, "\n")
	for index, line := range lines {
		gameIndex := index + 1
		game := parseLine(line)
		if game.blue <= MAX_BAG.blue && game.red <= MAX_BAG.red && game.green <= MAX_BAG.green {
			sum += gameIndex
		}
	}

	return
}

func part2(path string) (sum int) {
	lines := utils.GetLines(path, "\n")
	for _, line := range lines {
		game := parseLine(line)
		sum += game.blue * game.green * game.red
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
