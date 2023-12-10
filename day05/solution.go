package day05

import (
	"math"
	"strconv"
	"strings"
)

type Map struct {
	Seeds                 []int
	SeedToSoil            map[int]int
	SoilToFertilizer      map[int]int
	FertilizerToWater     map[int]int
	WaterToLight          map[int]int
	LightToTemperature    map[int]int
	TemperatureToHumidity map[int]int
	HumidityToLocation    map[int]int
}

func Part1(input string) int {
	lines := strings.Split(input, "\r\n")
	m := parseMap(lines)
	var soil, fertilizer, water, light, temperature, humidity, location int

	minLocation := math.MaxInt
	for _, seed := range m.Seeds {
		soil = match(seed, m.SeedToSoil)
		fertilizer = match(soil, m.SoilToFertilizer)
		water = match(fertilizer, m.FertilizerToWater)
		light = match(water, m.WaterToLight)
		temperature = match(light, m.LightToTemperature)
		humidity = match(temperature, m.TemperatureToHumidity)
		location = match(humidity, m.HumidityToLocation)
		if minLocation > location {
			minLocation = location
		}
	}

	return minLocation
}

func Part2(input string) int {
	lines := strings.Split(input, "\r\n")
	return len(lines)
}

func parseMap(lines []string) *Map {
	m := &Map{}
	var seeds []int
	var stsIdx, stfIdx, ftwIdx, wtlIdx, lttIdx, tthIdx, htlIdx int

	for i := range lines {
		if strings.HasPrefix(lines[i], "seeds") {
			splitSeeds := strings.Split(strings.ReplaceAll(lines[0], "seeds: ", ""), " ")
			for j := range splitSeeds {
				s, _ := strconv.Atoi(strings.TrimSpace(splitSeeds[j]))
				seeds = append(seeds, s)
			}
			m.Seeds = seeds
		} else {
			switch lines[i] {
			case "seed-to-soil map:":
				stsIdx = i
			case "soil-to-fertilizer map:":
				stfIdx = i
			case "fertilizer-to-water map:":
				ftwIdx = i
			case "water-to-light map:":
				wtlIdx = i
			case "light-to-temperature map:":
				lttIdx = i
			case "temperature-to-humidity map:":
				tthIdx = i
			case "humidity-to-location map:":
				htlIdx = i
			default:
				continue
			}
		}
	}

	m.SeedToSoil = iterate(lines, stsIdx+1, stfIdx-1)
	m.SoilToFertilizer = iterate(lines, stfIdx+1, ftwIdx-1)
	m.FertilizerToWater = iterate(lines, ftwIdx+1, wtlIdx-1)
	m.WaterToLight = iterate(lines, wtlIdx+1, lttIdx-1)
	m.LightToTemperature = iterate(lines, lttIdx+1, tthIdx-1)
	m.TemperatureToHumidity = iterate(lines, tthIdx+1, htlIdx-1)
	m.HumidityToLocation = iterate(lines, htlIdx+1, len(lines)-1)

	return m
}

func iterate(lines []string, start, end int) map[int]int {
	m := map[int]int{}

	for i := start; i < end; i++ {
		s := strings.Split(lines[i], " ")
		if len(s) == 3 {
			dest, _ := strconv.Atoi(strings.TrimSpace(s[0]))
			source, _ := strconv.Atoi(strings.TrimSpace(s[1]))
			length, _ := strconv.Atoi(strings.TrimSpace(s[2]))

			for j := 0; j < length; j++ {
				m[source+j] = dest + j
			}
		}
	}
	return m
}

func match(source int, m map[int]int) int {
	if _, ok := m[source]; ok {
		return m[source]
	}
	return source
}
