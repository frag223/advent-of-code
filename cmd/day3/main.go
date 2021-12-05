package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var input string

func main() {

	cache := map[int]map[int]int{}

	list := strings.Split(input, "\n")
	for _, bin := range list {
		for idx, r := range bin {
			pp, _ := strconv.Atoi(string(r))
			if _, ok := cache[idx][pp]; !ok {
				cache[idx] = map[int]int{
					pp: 0,
				}
				continue
			}

			cache[idx][pp] += 1

		}
	}
	fmt.Println(cache)
}

type Foo struct {
	Zero int
	One  int
}
