package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func getInput(path string) (times, distances []int) {
	lines := utils.GetLines(path, "\n")
	timeStr := lines[0]
	for _, numStr := range strings.Fields(timeStr)[1:] {
		if numStr == "" {
			continue
		}
		times = append(times, utils.ParseInt(numStr))
	}

	dostanceStr := lines[1]
	for _, numStr := range strings.Fields(dostanceStr)[1:] {
		if numStr == "" {
			continue
		}
		distances = append(distances, utils.ParseInt(numStr))
	}

	return
}

func getInput2(path string) (time, distance int) {
	lines := utils.GetLines(path, "\n")
	timeStr := lines[0]
	timeStr = strings.Split(timeStr, ":")[1]
	timeStr = strings.ReplaceAll(timeStr, " ", "")
	time = utils.ParseInt(timeStr)

	distanceStr := lines[1]
	distanceStr = strings.Split(distanceStr, ":")[1]
	distanceStr = strings.ReplaceAll(distanceStr, " ", "")
	distance = utils.ParseInt(distanceStr)

	return
}

func part1(path string) (total int) {
	times, distances := getInput(path)
	total = 1

	for raceIdx := 0; raceIdx < len(times); raceIdx++ {
		raceScore := 0
		time := times[raceIdx]
		distance := distances[raceIdx]

		for hold := 0; hold <= time; hold++ {
			timeLeft := time - hold
			speed := hold

			if timeLeft*speed > distance {
				raceScore++
			}
		}

		total *= raceScore
	}

	return
}

func part2(path string) (total int) {
	time, distance := getInput2(path)

	for hold := 0; hold <= time; hold++ {
		timeLeft := time - hold
		speed := hold

		if timeLeft*speed > distance {
			total++
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
