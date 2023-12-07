package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

const (
	HighCard int = iota
	Pair
	TwoPair
	Three
	Full
	Four
	Five
)

func getCardScore(card string, jokerRule bool) int {
	switch card {
	case "T":
		return 10
	case "J":
		if jokerRule {
			return 1
		}
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}

	return utils.ParseInt(card)
}

func calculateHandScore(cardScores []string, jokerRule bool) int {
	cardCount := map[string]int{}
	jokers := 0

	for _, card := range cardScores {
		if jokerRule && card == "J" {
			jokers++
			continue
		}

		val, ok := cardCount[card]
		if !ok {
			cardCount[card] = 1
		} else {
			cardCount[card] = val + 1
		}
	}

	switch len(cardCount) {
	case 0, 1:
		return Five
	case 2:
		for _, count := range cardCount {
			if count == 4 || (jokerRule && count+jokers == 4) {
				return Four
			}
		}
		return Full
	case 3:
		for _, count := range cardCount {
			if count == 3 || (jokerRule && count+jokers == 3) {
				return Three
			}
		}
		return TwoPair
	case 4:
		return Pair
	}

	return HighCard
}

type hand struct {
	handScore  int
	cardScores []int
	bid        int
}

func parseLine(line string, jokerRule bool) *hand {
	splitted := strings.Split(line, " ")

	cards := strings.Split(splitted[0], "")
	cardScores := utils.Map(cards, func(card string) int {
		return getCardScore(card, jokerRule)
	})

	return &hand{
		cardScores: cardScores,
		bid:        utils.ParseInt(splitted[1]),
		handScore:  calculateHandScore(cards, jokerRule),
	}
}

func compare(a, b *hand) int {
	if a.handScore > b.handScore {
		return 1
	} else if b.handScore > a.handScore {
		return -1
	}

	for i := 0; i < len(a.cardScores); i++ {
		if a.cardScores[i] > b.cardScores[i] {
			return 1
		} else if b.cardScores[i] > a.cardScores[i] {
			return -1
		}
	}

	return 0
}

func calculateTotal(path string, jokerRule bool) (total int) {
	lines := utils.GetLines(path, "\n")
	hands := utils.Map(lines, func(line string) *hand {
		return parseLine(line, jokerRule)
	})
	slices.SortFunc(hands, compare)

	for i, hand := range hands {
		rank := i + 1
		total += hand.bid * rank
	}

	return
}

func part1(path string) (total int) {
	return calculateTotal(path, false)
}

func part2(path string) (total int) {
	return calculateTotal(path, true)
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
