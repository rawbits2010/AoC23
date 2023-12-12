package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/rawbits2010/AoC23/internal/inputhandler"
)

func main() {

	lines := inputhandler.ReadInput()

	timeVals, distVals, err := parseRaceTimes(lines)
	if err != nil {
		fmt.Printf("couldn't parse input: %s\n", err)
		os.Exit(-1)
	}

	heldTimeRange := make([]int, 0)
	for raceIdx := 0; raceIdx < len(timeVals); raceIdx++ {
		currTime := timeVals[raceIdx]
		currDist := distVals[raceIdx]

		if currTime == 0 {
			fmt.Print("found a 0 time")
			os.Exit(0)
		}

		timeRange, err := calcHeldTimeRange(currTime, currDist)
		if err != nil {
			fmt.Printf("error while calcing part 1: %s\n", err)
			os.Exit(-1)
		}

		heldTimeRange = append(heldTimeRange, timeRange)
	}

	part1 := heldTimeRange[0]
	for heldTimeIdx, heldTime := range heldTimeRange {
		if heldTimeIdx == 0 {
			continue
		}
		part1 *= heldTime
	}

	timeVal, distVal, err := parseRaceTimesAsOne(lines)
	if err != nil {
		fmt.Printf("couldn't parse input: %s\n", err)
		os.Exit(-1)
	}

	part2, err := calcHeldTimeRange(timeVal, distVal)
	if err != nil {
		fmt.Printf("error while calcing part 2: %s\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Result - Part1: %d, Part2: %d\n", part1, part2)
}

// from the race time and the target distance, returns the number of possible win cases
func calcHeldTimeRange(currTime, currDist int) (int, error) {

	for heldTime := 0; heldTime < int(math.Floor(float64(currTime)/2)); heldTime++ {

		distTravel := (currTime - heldTime) * heldTime
		if distTravel > currDist {

			timeRange := (currTime - ((heldTime - 1) * 2)) - 1
			return timeRange, nil
		}
	}

	return -1, fmt.Errorf("couldn't found a longer distance for race with time: %d, and dist: %d", currTime, currDist)

}

// parses input lines and returns time and distance values in corresponding order
func parseRaceTimes(lines []string) ([]int, []int, error) {

	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("invalid number of input lines")
	}

	timeVals, err := parseRaceTimesLine(lines[0])
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't convert times: %w", err)
	}

	distVals, err := parseRaceTimesLine(lines[1])
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't convert distances: %w", err)
	}

	if len(timeVals) != len(distVals) {
		return nil, nil, fmt.Errorf("mismatching number of entries time: '%d', dist: '%d'", len(timeVals), len(distVals))
	}

	return timeVals, distVals, nil
}

func parseRaceTimesLine(line string) ([]int, error) {

	lineSplit := strings.Split(line, ":")
	if len(lineSplit) != 2 {
		return nil, fmt.Errorf("couldn't find time values")
	}
	lineValStrs := strings.Split(lineSplit[1], " ")

	lineVals := make([]int, 0)
	for _, valStr := range lineValStrs {
		if len(valStr) == 0 {
			continue
		}

		val, err := strconv.Atoi(valStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing value: '%s'", valStr)
		}

		lineVals = append(lineVals, val)
	}

	return lineVals, nil
}

// parses input lines and returns time and distance value
func parseRaceTimesAsOne(lines []string) (int, int, error) {

	if len(lines) != 2 {
		return -1, -1, fmt.Errorf("invalid number of input lines")
	}

	timeVal, err := parseRaceTimesLineAsOne(lines[0])
	if err != nil {
		return -1, -1, fmt.Errorf("couldn't convert time: %w", err)
	}

	distVal, err := parseRaceTimesLineAsOne(lines[1])
	if err != nil {
		return -1, -1, fmt.Errorf("couldn't convert distance: %w", err)
	}

	return timeVal, distVal, nil
}

func parseRaceTimesLineAsOne(line string) (int, error) {

	lineSplit := strings.Split(line, ":")
	if len(lineSplit) != 2 {
		return -1, fmt.Errorf("couldn't find time values")
	}
	lineValStrs := strings.Split(lineSplit[1], " ")

	var valStr strings.Builder
	for _, currValStr := range lineValStrs {
		if len(currValStr) == 0 {
			continue
		}

		_, err := valStr.WriteString(currValStr)
		if err != nil {
			return -1, fmt.Errorf("error concat string")
		}
	}

	value, err := strconv.Atoi(valStr.String())
	if err != nil {
		return -1, fmt.Errorf("error converting value '%s'", valStr.String())
	}

	return value, nil
}
