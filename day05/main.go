package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

type sourceDestinationEntry struct {
	source      int
	destination int
	rangeLen    int
}

func parseInput(path string) (seeds []int, maps [][]sourceDestinationEntry) {
	lines := utils.GetLines(path, "\n")

	seedsStr := lines[0]
	seedsStr = strings.TrimSpace(strings.Split(seedsStr, ":")[1])
	for _, seedStr := range strings.Split(seedsStr, " ") {
		seeds = append(seeds, utils.ParseInt(seedStr))
	}

	for _, lineStr := range lines[1:] {
		if lineStr == "" {
			continue
		}

		if !utils.IsInt(string(lineStr[0])) {
			maps = append(maps, []sourceDestinationEntry{})
			continue
		}

		numStrs := strings.Split(lineStr, " ")
		entry := sourceDestinationEntry{
			destination: utils.ParseInt(numStrs[0]),
			source:      utils.ParseInt(numStrs[1]),
			rangeLen:    utils.ParseInt(numStrs[2]),
		}
		maps[len(maps)-1] = append(maps[len(maps)-1], entry)
	}

	return
}

func part1(path string) (lowest int) {
	seeds, maps := parseInput(path)

	for _, seed := range seeds {
		curr := seed
		for _, destinationMap := range maps {
			for _, entry := range destinationMap {
				if curr >= entry.source && curr < entry.source+entry.rangeLen {
					curr = curr + entry.destination - entry.source
					break
				}
			}
		}

		if lowest == 0 || curr < lowest {
			lowest = curr
		}
	}

	return
}

type seedEntry struct {
	start    int
	rangeLen int
}

func part2(path string) (lowest int) {
	seeds, maps := parseInput(path)

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		rangeLen := seeds[i+1]
		for j := 0; j < rangeLen; j++ {
			seed := start + j
			curr := seed
			for _, destinationMap := range maps {
				for _, entry := range destinationMap {
					if curr >= entry.source && curr < entry.source+entry.rangeLen {
						curr = curr + entry.destination - entry.source
						break
					}
				}
			}

			if lowest == 0 || curr < lowest {
				lowest = curr
			}
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
