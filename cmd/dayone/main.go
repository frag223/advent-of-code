package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	count := 0
	lastValue := 0
	list := strings.Split(input, "\n")
	for _, item := range list {
		if item == "" {
			continue
		}
		parsedItem, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalln("error parsing item ", err)
		}
		if lastValue == 0 {
			lastValue = parsedItem
			continue
		}
		if parsedItem > lastValue {
			count++
		}
		lastValue = parsedItem
	}
	log.Println(count)
}
