package part1

func Run(input []string) int {
	var result int
	scratchCards := parse(input)

	for _, s := range scratchCards {
		result = result + s.getScore()
	}

	return result
}
