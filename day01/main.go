package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func getInput(path string) []string {
	lines := utils.GetLines(path, "\n")

	return lines
}

func part1(path string) (sum int) {
	lines := getInput(path)
	for _, line := range lines {
		digits := []int{}
		for _, char := range strings.Split(line, "") {
			digit, err := strconv.Atoi(char)
			if err == nil {
				digits = append(digits, digit)
			}
		}
		value := 10*digits[0] + digits[len(digits)-1]
		sum += value
	}

	return
}

var DIGITS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part2(path string) (sum int) {
	lines := utils.GetLines(path, "\n")
	for _, line := range lines {
		digits := []int{}
		acc := ""
		for _, char := range strings.Split(line, "") {
			digit, err := strconv.Atoi(char)

			if err == nil {
				digits = append(digits, digit)
				acc = ""
			} else {
				acc += char
				for str, num := range DIGITS {
					idx := strings.Index(acc, str)
					if idx != -1 {
						digits = append(digits, num)
						acc = acc[idx+1:]
						break
					}
				}
			}
		}

		value := 10*digits[0] + digits[len(digits)-1]
		sum += value
	}

	return
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./test-input2.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./input.txt"))
}
