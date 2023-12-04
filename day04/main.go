package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

func parseLine(lineStr string) (cardIndex int, winningNumbers map[int]bool, myNumbers []int) {
	winningNumbers = map[int]bool{}

	splitted := strings.Split(lineStr, ":")

	indexSplitted := strings.Split(splitted[0], " ")
	cardIndex = utils.ParseInt(indexSplitted[len(indexSplitted)-1])

	numbersStr := strings.TrimSpace(splitted[1])
	numbersStrSplitted := strings.Split(numbersStr, "|")

	winningNumbersStr := strings.TrimSpace(numbersStrSplitted[0])
	for _, winningNumberStr := range strings.Split(winningNumbersStr, " ") {
		if winningNumberStr == "" {
			continue
		}
		winningNumber := utils.ParseInt(winningNumberStr)
		winningNumbers[winningNumber] = true
	}

	myNumbersStr := strings.TrimSpace(numbersStrSplitted[1])
	for _, myNumberStr := range strings.Split(myNumbersStr, " ") {
		if myNumberStr == "" {
			continue
		}
		myNumber := utils.ParseInt(myNumberStr)
		myNumbers = append(myNumbers, myNumber)
	}

	return
}

func part1(path string) (total int) {
	lines := utils.GetLines(path, "\n")
	for _, lineStr := range lines {
		_, winningNumbers, myNumbers := parseLine(lineStr)
		cardScore := 0

		for _, myNumber := range myNumbers {
			has := winningNumbers[myNumber]
			if !has {
				continue
			}

			if cardScore == 0 {
				cardScore = 1
			} else {
				cardScore *= 2
			}
		}
		total += cardScore
	}

	return
}

func part2(path string) (total int) {
	lines := utils.GetLines(path, "\n")
	queue := utils.Queue[string]{}
	for _, lineStr := range lines {
		queue.Enqueue(lineStr)
	}
	cache := map[string][]string{}

	for queue.Len > 0 {
		lineStr, ok := queue.Dequeue()
		if !ok {
			return
		}

		total++

		cached, ok := cache[*lineStr]

		if ok {
			for _, nextLineStr := range cached {
				queue.Enqueue(nextLineStr)
			}
			continue
		}

		cardIndex, winningNumbers, myNumbers := parseLine(*lineStr)
		next := cardIndex

		for _, myNumber := range myNumbers {
			has := winningNumbers[myNumber]
			if has {
				line := lines[next]
				queue.Enqueue(line)
				cached = append(cached, line)
				next++
			}
		}
		cache[*lineStr] = cached
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
