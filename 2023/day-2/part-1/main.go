package main

import (
	"bufio"
	"fmt"
	"os"
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
	id    int
	red   int
	green int
	blue  int
}

func main() {
	lines := readFile("2023/day-2/input/data")

	sum := run(lines)

	fmt.Println("Sum of game ids: ", sum)
}

func run(input []string) int {
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
	return rules.red >= game.red && rules.green >= game.green && rules.blue >= game.blue
}

func fromLine(line string) Game {
	groups := strings.FieldsFunc(line, Split)

	var id int
	var red int
	var green int
	var blue int

	for _, group := range groups {
		if strings.Contains(group, "Game") {
			id = extractDigits(group)
		}

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

	return Game{id, red, green, blue}
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

func Split(r rune) bool {
	return r == ':' || r == ',' || r == ';'
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
