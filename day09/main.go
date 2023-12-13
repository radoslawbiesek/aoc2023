package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func getInput(path string) (sequences [][]int) {
	lines := utils.GetLines(path, "\n")
	for _, line := range lines {
		sequence := utils.Map(strings.Fields(line), utils.ParseInt)
		sequences = append(sequences, sequence)
	}

	return
}

func hasAllZeros(s []int) bool {
	return !slices.ContainsFunc(s, func(el int) bool {
		return el != 0
	})
}

func last[T any](s []T) T {
	return s[len(s)-1]
}

func solution(path string, reverse bool) (total int) {
	for _, sequence := range getInput(path) {
		if reverse {
			slices.Reverse(sequence)
		}
		history := [][]int{sequence}

		for !hasAllZeros(last(history)) {
			nextSequence := []int{}
			for i := 0; i < len(last(history))-1; i++ {
				nextSequence = append(nextSequence, last(history)[i+1]-last(history)[i])
			}
			history = append(history, nextSequence)
		}

		history[len(history)-1] = append(last(history), 0)
		for i := len(history) - 2; i >= 0; i-- {
			history[i] = append(history[i], last(history[i])+last(history[i+1]))
		}

		total += last(history[0])
	}

	return
}

func part1(path string) int {
	return solution(path, false)
}

func part2(path string) int {
	return solution(path, true)
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
