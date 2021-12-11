package main

import (
	"fmt"

	"github.com/frag223/advent-of-code/internal/bingo"
)

func main() {
	v := bingo.New(5)
	fmt.Println(v)
	v.Add(1)
	v.Add(2)
	v.Add(3)
	v.Add(4)
	v.Add(5)
	v.Add(6)
	v.Add(7)
	v.Mark(1)
	v.Print()

}
