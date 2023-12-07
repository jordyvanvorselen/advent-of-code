package parsers

import (
	"2023/day7/part2/models"
	"2023/day7/part2/utils"
	"strings"
)

func Parse(lines []string) models.Hands {
	var result []models.Hand

	for _, line := range lines {
		words := strings.Fields(line)

		bid := utils.ToInt(words[1])
		rawCards := words[0]

		var handCards []models.Card

		for _, rc := range rawCards {
			handCards = append(handCards, models.Card{Label: string(rc)})
		}

		result = append(result, models.Hand{Cards: handCards, Bid: bid})
	}

	return result
}
