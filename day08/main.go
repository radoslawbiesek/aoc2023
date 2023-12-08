package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

type instructions = []string
type nodeMap = map[string][]string

func getInput(path string) (*instructions, *nodeMap) {
	lines := utils.GetLines(path, "\n")

	instructions := strings.Split(lines[0], "")

	nodeMap := nodeMap{}
	for _, line := range lines[2:] {
		splitted := strings.Split(line, "=")

		key := strings.TrimSpace(splitted[0])

		dirsStr := strings.ReplaceAll(splitted[1], " ", "")
		dirsStr = strings.ReplaceAll(dirsStr, "(", "")
		dirsStr = strings.ReplaceAll(dirsStr, ")", "")
		nodeMap[key] = strings.Split(dirsStr, ",")
	}

	return &instructions, &nodeMap
}

func part1(path string) (steps int) {
	const START = "AAA"
	const END = "ZZZ"
	instructions, nodesMap := getInput(path)

	curr := START
	for curr != END {
		dir := (*instructions)[steps%len(*instructions)]
		node, _ := (*nodesMap)[curr]

		if dir == "L" {
			curr = node[0]
		} else {
			curr = node[1]
		}
		steps++
	}

	return
}

func part2(path string) int {
	instructions, nodesMap := getInput(path)
	minSteps := []int{}

	for pos := range *nodesMap {
		if !strings.HasSuffix(pos, "A") {
			continue
		}

		curr := pos
		step := 0
		for {
			if strings.HasSuffix(curr, "Z") {
				minSteps = append(minSteps, step)
				break
			}

			dir := (*instructions)[step%len(*instructions)]
			node, _ := (*nodesMap)[curr]
			if dir == "L" {
				curr = node[0]
			} else {
				curr = node[1]
			}
			step++
		}
	}

	return utils.LCMSlice(minSteps)
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
