package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fishData := parseInput(input)

	// part1 := predictPopulation(fishData, 80)
	// fmt.Printf("part 1: %d\n", part1)

	part2 := predictPopulation(fishData, 256)
	fmt.Printf("part 2: %d\n", part2)
}

// NOTE did not solve this

// Use a circular buffer to keep track of the population of fishes.
// At each generation, fishes who reach a 0 lifespan give birth to new fishes
// (with life=8), and their life is reset to 6.
//
// Initial state:
// life:  0   1   2   3   4   5   6   7   8
// count: 0   1   1   2   1   0   0   0   0
//
// 1st generation
// life:  8   0   1   2   3   4   5   6   7
// count: 0   1   1   2   1   0   0   0   0
//
// 2nd generation
// life:  7   8   0   1   2   3   4   5   6
// count: 0  (1)  1   2   1   0   0   0   0+1
//
// 3rd generation
// life:  6   7   8   0   1   2   3   4   5
// count: 0+1 1  (1)  2   1   0   0   0   1
//
// 4th generation
// life:  5   6   7   8   0   1   2   3   4
// count: 1   1+2 1  (2)  1   0   0   0   1
func predictPopulation(input []uint64, gens int) uint64 {
	var buf [9]uint64
	copy(buf[:], input)

	born, reset := 8, 6
	for i := 0; i < gens; i++ {
		born, reset = incr(born), incr(reset)
		buf[reset] += buf[born]
	}

	var result uint64
	for _, count := range buf {
		result += count
	}
	return result
}

func incr(index int) int {
	if index == 8 {
		return 0
	}
	return index + 1
}

func parseInput(line string) []uint64 {
	inputStrings := strings.Split(line, ",")
	var input [9]uint64
	for _, s := range inputStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		input[n]++
	}
	return input[:]
}
