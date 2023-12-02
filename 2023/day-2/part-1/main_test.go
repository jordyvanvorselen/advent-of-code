package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 57
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", // No
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", // No
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		"Game 10: 2 blue; 10 green, 4 blue, 3 red; 5 green, 4 red, 4 blue; 1 red, 3 blue, 4 green; 2 blue, 5 red, 3 green; 3 green, 2 red, 2 blue",
		"Game 11: 4 blue, 19 green; 19 blue, 12 green, 17 red; 11 red, 10 blue, 17 green; 9 green, 18 blue; 14 green, 9 red, 18 blue; 15 blue, 6 green, 19 red", // No
		"Game 12: 1 green, 6 blue, 2 red; 6 blue, 2 red, 8 green; 2 green, 2 red, 7 blue; 1 red, 3 blue, 6 green",
		"Game 13: 2 red, 11 blue, 4 green; 2 red, 7 blue; 9 green, 1 red, 12 blue; 13 blue, 8 green; 11 blue, 8 green, 1 red; 1 red, 2 blue",
		"Game 14: 1 green, 4 blue, 11 red; 11 green, 6 blue, 7 red; 7 green, 6 blue, 4 red; 12 blue, 10 red, 11 green",
		"Game 15: 1 green, 19 red, 3 blue; 11 red, 3 blue; 20 red, 4 blue", // No
	}

	result := run(input)

	if result != expected {
		t.Errorf("got %d but expected %d", result, expected)
	}
}
