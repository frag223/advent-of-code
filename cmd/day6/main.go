package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	// fishData := parseText(input)
	fishData := parseInput(input)
	newData := [][]int{fishData[0]}
	days := 80
	for i := 0; i < days; i++ {
		row := calculateNextRow(newData[i])
		newData = append(newData, row)
	}
	fmt.Println(len(newData[len(newData)-1]))
}

func calculateNextState(current int) (int, bool) {
	switch current {
	case 0:
		return 6, true
	default:
		return current - 1, false
	}
}

func calculateNextRow(row []int) []int {
	newFish := []int{}
	nextRow := []int{}
	for _, item := range row {
		next, add := calculateNextState(item)
		if add {
			newFish = append(newFish, 8)
		}
		nextRow = append(nextRow, next)
	}
	if len(newFish) > 0 {
		nextRow = append(nextRow, newFish...)
	}
	return nextRow
}
