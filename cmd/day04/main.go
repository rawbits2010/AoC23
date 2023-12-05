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

	var part1 int

	cardsCount := make([]int, len(lines))

	for idx, line := range lines {

		cardNum, winList, ownList, err := parseCard(line)
		if err != nil {
			fmt.Printf("error processing line '%s': %s\n", line, err)
			os.Exit(-1)
		}

		matches := GetMatches(winList, ownList)
		fmt.Printf("card: %d - wins: %d\n", cardNum, matches)

		score := int(math.Pow(2, float64(matches-1)))
		part1 += score

		cardsCount[idx]++
		for i := 1; i <= matches; i++ {
			if idx+i > len(lines)-1 {
				break
			}
			cardsCount[idx+i] += cardsCount[idx]
		}

	}

	var part2 int
	for _, count := range cardsCount {
		part2 += count
	}

	fmt.Printf("Result - Part 1: %d, Part 2: %d", part1, part2)
}

// parses card input line as card number, winning numbers, your numbers
func parseCard(line string) (int, []int, []int, error) {

	splitted := strings.Split(line, ":")
	if len(splitted) != 2 {
		return -1, nil, nil, fmt.Errorf("couldn't split at ':'")
	}

	cardNumStr := strings.TrimSpace(splitted[0][4:])
	cardNum, err := strconv.Atoi(cardNumStr)
	if err != nil {
		return -1, nil, nil, fmt.Errorf("couldn't convert card number '%d': %w", cardNum, err)
	}

	numberSplit := strings.Split(splitted[1], "|")
	if len(numberSplit) != 2 {
		return -1, nil, nil, fmt.Errorf("couldn't split at '|'")
	}

	winNums, err := parseTheNumbers(numberSplit[0])
	if err != nil {
		return -1, nil, nil, fmt.Errorf("couldn't parse winning numbers '%s': %w", numberSplit[0], err)
	}

	ownNums, err := parseTheNumbers(numberSplit[1])
	if err != nil {
		return -1, nil, nil, fmt.Errorf("couldn't parse own numbers '%s': %w", numberSplit[1], err)
	}

	return cardNum, winNums, ownNums, nil
}

func parseTheNumbers(numbersStr string) ([]int, error) {

	numbers := make([]int, 0)
	split := strings.Split(strings.TrimSpace(numbersStr), " ")
	for _, numStr := range split {
		if len(numStr) == 0 {
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			return nil, fmt.Errorf("couldn't convert winning number '%s' in numbers '%s': %w", numStr, numbersStr, err)
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}

func GetMatches(winNums, ownNums []int) int {

	var match int
	for _, w := range winNums {
		for _, o := range ownNums {
			if w == o {
				match++
				break
			}
		}
	}

	return match
}
