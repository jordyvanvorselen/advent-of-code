package models

type Hand struct {
	Cards []Card
	Bid   int
}

func (h Hand) Type() HandType {
	raw := h.rawCards()
	charOccurs := countOccurrences(raw)

	if len(charOccurs) == 1 {
		return FiveOfAKind
	}

	if len(charOccurs) == 2 {
		if pot := charOccurs[raw[:1]]; pot == 1 || pot == 4 {
			return FourOfAKind
		}

		return FullHouse
	}

	if len(charOccurs) == 3 {
		for _, occurs := range charOccurs {
			if occurs == 3 {
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

			if res := cardB.isStrongerThan(cardA); res != 0 {
				return res == 1
			}
		}
	}

	return typeA < typeB
}
