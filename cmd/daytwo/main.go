package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed test.txt
var input string

func main() {
	cache := map[string][]int{}
	list := strings.Split(input, "\n")
	for _, item := range list {
		if item == "" {
			continue
		}
		measurements := strings.Split(item, " ")
		value, err := strconv.Atoi(measurements[0])
		if err != nil {
			log.Println("invalid number: ", measurements[0])
			continue
		}
		for i := 1; i < len(measurements); i++ {
			if measurements[i] != "" {
				cache[measurements[i]] = append(cache[measurements[i]], value)
			}
		}

	}
	fmt.Println(cache)
}
