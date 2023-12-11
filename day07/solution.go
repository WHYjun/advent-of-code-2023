package day07

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []int
	Bid   int
	Score int
}

const (
	fiveofakind  = 7
	fourofakind  = 6
	fullhouse    = 5
	threeofakind = 4
	twopair      = 3
	onepair      = 2
	highcard     = 1
)

func Part1(input string) int {
	lines := strings.Split(input, "\r\n")
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		s := strings.Split(line, " ")
		cards := getCards(s[0])
		bid, _ := strconv.Atoi(s[1])
		h := Hand{
			Cards: cards,
			Bid:   bid,
			Score: getScore(cards),
		}
		hands[i] = h
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Score == b.Score {
			for i := 0; i < 5; i++ {
				if a.Cards[i] == b.Cards[i] {
					continue
				}
				return cmp.Compare(a.Cards[i], b.Cards[i])
			}
		}
		return cmp.Compare(a.Score, b.Score)
	})

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += ((i + 1) * hands[i].Bid)
	}

	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\r\n")
	return len(lines)
}

func getCards(hand string) []int {
	cards := make([]int, 5)
	for i, c := range hand {
		s := string(c)
		if s == "A" {
			cards[i] = 14
		} else if s == "K" {
			cards[i] = 13
		} else if s == "Q" {
			cards[i] = 12
		} else if s == "J" {
			cards[i] = 11
		} else if s == "T" {
			cards[i] = 10
		} else {
			v, _ := strconv.Atoi(s)
			cards[i] = v
		}
	}
	return cards
}

func getScore(cards []int) int {
	cardCount := map[int]int{}
	for _, card := range cards {
		if _, ok := cardCount[card]; ok {
			cardCount[card]++
		} else {
			cardCount[card] = 1
		}
	}

	if len(cardCount) == 5 {
		return highcard
	}
	if len(cardCount) == 1 {
		return fiveofakind
	}
	if len(cardCount) == 2 {
		for _, v := range cardCount {
			if v == 1 || v == 4 {
				return fourofakind
			}
		}
		return fullhouse
	}
	if len(cardCount) == 3 {
		for _, v := range cardCount {
			if v == 2 {
				return twopair
			}
		}
		return threeofakind
	}
	if len(cardCount) == 4 {
		return onepair
	}
	return 0
}
