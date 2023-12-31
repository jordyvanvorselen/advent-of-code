package part1

import (
	"strconv"
	"strings"
	"unicode"
)

type Rules struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	groups []Group
}

type Group struct {
	red   int
	green int
	blue  int
}

func Run(input []string) int {
	rules := Rules{red: 12, green: 13, blue: 14}
	var games []Game
	var result int

	for _, line := range input {
		games = append(games, fromLine(line))
	}

	for _, game := range games {
		if rules.allow(game) {
			result = result + game.id
		}
	}

	return result
}

func (rules Rules) allow(game Game) bool {
	for _, group := range game.groups {
		if rules.red < group.red || rules.green < group.green || rules.blue < group.blue {
			return false
		}
	}

	return true
}

func fromLine(line string) Game {
	splitLine := strings.Split(line, ":")

	gameInfo := splitLine[0]
	groupInfo := splitLine[1]

	rawGroups := strings.Split(groupInfo, ";")

	var finalGroups []Group
	id := extractDigits(gameInfo)

	for _, rawGroup := range rawGroups {
		groups := strings.Split(rawGroup, ",")

		var red int
		var green int
		var blue int

		for _, group := range groups {
			if strings.Contains(group, "red") {
				red = red + extractDigits(group)
			}

			if strings.Contains(group, "green") {
				green = green + extractDigits(group)
			}

			if strings.Contains(group, "blue") {
				blue = blue + extractDigits(group)
			}
		}

		finalGroups = append(finalGroups, Group{red, green, blue})
	}

	return Game{id: id, groups: finalGroups}
}

func extractDigits(input string) int {
	var result string
	for _, char := range input {
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}

	value, _ := strconv.Atoi(result)
	return value
}
