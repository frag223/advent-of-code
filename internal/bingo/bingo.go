package bingo

import (
	"fmt"
)

type Board struct {
	Size           int
	Grid           [][]Number
	WinningNumbers WinningNumbers
	GameComplete   bool
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
	fmt.Println(b.Grid)
}

func (b *Board) IsBingo() bool {
	if b.GameComplete {
		return b.GameComplete
	}
	for i := 0; i < b.Size; i++ {
		row := getRow(i, b.Grid)
		column := getColumn(i, b.Grid)
		if isBingo(row) {
			b.WinningNumbers = WinningNumbers{
				Numbers:   row,
				Direction: Row,
			}
			b.GameComplete = true
		}
		if isBingo(column) {
			b.WinningNumbers = WinningNumbers{
				Numbers:   column,
				Direction: Column,
			}
			b.GameComplete = true
		}
	}
	return b.GameComplete
}

func (b *Board) UnselectedNumbers() []Number {
	var result []Number
	for _, row := range b.Grid {
		for _, item := range row {
			if !item.Selected {
				result = append(result, item)
				continue
			}
		}
	}
	return result
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
