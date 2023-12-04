package part1

type ScratchCard struct {
	winning []int
	numbers []int
}

func (s ScratchCard) getScore() int {
	intersects := bOverlapCountInA(s.winning, s.numbers)

	result := 0
	for i := 0; i < intersects; i++ {
		if i == 0 {
			result++
			continue
		}

		result = result * 2
	}

	return result
}
