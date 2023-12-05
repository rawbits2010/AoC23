package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rawbits2010/AoC23/internal/inputhandler"
	"golang.org/x/exp/maps"
)

var numStrs = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var numDigits = "0123456789"

func main() {

	lines := inputhandler.ReadInput()

	var result int
	for _, line := range lines {

		firstNum := findFirstNumber(line)
		//fmt.Printf("1st num: '%s'\n", firstNum)

		secondNum := findLastNumber(line)
		//fmt.Printf("2nd num: '%s'\n", firstNum)

		numberStr := firstNum + secondNum
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Printf("Couldn't convert '%s' to a number in line: %s", numberStr, line)
			os.Exit(-1)
		}
		//fmt.Printf("num '%s' conv: '%d'\n", numberStr, number)

		result += number
	}

	fmt.Printf("Result: %d", result)
}

func findFirstNumber(line string) string {

	var idx int

	numIdx := len(line)
	var numStr string
	for _, numstr := range maps.Keys(numStrs) {
		idx = strings.Index(line, numstr)
		if idx != -1 {
			//fmt.Printf("found '%s' at '%d' in %s\n", numstr, idx, line)

			if idx < numIdx {
				numIdx = idx
				numStr = numstr
			}
		}
	}

	digitIdx := strings.IndexAny(line, numDigits)
	if digitIdx != -1 {
		//fmt.Printf("found '%s' at '%d' in %s\n", string(line[digitIdx]), digitIdx, line)

		if digitIdx < numIdx {
			return string(line[digitIdx])
		}
	}

	return numStrs[numStr]
}

func findLastNumber(line string) string {

	var idx int

	numIdx := -1
	var numStr string
	for _, numstr := range maps.Keys(numStrs) {
		idx = strings.LastIndex(line, numstr)
		if idx != -1 {
			//fmt.Printf("found '%s' at '%d' in %s\n", numstr, idx, line)

			if idx > numIdx {
				numIdx = idx
				numStr = numstr
			}
		}
	}

	digitIdx := strings.LastIndexAny(line, numDigits)
	if digitIdx != -1 {
		//fmt.Printf("found '%s' at '%d' in %s\n", string(line[digitIdx]), digitIdx, line)

		if digitIdx > numIdx {
			return string(line[digitIdx])
		}
	}

	return numStrs[numStr]
}
