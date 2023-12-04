package part1

func Run(input []string) int {
	var result int

	scratchCards := parse(input)

	for i, s := range scratchCards {
		n := s.GetWinningNumberCount()

		// 0 card 1 -> 2     :  1
		// 1 card 2 -> 1     :  2
		// 2 card 3 -> 1     :  4
		// 3 card 4 -> 0     :  5

		// idx 3, n = 5
		// 4, 5, 6, 7, 8
		// start = i + 1
		// end = i + 5
		for i2, copiedCard := range scratchCards[i+1 : i+n+1] {
			scratchCards[i+1+i2].SetAmount(copiedCard.amount + s.amount)
		}
	}

	for _, s := range scratchCards {
		result = result + s.amount
	}

	return result
}
