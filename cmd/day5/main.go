package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {

	positions := parseInput(input)
	d := NewDiagram(1000, false)

	for _, line := range positions {
		d.PlaceLine(line[0], line[1])
	}

	fmt.Println(d.GetCrossedPositions())
	fmt.Println(len(d.GetCrossedPositions()))
}
