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
	horizontal := 0
	depth := 0
	aim := 0

	instructions := strings.Split(input, "\n")
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}
		items := strings.Split(instruction, " ")
		direction := items[0]
		value, _ := strconv.Atoi(items[1])

		switch direction {
		case "forward":
			horizontal += value
			depth += (aim * value)
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			continue
		}
	}
	journey := horizontal * depth
	fmt.Println("journey", journey)
}
