package main

import (
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	list := strings.Split(input, "\n")
	convertedData := [][]int{}
	for _, row := range list {
		fishData := rowToFishes(row)
		convertedData = append(convertedData, fishData)
	}
	return convertedData
}

func rowToFishes(rawMeasurement string) []int {
	dirtyData := strings.Split(rawMeasurement, ":")
	rawFishes := strings.Split(
		strings.TrimSpace(dirtyData[1]),
		",",
	)
	var fishData []int
	for _, item := range rawFishes {
		fish, _ := strconv.Atoi(item)
		fishData = append(fishData, fish)
	}
	return fishData
}
