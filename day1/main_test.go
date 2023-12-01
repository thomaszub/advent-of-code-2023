package day1

import "testing"

func TestSolution(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	exp := 29 + 83 + 13 + 24 + 42 + 14 + 76
	act, err := solve(input)
	if err != nil {
		t.Errorf("Did not expect a error: %s", err.Error())
	}
	if act != exp {
		t.Errorf("Wrong result, expected: %d, actual: %d", exp, act)
	}
}

func TestGetLineValue(t *testing.T) {
	input := "aoneb2c3d4e5f"
	exp := 15
	act := getLineValue(input)
	if act != exp {
		t.Errorf("Wrong result, expected: %d, actual: %d", exp, act)
	}
}

func TestReplaceWords(t *testing.T) {
	input := "fivefive"
	exp := "5ive5ive"
	act := replaceWords(input)
	if act != exp {
		t.Errorf("Wrong result, expected: %s, actual: %s", exp, act)
	}
}
