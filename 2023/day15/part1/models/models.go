package models

type InitItem struct {
	Raw string
}

func (i InitItem) Hash() int {
	var result int

	for _, c := range i.Raw {
		result = ((result + int(c)) * 17) % 256
	}

	return result
}
