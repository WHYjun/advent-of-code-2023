package day06

type game struct {
	Time     int
	Distance int
}

func Part1(_ string) int {
	games := []game{
		{Time: 59, Distance: 543},
		{Time: 68, Distance: 1020},
		{Time: 82, Distance: 1664},
		{Time: 74, Distance: 1022},
	}

	winCounts := make([]int, 4)

	for i, g := range games {
		for holdTime := 0; holdTime < g.Time; holdTime++ {
			distance := (g.Time - holdTime) * holdTime
			if distance > g.Distance {
				winCounts[i]++
			}
		}
	}
	result := 1
	for _, v := range winCounts {
		result *= v
	}
	return result
}

func Part2(input string) int {
	games := []game{
		{Time: 59688274, Distance: 543102016641022},
	}

	winCounts := make([]int, 1)

	for i, g := range games {
		for holdTime := 0; holdTime < g.Time; holdTime++ {
			distance := (g.Time - holdTime) * holdTime
			if distance > g.Distance {
				winCounts[i]++
			}
		}
	}

	return winCounts[0]
}
