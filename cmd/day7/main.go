package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed test.txt
var input string

func main() {
	pos := parseInput(input)
	sort.Ints(pos)

	lastPos := pos[len(pos)-1]

}

func parseInput(input string) []int {
	list := strings.Split(input, ",")

	var result []int
	for _, item := range list {

		i, _ := strconv.Atoi(item)
		result = append(result, i)
	}
	return result
}
