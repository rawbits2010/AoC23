package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/rawbits2010/AoC23/internal/inputhandler"
)

type Card int

const (
	CardInv Card = iota
	Card2
	Card3
	Card4
	Card5
	Card6
	Card7
	Card8
	Card9
	CardT
	CardJ
	CardQ
	CardK
	CardA
)
const CardMaxVal = int(CardA)

type Hand struct {
	Cards []Card
	Type  HandType
	Bid   int
}

type HandType int

const (
	HTInv HandType = iota
	HTHighCard
	HTOnePair
	HTTwoPair
	HTThreeOfAKind
	HTFullHouse
	HTFourOfAKind
	HTFiveOfKind
)

func main() {

	lines := inputhandler.ReadInput()

	hands := make([]Hand, 0)
	for _, line := range lines {
		cards, bid, err := parseHand(line)
		if err != nil {
			fmt.Printf("error parsing line '%s': %s", line, err)
			os.Exit(-1)
		}
		hands = append(hands, Hand{Cards: cards, Bid: bid})
	}

	for handIdx := 0; handIdx < len(hands); handIdx++ {
		handType, err := determineHand(hands[handIdx].Cards)
		if err != nil {
			fmt.Printf("error determining hand type: '%s': %s", lines[handIdx], err)
			os.Exit(-1)
		}
		hands[handIdx].Type = handType
	}

	slices.SortFunc(hands, handSortingFunc)

	var part1 int
	for handIdx, currHand := range hands {
		part1 += currHand.Bid * (handIdx + 1)
	}

	fmt.Printf("Result - Part1: %d", part1)
}

func handSortingFunc(p, q Hand) int {
	if p.Type == q.Type {
		for cardIdx := 0; cardIdx < len(p.Cards); cardIdx++ {
			comp := compareCards(p.Cards[cardIdx], q.Cards[cardIdx])
			if comp == 0 {
				continue
			}
			return comp
		}
		fmt.Print("we got a tie")
		os.Exit(-1)
	}
	if p.Type < q.Type {
		return -1
	}
	return 1
}

// returns -1 if right is greater, 0 if equal, 1 if left is greater
func compareCards(f, s Card) int {
	if f == s {
		return 0
	}
	if f < s {
		return -1
	}
	return 1
}

func determineHand(cards []Card) (HandType, error) {

	lastCardCount := 0
	waiting := false
	for cardVal := 1; cardVal <= CardMaxVal; cardVal++ {
		cardCount := 0
		for _, currCard := range cards {
			if int(currCard) == cardVal {
				cardCount++
			}
		}

		switch cardCount {
		case 5:
			return HTFiveOfKind, nil

		case 4:
			return HTFourOfAKind, nil

		case 3:
			if waiting {
				if lastCardCount == 2 {
					return HTFullHouse, nil
				}
			}
			lastCardCount = cardCount
			waiting = true

		case 2:
			if waiting {
				if lastCardCount == 3 {
					return HTFullHouse, nil
				}
				if lastCardCount == 2 {
					return HTTwoPair, nil
				}
			}
			lastCardCount = cardCount
			waiting = true

		}

		if !waiting {
			lastCardCount = cardCount
		}
	}

	switch lastCardCount {
	case 3:
		return HTThreeOfAKind, nil
	case 2:
		return HTOnePair, nil
	}

	return HTHighCard, nil
}

// parses a hand and returns the cards in the parsed order and the bid
func parseHand(line string) ([]Card, int, error) {

	lineSplit := strings.Split(line, " ")
	if len(lineSplit) != 2 {
		return nil, 0, fmt.Errorf("invalid line format")
	}

	hand := make([]Card, 0, 5)
	for _, card := range lineSplit[0] {
		cardConv, err := convertToCard(card)
		if err != nil {
			return nil, 0, fmt.Errorf("error processing hand: %w", err)
		}
		hand = append(hand, cardConv)
	}

	bid, err := strconv.Atoi(lineSplit[1])
	if err != nil {
		return nil, 0, fmt.Errorf("error converting bid '%s'", lineSplit[1])
	}

	return hand, bid, nil
}

func convertToCard(card rune) (Card, error) {
	switch card {
	case 'A':
		return CardA, nil
	case 'K':
		return CardK, nil
	case 'Q':
		return CardQ, nil
	case 'J':
		return CardJ, nil
	case 'T':
		return CardT, nil
	case '9':
		return Card9, nil
	case '8':
		return Card8, nil
	case '7':
		return Card7, nil
	case '6':
		return Card6, nil
	case '5':
		return Card5, nil
	case '4':
		return Card4, nil
	case '3':
		return Card3, nil
	case '2':
		return Card2, nil
	default:
		return CardInv, fmt.Errorf("invalid card type: %c", card)
	}
}
