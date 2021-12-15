package main

import (
	"strconv"
	"strings"
)

func parseInput(input string) [][]Position {
	list := strings.Split(input, "\n")
	ff := [][]Position{}
	for _, row := range list {
		positionsRaw := strings.Split(row, "->")
		firstPosition := strings.Split(
			strings.TrimSpace(positionsRaw[0]),
			",",
		)
		secondPosition := strings.Split(
			strings.TrimSpace(positionsRaw[1]),
			",",
		)
		x1, _ := strconv.Atoi(firstPosition[0])
		y1, _ := strconv.Atoi(firstPosition[1])
		x2, _ := strconv.Atoi(secondPosition[0])
		y2, _ := strconv.Atoi(secondPosition[1])
		ff = append(ff, []Position{
			{
				X: x1,
				Y: y1,
			},
			{
				X: x2,
				Y: y2,
			},
		})
	}
	return ff
}
