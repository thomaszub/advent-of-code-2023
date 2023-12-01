package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Execute() error {
	bytes, err := os.ReadFile("./input/day1.txt")
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
	sum := 0
	for _, l := range lines {
		sum += getLineValue(l)
	}
	return sum, nil
}

func getLineValue(line string) int {
	repLine := replaceWords(line)
	firstSet := false
	first := 0
	last := 0
	for _, r := range repLine {
		if !unicode.IsDigit(r) {
			continue
		}
		v, _ := strconv.Atoi(string(r))
		if !firstSet {
			firstSet = true
			first = v
		}
		last = v
	}
	return first*10 + last
}

var wordToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type indexedKey struct {
	key   string
	index int
}

type byIdx []indexedKey

func (b byIdx) Len() int {
	return len(b)
}

func (b byIdx) Less(i int, j int) bool {
	return b[i].index < b[j].index
}

func (b byIdx) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func replaceWords(line string) string {
	found := []indexedKey{}
	for k := range wordToDigit {
		idx := strings.Index(line, k)
		if idx > -1 {
			found = append(found, indexedKey{k, idx})
		}
	}
	if len(found) == 0 {
		return line
	}
	sort.Sort(byIdx(found))
	replLine := line
	for _, k := range found {
		replLine = replLine[:k.index] + wordToDigit[k.key] + replLine[k.index+1:]
	}
	return replaceWords(replLine)
}
