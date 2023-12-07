package models

import "fmt"

type Hand struct {
	Cards []Card
	Bid   int
}

func (h Hand) Type() HandType {
	raw := h.rawCards()
	charOccurs := countOccurrences(raw)
	jokerOccurs := countJokerOccurs(raw)

	delete(charOccurs, "J")

	if typeOfCards := len(charOccurs); typeOfCards == 1 || typeOfCards == 0 {
		return FiveOfAKind
	}

	if len(charOccurs) == 2 {
		for _, occur := range charOccurs {
			if occur+jokerOccurs == 4 {
				return FourOfAKind
			}
		}

		return FullHouse
	}

	if len(charOccurs) == 3 {
		for _, occurs := range charOccurs {
			if occurs+jokerOccurs == 3 {
				return ThreeOfAKind
			}
		}
	}

	pairs := countPairs(raw)

	if pairs == 2 {
		return TwoPair
	}

	if pairs == 1 {
		return OnePair
	}

	// It's not possible to have a TwoPair when having a joker.
	// At most, you have nothing + a joker which makes it OnePair.
	if jokerOccurs > 0 {
		return OnePair
	}

	return HighCard
}

func (h Hand) rawCards() string {
	var result string
	for _, c := range h.Cards {
		result = result + c.Label
	}

	return result
}

func countOccurrences(input string) map[string]int {
	charCount := make(map[string]int)

	for _, char := range input {
		charCount[string(char)]++
	}

	return charCount
}

func countJokerOccurs(input string) int {
	count := 0
	for _, c := range input {
		if string(c) == "J" {
			count++
		}
	}
	return count
}

func countPairs(input string) int {
	var count int

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				count++
			}
		}
	}

	return count
}

type Hands []Hand

func (hs Hands) Len() int {
	return len(hs)
}

func (hs Hands) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func (hs Hands) Less(i, j int) bool {
	typeA := hs[i].Type()
	typeB := hs[j].Type()

	if typeA == typeB {
		for idx, cardA := range hs[i].Cards {
			cardB := hs[j].Cards[idx]

			if cardA.Label == "J" && cardB.Label == "K" {
				fmt.Println(".")
			}

			if res := cardB.isStrongerThan(cardA); res != 0 {
				return res == 1
			}
		}
	}

	return typeA < typeB
}
