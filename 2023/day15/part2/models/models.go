package models

type InitItem struct {
	Label       string
	Op          Operation
	FocalLength int
}

type Operation string

const (
	ADD    Operation = "="
	REMOVE Operation = "-"
)

func (i InitItem) Hash() int {
	var result int

	for _, c := range i.Label {
		result = ((result + int(c)) * 17) % 256
	}

	return result
}
