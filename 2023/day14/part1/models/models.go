package models

type RockPositionList struct {
	Rows    []string
	Columns []string
}

func (rpl RockPositionList) TiltNorth() RockPositionList {
	var columns []string

	for colIdx, column := range rpl.Columns {
		newColumn := column

		for charIdx, c := range column {
			char := string(c)
			nextChar := string(column[charIdx+1])

			if char == "O" && nextChar == "." {
				shiftPlaces := freeFromIndex(column, charIdx)
				canShiftPlaces
			}
		}
	}

	return RockPositionList{Rows: []string{}, Columns: columns}
}

func freeFromIndex(line string, index int) int {
	var result int
	for _, c := range line[index+1:] {
		if string(c) == "." {
			result++
			continue
		}

		break
	}

	return result
}
