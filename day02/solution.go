package day02

import (
	"strconv"
	"strings"
)

type CubeCounts struct {
	Red   int
	Green int
	Blue  int
}

func Part1(input string) int {
	sum := 0
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		if playGameOne(line) {
			sum += getGameID(line)
		}
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\r\n")
	sum := 0
	for _, line := range lines {
		sum += playGameTwo(line)
	}
	return sum
}

func getGameID(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	game := strings.Split(s, ":")[0]
	id, _ := strconv.Atoi(strings.ReplaceAll(game, "Game", ""))
	return id
}

func playGameOne(s string) bool {
	sets := strings.Split(strings.Split(s, ":")[1], ";")
	for _, set := range sets {
		count := CubeCounts{}
		blocks := strings.Split(set, ",")
		for _, block := range blocks {
			block = strings.ReplaceAll(block, " ", "")
			if strings.HasSuffix(block, "red") {
				count.Red, _ = strconv.Atoi(strings.ReplaceAll(block, "red", ""))
			} else if strings.HasSuffix(block, "green") {
				count.Green, _ = strconv.Atoi(strings.ReplaceAll(block, "green", ""))
			} else if strings.HasSuffix(block, "blue") {
				count.Blue, _ = strconv.Atoi(strings.ReplaceAll(block, "blue", ""))
			}
		}
		if !isValidGame(count) {
			return false
		}
	}
	return true
}

func isValidGame(count CubeCounts) bool {
	maxCount := CubeCounts{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	return count.Red <= maxCount.Red && count.Green <= maxCount.Green && count.Blue <= maxCount.Blue
}

func playGameTwo(s string) int {
	minCubeCounts := CubeCounts{
		Red:   -1,
		Green: -1,
		Blue:  -1,
	}
	sets := strings.Split(strings.Split(s, ":")[1], ";")
	for _, set := range sets {
		count := CubeCounts{}
		blocks := strings.Split(set, ",")
		for _, block := range blocks {
			block = strings.ReplaceAll(block, " ", "")
			if strings.HasSuffix(block, "red") {
				count.Red, _ = strconv.Atoi(strings.ReplaceAll(block, "red", ""))
			} else if strings.HasSuffix(block, "green") {
				count.Green, _ = strconv.Atoi(strings.ReplaceAll(block, "green", ""))
			} else if strings.HasSuffix(block, "blue") {
				count.Blue, _ = strconv.Atoi(strings.ReplaceAll(block, "blue", ""))
			}
		}
		if count.Red > minCubeCounts.Red {
			minCubeCounts.Red = count.Red
		}
		if count.Blue > minCubeCounts.Blue {
			minCubeCounts.Blue = count.Blue
		}
		if count.Green > minCubeCounts.Green {
			minCubeCounts.Green = count.Green
		}
	}
	return minCubeCounts.Blue * minCubeCounts.Green * minCubeCounts.Red
}
