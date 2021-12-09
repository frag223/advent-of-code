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
	list := strings.Split(input, "\n")

	data := extract(list)
	tfm := transform(data)
	gamma := getGamma(tfm)
	epsilon := getEpsilon(tfm)
	result := productBinary(gamma, epsilon)
	fmt.Println(result)
	foo := getOxygen(list)
	poo := getC02(list)

	fmt.Println(foo)
	fmt.Println(poo)

	lifeSupport := productBinary(foo, poo)

	fmt.Println(lifeSupport)

}

type BinaryDiagnostic struct {
	Order int
	Zero  int
	One   int
}

func (b *BinaryDiagnostic) MostCommon() int {
	if b.Zero > b.One {
		return 0
	}
	return 1
}

func (b *BinaryDiagnostic) LeastCommon() int {
	if b.Zero > b.One {
		return 1
	}
	return 0
}

type Collection []BinaryDiagnostic

func (a Collection) Len() int           { return len(a) }
func (a Collection) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Collection) Less(i, j int) bool { return a[i].Order < a[j].Order }

func extract(list []string) map[int]map[int]int {
	cache := map[int]map[int]int{}
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

func getOxygen(list []string) string {
	bLength := len(list[0])
	tmpList := list
	for i := 0; i < bLength; i++ {
		if len(tmpList) <= 1 {
			return tmpList[0]
		}
		data := extract(tmpList)
		tfm := transform(data)

		newList := filter(func(val string) bool {
			cmp := tfm[i].MostCommon()
			return string(val[i]) == fmt.Sprintf("%d", cmp)
		}, tmpList)
		tmpList = newList
	}
	return tmpList[0]
}
func getC02(list []string) string {
	bLength := len(list[0])
	tmpList := list
	for i := 0; i < bLength; i++ {
		if len(tmpList) == 1 {
			return tmpList[0]
		}
		data := extract(list)
		tfm := transform(data)

		newList := filter(func(val string) bool {
			cmp := tfm[i].LeastCommon()
			return string(val[i]) == fmt.Sprintf("%d", cmp)
		}, tmpList)
		tmpList = newList
	}
	return tmpList[0]
}

func filter(fn predicateFunc, data []string) []string {
	result := []string{}
	for _, item := range data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

type predicateFunc = func(val string) bool
