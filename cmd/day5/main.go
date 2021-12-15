package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var input string

func main() {

	i := parseInput(input)

	fmt.Println(i)

	d := NewDiagram(10, true)
	d.PlaceLine(Position{
		X: 1,
		Y: 1,
	}, Position{X: 1, Y: 9})
	fmt.Println("--------")
}
