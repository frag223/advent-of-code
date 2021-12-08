package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed test.txt
var input string

func main() {

	data := extract(input)
	tfm := transform(data)
	gamma := getGamma(tfm)
	epsilon := getEpsilon(tfm)
	fmt.Println(gamma)
	fmt.Println(epsilon)
	result := productBinary(gamma, epsilon)
	fmt.Println(result)
}

type BinaryDiagnostic struct {
	Order int
	Zero  int
	One   int
}

type Collection []BinaryDiagnostic

func (a Collection) Len() int           { return len(a) }
func (a Collection) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Collection) Less(i, j int) bool { return a[i].Order < a[j].Order }

func extract(input string) map[int]map[int]int {
	cache := map[int]map[int]int{}
	list := strings.Split(input, "\n")
	for _, bin := range list {
		for idx, r := range bin {
			pp, _ := strconv.Atoi(string(r))
			if _, ok := cache[idx][pp]; !ok {
				cache[idx] = map[int]int{
					0: 0,
					1: 0,
				}
			}

			cache[idx][pp] += 1

		}
	}
	return cache
}

func transform(data map[int]map[int]int) Collection {
	var col Collection
	for key, value := range data {
		v := BinaryDiagnostic{
			Order: key,
			Zero:  value[0],
			One:   value[1],
		}
		col = append(col, v)
	}
	sort.Sort(col)
	return col
}

func getGamma(data Collection) string {
	binary := ""
	for _, item := range data {
		if item.One > item.Zero {
			binary += fmt.Sprintf("%d", 1)
			continue
		}
		binary += fmt.Sprintf("%d", 0)
	}
	return binary
}

func getEpsilon(data Collection) string {
	binary := ""
	for _, item := range data {
		if item.One > item.Zero {
			binary += fmt.Sprintf("%d", 0)
			continue
		}
		binary += fmt.Sprintf("%d", 1)
	}
	return binary
}

func productBinary(first, second string) int64 {
	f, _ := strconv.ParseInt(first, 2, 64)
	s, _ := strconv.ParseInt(second, 2, 64)
	return f * s
}
