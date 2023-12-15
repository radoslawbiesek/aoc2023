package main

import (
	"fmt"
	"slices"
	"strings"

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

type lens struct {
	label  string
	length int
}

func part2(path string) (total int) {
	boxes := [256][]lens{}
	for _, line := range utils.GetLines(path, ",") {
		var operation, label string
		var length int

		if strings.Contains(line, "-") {
			operation = "-"
			label = strings.Replace(line, operation, "", 1)
		} else {
			operation = "="
			splitted := strings.Split(line, operation)
			label = splitted[0]
			length = utils.ParseInt(splitted[1])
		}

		boxIndex := hash(label)
		box := boxes[boxIndex]
		index := slices.IndexFunc(box, func(lens lens) bool {
			return lens.label == label
		})

		switch operation {
		case "-":
			if index != -1 {
				box = append(box[:index], box[index+1:]...)
			}
		case "=":
			lens := lens{label: label, length: length}
			if index == -1 {
				box = append(box, lens)
			} else {
				box[index] = lens
			}
		}

		boxes[boxIndex] = box
	}

	for boxIndex, box := range boxes {
		for lensIndex, lens := range box {
			total += (boxIndex + 1) * (lensIndex + 1) * lens.length
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
