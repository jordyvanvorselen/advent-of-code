package part1

func Run(input []string) int {
	var result int

	scratchCards := parse(input)

	for i, s := range scratchCards {
		n := s.GetWinningNumberCount()

		for i2, copiedCard := range scratchCards[i+1 : i+n+1] {
			scratchCards[i+1+i2].SetAmount(copiedCard.amount + s.amount)
		}
	}

	for _, s := range scratchCards {
		result = result + s.amount
	}

	return result
}
