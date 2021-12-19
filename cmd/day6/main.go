package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var input string

func main() {
	fishData := parseInput(input)
	fmt.Println(fishData)
}
