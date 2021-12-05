package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	problem1()
	problem2()

}

func problem1() {
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

func problem2() {
	count := 0
	lastVal := 0
	results := []int{}
	list := strings.Split(input, "\n")

	for i := 0; i+3 < len(list); i += 2 {
		j := i + 1
		val1, _ := strconv.Atoi(list[i])
		val2, _ := strconv.Atoi(list[i+1])
		val3, _ := strconv.Atoi(list[i+2])
		results = append(results, sum(val1, val2, val3))
		val4, _ := strconv.Atoi(list[j])
		val5, _ := strconv.Atoi(list[j+1])
		val6, _ := strconv.Atoi(list[j+2])
		results = append(results, sum(val4, val5, val6))
	}

	for _, item := range results {
		if lastVal == 0 {
			lastVal = item
			continue
		}
		if item > lastVal {
			count++
		}
		lastVal = item
	}

	fmt.Println(count)
}

func sum(list ...int) int {
	result := 0
	for _, item := range list {
		result += item
	}
	return result
}
