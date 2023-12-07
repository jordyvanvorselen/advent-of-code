package models

type HandType int

const (
	FiveOfAKind  HandType = 6
	FourOfAKind  HandType = 5
	FullHouse    HandType = 4
	ThreeOfAKind HandType = 3
	TwoPair      HandType = 2
	OnePair      HandType = 1
	HighCard     HandType = 0
)
