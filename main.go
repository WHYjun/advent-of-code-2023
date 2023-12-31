package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/WHYjun/advent-of-code-2023/day01"
	"github.com/WHYjun/advent-of-code-2023/day02"
	"github.com/WHYjun/advent-of-code-2023/day04"
	"github.com/WHYjun/advent-of-code-2023/day05"
	"github.com/WHYjun/advent-of-code-2023/day06"
	"github.com/WHYjun/advent-of-code-2023/day07"
	"github.com/WHYjun/advent-of-code-2023/utils"
)

func main() {
	day := mustParseDay()
	fmt.Printf("Running day %02d\n", day)

	switch day {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.MustReadFile(day)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.MustReadFile(day)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.MustReadFile(day)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.MustReadFile(day)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.MustReadFile(day)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.MustReadFile(day)))
	case 5:
		fmt.Printf("part 1: %v\n", day05.Part1(utils.MustReadFile(day)))
		fmt.Printf("part 2: %v\n", day05.Part2(utils.MustReadFile(day)))
	case 6:
		fmt.Printf("part 1: %v\n", day06.Part1(utils.MustReadFile(day)))
		fmt.Printf("part 2: %v\n", day06.Part2(utils.MustReadFile(day)))
	case 7:
		fmt.Printf("part 1: %v\n", day07.Part1(utils.MustReadFile(day)))
		fmt.Printf("part 2: %v\n", day07.Part2(utils.MustReadFile(day)))
	default:
		utils.PanicError(fmt.Errorf("no such day: %d", day))
	}
}

func mustParseDay() int {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		utils.PanicError(fmt.Errorf("should have a vaild integer day (1-25) - %s", err.Error()))
	}
	return day
}
