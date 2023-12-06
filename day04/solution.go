package day04

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\r\n")
	sum := 0.0
	for _, line := range lines {
		w := 0.0
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers, cardNumbers := strings.Split(strings.TrimSpace(numbers[0]), " "), strings.Split(strings.TrimSpace(numbers[1]), " ")
		winningNumberSet := make(map[int]bool, 0)
		for i := range winningNumbers {
			val := strings.TrimSpace(winningNumbers[i])
			if val != "" {
				key, _ := strconv.Atoi(val)
				winningNumberSet[key] = true
			}
		}
		for i := range cardNumbers {
			val := strings.TrimSpace(cardNumbers[i])
			if val != "" {
				key, _ := strconv.Atoi(val)
				if _, ok := winningNumberSet[key]; ok {
					w++
				}
			}
		}
		if w != 0 {
			sum += math.Exp2(w - 1.0)
		}
	}

	return int(sum)
}

func Part2(input string) int {
	lines := strings.Split(input, "\r\n")
	tickets := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		tickets[i] = 1
	}

	for i := range lines {
		w := 0
		numbers := strings.Split(strings.Split(lines[i], ":")[1], "|")
		winningNumbers, cardNumbers := strings.Split(strings.TrimSpace(numbers[0]), " "), strings.Split(strings.TrimSpace(numbers[1]), " ")
		winningNumberSet := make(map[int]bool, 0)
		for j := range winningNumbers {
			val := strings.TrimSpace(winningNumbers[j])
			if val != "" {
				key, _ := strconv.Atoi(val)
				winningNumberSet[key] = true
			}
		}
		for j := range cardNumbers {
			val := strings.TrimSpace(cardNumbers[j])
			if val != "" {
				key, _ := strconv.Atoi(val)
				if _, ok := winningNumberSet[key]; ok {
					w++
				}
			}
		}
		for k := i + 1; k <= i+w; k++ {
			tickets[k] += tickets[i]
		}
	}

	sum := 0
	for _, v := range tickets {
		sum += v
	}
	return sum
}
