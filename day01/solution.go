package day01

import (
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\r\n")

	sum := 0
	for _, line := range lines {
		firstDigit, lastDigit := firstNumber(line), lastNumber(line)
		if lastDigit == 0 {
			lastDigit = firstDigit
		}
		sum += firstDigit*10 + lastDigit
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\r\n")
	sum := 0
	for _, line := range lines {
		sum += firstNumberOrDigit(line)*10 + lastNumberOrDigit(line)
	}
	return sum
}

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func firstNumber(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
	}
	return 0
}

func firstNumberOrDigit(s string) int {
	acc := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc += string(s[i])

		for i, d := range digits {
			if strings.HasSuffix(acc, d) {
				return i + 1
			}
		}
	}
	return 0
}

func lastNumber(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
	}
	return 0
}

func lastNumberOrDigit(s string) int {
	acc := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc = string(s[i]) + acc

		for i, d := range digits {
			if strings.HasPrefix(acc, d) {
				return i + 1
			}
		}
	}
	return 0
}
