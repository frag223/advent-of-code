package bingo

import (
	"encoding/json"
	"fmt"
)

type Board struct {
	Size int
	Grid [][]Number
}

type Number struct {
	Number   int
	Selected bool
}

func New(size int) Board {
	grid := make([][]Number, size)
	return Board{
		Size: size,
		Grid: grid,
	}
}

func (b *Board) Add(number int) {

	for index, row := range b.Grid {
		if len(row) >= b.Size {
			continue
		}
		b.Grid[index] = append(b.Grid[index], Number{
			Number:   number,
			Selected: false,
		})
		break
	}
}

func (b *Board) Mark(number int) bool {
	for idx, row := range b.Grid {
		for i, column := range row {
			if column.Number == number {
				b.Grid[idx][i] = Number{
					Number:   column.Number,
					Selected: true,
				}
				return true
			}
		}
	}
	return false
}

func (b *Board) Print() {
	d, _ := json.MarshalIndent(b.Grid, "\n", "  ")
	fmt.Println(string(d))
}

func (b *Board) IsBingo() bool {
	for i := 0; i < 5; i++ {
		row := getRow(i, b.Grid)
		column := getColumn(i, b.Grid)
		if isBingo(row) || isBingo(column) {
			return true
		}
	}
	return false
}

func getRow(rowID int, grid [][]Number) []Number {
	return grid[rowID]
}

func getColumn(columnID int, grid [][]Number) []Number {
	result := []Number{}
	for _, row := range grid {
		result = append(result, row[columnID])
	}
	return result
}

func isBingo(row []Number) bool {
	for _, item := range row {
		if !item.Selected {
			return false
		}
	}
	return true
}
