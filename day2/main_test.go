package day2

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	exp := 2286
	act, err := solve(input)
	if err != nil {
		t.Errorf("Error not expected: %s", err.Error())
	}
	if act != exp {
		t.Errorf("Wrong result, expected: %d, actual: %d", exp, act)
	}
}

func TestParse(t *testing.T) {
	input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	exp := game{
		id: 3,
		sets: []set{
			{
				red:   20,
				green: 8,
				blue:  6,
			},
			{
				red:   4,
				green: 13,
				blue:  5,
			},
			{
				red:   1,
				green: 5,
				blue:  0,
			},
		},
	}
	act, err := parse(input)
	if err != nil {
		t.Errorf("Got error parsing input: %s", err.Error())
	}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf("Wrong result, expected: %d, actual: %d", exp, act)
	}
}
