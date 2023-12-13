package models

type MirrorList struct {
	Rows    []string
	Columns []string
}

func (m MirrorList) FindPerfectRowReflectionIndex() (int, bool) {
	return findPerfectReflectionIndex(m.Rows)
}

func (m MirrorList) FindPerfectColumnReflectionIndex() (int, bool) {
	return findPerfectReflectionIndex(m.Columns)
}

func findPerfectReflectionIndex(lines []string) (int, bool) {
	for i, line := range lines {
		if i+1 == len(lines) {
			break
		}

		nextLine := lines[i+1]

		if line == nextLine {
			if isPerfectReflection(lines, i) {
				return i, true
			}
		}
	}

	return -1, false
}

func isPerfectReflection(lines []string, index int) bool {
	indexA := index
	indexB := index + 1

	for {
		if indexA < 0 {
			break
		}

		if indexB >= len(lines) {
			break
		}

		if lines[indexA] != lines[indexB] {
			return false
		}

		indexA--
		indexB++
	}

	return true
}
