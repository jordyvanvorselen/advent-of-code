package utils

type String string

func (s String) ReplaceAt(index int, new string) string {
	var newString string

	for i, c := range s {
		if index == i {
			newString += new
		} else {
			newString += string(c)
		}
	}

	return newString
}
