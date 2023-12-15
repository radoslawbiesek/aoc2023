package main

import (
	"fmt"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func hash(str string) (result int) {
	for i := 0; i < len(str); i++ {
		code := int(str[i])
		result += code
		result *= 17
		result = result % 256
	}
	return
}

func part1(path string) (total int) {
	for _, line := range utils.GetLines(path, ",") {
		total += hash(line)
	}

	return
}

func part2(path string) (total int) {
	return
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	// fmt.Printf("Part 2: %d\n", part2("./test-input.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
	// fmt.Printf("Part 2: %d\n", part2("./input.txt"))
}
