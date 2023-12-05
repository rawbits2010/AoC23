package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rawbits2010/AoC23/internal/inputhandler"
)

const (
	redMax   = 12
	greenMax = 13
	blueMax  = 14
)

func main() {

	lines := inputhandler.ReadInput()

	var part1 int
	var part2 int
	for _, line := range lines {
		gameNum, redCubes, greenCubes, blueCubes, err := parseLine(line)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		// Part 1
		if IsValidCubeCount(redCubes, redMax) {
			if IsValidCubeCount(greenCubes, greenMax) {
				if IsValidCubeCount(blueCubes, blueMax) {

					part1 += gameNum
				}
			}
		}

		// Part 2
		gamePower := GetPowerOfGame(redCubes, greenCubes, blueCubes)
		part2 += gamePower

	}

	fmt.Printf("Result - Part 1: %d, Part 2: %d", part1, part2)
}

// parses a line and returns game #, reds, greens, blues counts by pulls
func parseLine(line string) (int, []int, []int, []int, error) {

	splitted := strings.Split(line, ":")
	if len(splitted) != 2 {
		return -1, nil, nil, nil, fmt.Errorf("couldn't find ':' in line: %s", line)
	}

	gameNumStr := strings.Split(splitted[0], " ")
	if len(gameNumStr) == 0 {
		return -1, nil, nil, nil, fmt.Errorf("couldn't split game number '%s' in line: %s", splitted[0], line)
	}
	gameNum, err := strconv.Atoi(gameNumStr[1])
	if err != nil {
		return -1, nil, nil, nil, fmt.Errorf("couldn't convert game number '%s' in line: %s", gameNumStr, line)
	}

	redBalls := make([]int, 0, 10)
	greenBalls := make([]int, 0, 10)
	blueBalls := make([]int, 0, 10)
	pulls := strings.Split(splitted[1], ";")
	for _, pull := range pulls {
		red, green, blue, err := parseRedGreenBlue(pull)
		if err != nil {
			return gameNum, nil, nil, nil, fmt.Errorf("error parsing balls in line '%s': %w", line, err)
		}
		redBalls = append(redBalls, red)
		greenBalls = append(greenBalls, green)
		blueBalls = append(blueBalls, blue)
	}

	return gameNum, redBalls, greenBalls, blueBalls, nil
}

// returns red, green, blue for a pull
func parseRedGreenBlue(pullLine string) (int, int, int, error) {

	splitted := strings.Split(pullLine, ",")
	if len(splitted) == 0 {
		return -1, -1, -1, fmt.Errorf("couldn't find ',' in pull: %s", pullLine)
	}

	var red, green, blue int
	for _, cubes := range splitted {

		cubeSplit := strings.Split(strings.TrimSpace(cubes), " ")
		if len(cubeSplit) != 2 {
			return -1, -1, -1, fmt.Errorf("couldn't find ball count in: %s", cubes)
		}

		count, err := strconv.Atoi(cubeSplit[0])
		if err != nil {
			return -1, -1, -1, fmt.Errorf("invalid ball count '%s' in: %s", cubeSplit[0], cubes)
		}

		switch cubeSplit[1] {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count
		}

	}

	return red, green, blue, nil
}

func IsValidCubeCount(cubePulls []int, max int) bool {
	for _, count := range cubePulls {
		if count > max {
			return false
		}
	}
	return true
}

func GetPowerOfGame(redCubes, greenCubes, blueCubes []int) int {

	redCountMax := GetMaxCubes(redCubes)
	greenCountMax := GetMaxCubes(greenCubes)
	blueCountMax := GetMaxCubes(blueCubes)

	return redCountMax * greenCountMax * blueCountMax
}

func GetMaxCubes(cubeCounts []int) int {
	max := 0
	for _, count := range cubeCounts {
		if count > max {
			max = count
		}
	}
	return max
}
