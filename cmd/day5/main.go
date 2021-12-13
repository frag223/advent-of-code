package main

import "fmt"

func main() {

	d := NewDiagram(10, true)
	d.PlaceLine(Position{
		X: 1,
		Y: 1,
	}, Position{X: 1, Y: 9})
	fmt.Println("--------")
	fmt.Println(d.Grid)
}
