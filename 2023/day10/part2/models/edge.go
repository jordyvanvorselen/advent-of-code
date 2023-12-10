package models

type Edge struct {
	To, Weight int
}

func EdgeToRight(index int) Edge {
	return Edge{To: index + 1, Weight: 1}
}

func EdgeToLeft(index int) Edge {
	return Edge{To: index - 1, Weight: 1}
}

func EdgeToTop(index int, lineLength int) Edge {
	return Edge{To: index - lineLength, Weight: 1}
}

func EdgeToBottom(index int, lineLength int) Edge {
	return Edge{To: index + lineLength, Weight: 1}
}
