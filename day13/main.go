package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func columnStr(pattern string, index int) (str string) {
	for _, row := range strings.Split(pattern, "\n") {
		str += string(row[index])
	}
	return
}

func compare(str1, str2 string) bool {
	count := 0
	for i1, b1 := range str1 {
		if b1 == rune(str2[i1]) {
			continue
		}
		count++
		if count > 1 {
			return false
		}
	}
	return true
}

func countHorizontal(pattern string) int {
	rows := strings.Split(pattern, "\n")
	for yStart := 0; yStart < len(rows); yStart++ {
		isReflection := false
		for d := 0; yStart-d >= 0 && yStart+d < len(rows)-1; d++ {
			if rows[yStart-d] == rows[yStart+d+1] {
				isReflection = true
			} else {
				isReflection = false
				break
			}
		}
		if isReflection {
			return yStart
		}
	}
	return -1
}

func countHorizontalWithSmudge(pattern string) int {
	rows := strings.Split(pattern, "\n")
	for yStart := 0; yStart < len(rows); yStart++ {
		isReflection := false
		smudge := false
		for d := 0; yStart-d >= 0 && yStart+d < len(rows)-1; d++ {
			row1 := rows[yStart-d]
			row2 := rows[yStart+d+1]

			if row1 == row2 {
				isReflection = true
				continue
			}

			if compare(row1, row2) && !smudge {
				isReflection = true
				smudge = true
			} else {
				isReflection = false
				break
			}
		}
		if isReflection && smudge {
			return yStart
		}
	}
	return -1
}

func countVertical(pattern string) int {
	rows := strings.Split(pattern, "\n")
	for xStart := 0; xStart < len(rows[0]); xStart++ {
		isReflection := false
		for d := 0; xStart-d >= 0 && xStart+d < len(rows[0])-1; d++ {
			if columnStr(pattern, xStart-d) == columnStr(pattern, xStart+d+1) {
				isReflection = true
			} else {
				isReflection = false
				break
			}
		}
		if isReflection {
			return xStart
		}
	}
	return -1
}

func countVerticalWithSmudge(pattern string) int {
	rows := strings.Split(pattern, "\n")
	for xStart := 0; xStart < len(rows[0]); xStart++ {
		isReflection := false
		smudge := false
		for d := 0; xStart-d >= 0 && xStart+d < len(rows[0])-1; d++ {
			col1 := columnStr(pattern, xStart-d)
			col2 := columnStr(pattern, xStart+d+1)

			if col1 == col2 {
				isReflection = true
				continue
			}

			if compare(col1, col2) && !smudge {
				smudge = true
				isReflection = true
			} else {
				isReflection = false
				break
			}
		}
		if isReflection && smudge {
			return xStart
		}
	}
	return -1
}

func part1(path string) (total int) {
	for _, pattern := range utils.GetLines(path, "\n\n") {
		horizontal := countHorizontal(pattern)
		total += (horizontal + 1) * 100

		if horizontal == -1 {
			vertical := countVertical(pattern)
			total += (vertical + 1)
		}
	}
	return
}

func part2(path string) (total int) {
	for _, pattern := range utils.GetLines(path, "\n\n") {
		horizontal := countHorizontalWithSmudge(pattern)
		total += (horizontal + 1) * 100

		if horizontal == -1 {
			vertical := countVerticalWithSmudge(pattern)
			total += (vertical + 1)
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
