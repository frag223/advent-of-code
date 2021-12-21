package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var input string

func main() {
	fishData := parseInput(input)
	days := len(fishData) - 1
	newData := [][]int{fishData[0]}
	for i := 0; i < days; i++ {
		row := calculateNextRow(newData[i])
		newData = append(newData, row)
	}
	fmt.Println(len(newData))
	newTotal := len(newData[len(newData)-1])
	fmt.Println(newTotal)
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
	nextRow = append(nextRow, newFish...)
	return nextRow
}
