package part2

import (
	"2023/day15/part2/models"
	"2023/day15/part2/parsers"
	"github.com/samber/lo"
	"slices"
)

func Run(input []string) int {
	var result int

	items := parsers.Parse(input)

	boxes := make(map[int][]models.InitItem)

	for _, item := range items {
		box := item.Hash()

		if item.Op == models.REMOVE {
			for i, val := range boxes[box] {
				if val.Label == item.Label {
					boxes[box] = slices.Delete(boxes[box], i, i+1)
				}
			}

			continue
		}

		_, idx, found := lo.FindIndexOf(boxes[box], func(ii models.InitItem) bool {
			return ii.Label == item.Label
		})

		if found {
			boxes[box][idx] = item
		} else {
			boxes[box] = append(boxes[box], item)
		}
	}

	for i, box := range boxes {
		count := 0

		for _, item := range box {
			count++
			result += (i + 1) * count * item.FocalLength
		}
	}

	return result
}
