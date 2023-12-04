package part1

type ScratchCard struct {
	amount  int
	winning []int
	numbers []int
}

func (s *ScratchCard) GetWinningNumberCount() int {
	return bOverlapCountInA(s.winning, s.numbers)
}

func (s *ScratchCard) SetAmount(i int) {
	s.amount = i
}
