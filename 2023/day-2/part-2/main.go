package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	id     int
	groups []Group
}

type Group struct {
	red   int
	green int
	blue  int
}

func main() {
	lines := readFile("2023/day-2/input/data")

	sum := run(lines)

	fmt.Println("Sum of power of sets: ", sum)
}

func run(input []string) int {
	var games []Game
	var result int

	for _, line := range input {
		games = append(games, fromLine(line))
	}

	for _, game := range games {
		maxSet := game.findMaxSet()

		result = result + (maxSet.red * maxSet.green * maxSet.blue)
	}

	return result
}

func (game Game) findMaxSet() Group {
	var red int
	var green int
	var blue int

	for _, group := range game.groups {
		if group.red > red {
			red = group.red
		}

		if group.green > green {
			green = group.green
		}

		if group.blue > blue {
			blue = group.blue
		}
	}

	return Group{red, green, blue}
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

func readFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
