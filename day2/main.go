package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}

func Execute() error {
	bytes, err := os.ReadFile("./input/day2.txt")
	if err != nil {
		return err
	}
	text := string(bytes)
	lines := strings.Split(text, "\n")
	res, err := solve(lines)
	if err != nil {
		return err
	}
	fmt.Printf("The solution is: %d\n", res)
	return nil
}

func solve(lines []string) (int, error) {
	result := 0
	for _, l := range lines {
		g, err := parse(l)
		if err != nil {
			return result, err
		}
		result += power(g.sets)
	}
	return result, nil
}

func power(sets []set) int {
	red := 0
	blue := 0
	green := 0
	for _, s := range sets {
		red = int(math.Max(float64(red), float64(s.red)))
		blue = int(math.Max(float64(blue), float64(s.blue)))
		green = int(math.Max(float64(green), float64(s.green)))
	}
	return red * blue * green
}

func parse(line string) (game, error) {
	game := game{
		id:   0,
		sets: []set{},
	}

	sp := strings.Split(line, ": ")
	id, err := strconv.Atoi(strings.Split(sp[0], " ")[1])
	if err != nil {
		return game, err
	}
	game.id = id

	setStrings := strings.Split(sp[1], ";")
	for _, s := range setStrings {
		set, err := parseSet(s)
		if err != nil {
			return game, err
		}
		game.sets = append(game.sets, set)
	}
	return game, nil
}

func parseSet(line string) (set, error) {
	val := set{0, 0, 0}
	colors := strings.Split(line, ",")
	for _, c := range colors {
		kv := strings.Split(strings.TrimSpace(c), " ")
		v, err := strconv.Atoi(kv[0])
		if err != nil {
			return val, err
		}
		if kv[1] == "green" {
			val.green = v
		}
		if kv[1] == "blue" {
			val.blue = v
		}
		if kv[1] == "red" {
			val.red = v
		}
	}
	return val, nil
}
