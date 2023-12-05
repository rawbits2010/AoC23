package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rawbits2010/AoC23/internal/inputhandler"
)

// numbers 0-9: 48-57 byte code
// '.' is 46

type Part struct {
	partNumberStr string
	partNumber    int
	symbolPosX    int
	symbolPosY    int
}

func main() {

	lines := inputhandler.ReadInput()

	partList := make([]Part, 0)

	for lineIdx, line := range lines {

		currNumber := ""
		hasSymbol := false
		var symX, symY int
		for charIdx, char := range line {

			// a number
			if IsNumber(byte(char)) {

				// check surroundings
				if !hasSymbol {

					// need left side
					if len(currNumber) == 0 && charIdx > 0 {

						// top-left
						if lineIdx > 0 {
							if IsSymbol(lines[lineIdx-1][charIdx-1]) {
								hasSymbol = true
								symX = charIdx - 1
								symY = lineIdx - 1
								goto symbol_found
							}
						}

						// left
						if IsSymbol(lines[lineIdx][charIdx-1]) {
							hasSymbol = true
							symX = charIdx - 1
							symY = lineIdx
							goto symbol_found
						}

						// bottom-left
						if lineIdx < len(lines)-1 {
							if IsSymbol(lines[lineIdx+1][charIdx-1]) {
								hasSymbol = true
								symX = charIdx - 1
								symY = lineIdx + 1
								goto symbol_found
							}
						}
					}

					// above
					if lineIdx > 0 {
						if IsSymbol(lines[lineIdx-1][charIdx]) {
							hasSymbol = true
							symX = charIdx
							symY = lineIdx - 1
							goto symbol_found
						}
					}

					// below
					if lineIdx < len(lines)-1 {
						if IsSymbol(lines[lineIdx+1][charIdx]) {
							hasSymbol = true
							symX = charIdx
							symY = lineIdx + 1
							goto symbol_found
						}
					}

					// need right side
					if charIdx < len(line)-1 && !IsNumber(lines[lineIdx][charIdx+1]) {

						// top-left
						if lineIdx > 0 {
							if IsSymbol(lines[lineIdx-1][charIdx+1]) {
								hasSymbol = true
								symX = charIdx + 1
								symY = lineIdx - 1
								goto symbol_found
							}
						}

						// left
						if IsSymbol(lines[lineIdx][charIdx+1]) {
							hasSymbol = true
							symX = charIdx + 1
							symY = lineIdx
							goto symbol_found
						}

						// bottom-left
						if lineIdx < len(lines)-1 {
							if IsSymbol(lines[lineIdx+1][charIdx+1]) {
								hasSymbol = true
								symX = charIdx + 1
								symY = lineIdx + 1
								goto symbol_found
							}
						}

					}

				}

			symbol_found:
				currNumber += string(char)

				if len(currNumber) > 0 {
					// end of this number
					if charIdx == len(line)-1 {

						if hasSymbol {
							partList = append(partList, Part{partNumberStr: currNumber, symbolPosX: symX, symbolPosY: symY})
						}
						currNumber = ""
						hasSymbol = false
					}
				}
			} else {

				if len(currNumber) > 0 {
					// end of this number
					if hasSymbol {
						partList = append(partList, Part{partNumberStr: currNumber, symbolPosX: symX, symbolPosY: symY})
					}
					currNumber = ""
					hasSymbol = false
				}
			}

		}
	}

	var part1 int
	for partIdx, part := range partList {
		num, err := strconv.Atoi(part.partNumberStr)
		if err != nil {
			fmt.Printf("couldn't parse part number: %s", part.partNumber)
			os.Exit(-1)
		}
		partList[partIdx].partNumber = num

		part1 += num
	}

	var part2 int
	for i := 0; i < len(partList); i++ {
		currPart := partList[i]

		// has *
		if lines[currPart.symbolPosY][currPart.symbolPosX] == 42 {

			// is there a pair?
			for j := i + 1; j < len(partList); j++ {
				currToTest := partList[j]

				if currPart.symbolPosX == currToTest.symbolPosX && currPart.symbolPosY == currToTest.symbolPosY {

					part2 += currPart.partNumber * currToTest.partNumber

					break // assume can be only one
				}
			}
		}
	}

	fmt.Printf("Result - Part 1: %d, Part 2: %d", part1, part2)
}

func IsNumber(char byte) bool {
	if char < 48 || char > 57 {
		return false
	}
	return true
}

func IsSymbol(char byte) bool {
	if IsNumber(char) {
		return false
	}
	if char == 46 {
		return false
	}
	return true
}
