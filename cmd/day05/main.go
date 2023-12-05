package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/rawbits2010/AoC23/internal/inputhandler"
)

type MapElement struct {
	Source      int
	Destination int
	Range       int
}

var categories = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func main() {

	lines := inputhandler.ReadInput()

	seedIDs, categoryMaps, err := parseInput(lines)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	lowestLocationPart1 := math.MaxInt
	for _, source := range seedIDs {

		var dest int
		for _, currCategory := range categories {
			dest = ConvertNumber(categoryMaps[currCategory], source)
			source = dest
		}

		if dest < lowestLocationPart1 {
			lowestLocationPart1 = dest
		}
	}

	lowestLocationPart2 := math.MaxInt
	for i := 0; i < len(seedIDs)-1; i += 2 {

		for source := seedIDs[i]; source < seedIDs[i]+seedIDs[i+1]; source++ {

			dest := source
			for _, currCategory := range categories {
				dest = ConvertNumber(categoryMaps[currCategory], dest)
			}

			if dest < lowestLocationPart2 {
				lowestLocationPart2 = dest
			}
		}
	}

	fmt.Printf("Result - Part 1: %d, Part 2: %d", lowestLocationPart1, lowestLocationPart2)
}

func parseInput(lines []string) ([]int, map[string][]MapElement, error) {

	seedIDs := make([]int, 0)
	categoryMaps := make(map[string][]MapElement)

	var currCategory string
	var mapPart []MapElement
lineProcessor:
	for lineIdx, line := range lines {

		// seeds
		if lineIdx == 0 {

			seedsStr := strings.Split(strings.TrimSpace(line[6:]), " ")
			for _, seed := range seedsStr {

				seedNum, err := strconv.Atoi(strings.TrimSpace(seed))
				if err != nil {
					return nil, nil, fmt.Errorf("couldn't convert seed id '%s'", seed)
				}

				seedIDs = append(seedIDs, seedNum)
			}

			continue
		}

		if len(line) == 0 {
			if len(currCategory) != 0 {
				categoryMaps[currCategory] = mapPart
			}
			currCategory = ""
			continue
		}

		if len(currCategory) == 0 {
			for _, category := range categories {
				if strings.HasPrefix(line, category) {
					mapPart = make([]MapElement, 0)
					currCategory = category
					continue lineProcessor
				}
			}
			return nil, nil, fmt.Errorf("expected a category in line: %s", line)
		}

		mapRanges := strings.Split(line, " ")
		if len(mapRanges) != 3 {
			return nil, nil, fmt.Errorf("invalid mapping range in line: %s", line)
		}

		dest, err := strconv.Atoi(mapRanges[0])
		if err != nil {
			return nil, nil, fmt.Errorf("couldn't convert destination category '%s' in line: %s", mapRanges[0], line)
		}
		source, err := strconv.Atoi(mapRanges[1])
		if err != nil {
			return nil, nil, fmt.Errorf("couldn't convert source category '%s' in line: %s", mapRanges[1], line)
		}
		ranges, err := strconv.Atoi(mapRanges[2])
		if err != nil {
			return nil, nil, fmt.Errorf("couldn't convert range '%s' in line: %s", mapRanges[2], line)
		}

		mapPart = append(mapPart, MapElement{Source: source, Destination: dest, Range: ranges})
	}
	categoryMaps[currCategory] = mapPart

	return seedIDs, categoryMaps, nil
}

func ConvertNumber(categoryMap []MapElement, source int) int {

	for _, elem := range categoryMap {
		if source >= elem.Source && source < elem.Source+elem.Range {
			return elem.Destination + (source - elem.Source)
		}
	}

	return source
}
