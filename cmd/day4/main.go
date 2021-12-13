package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/frag223/advent-of-code/internal/bingo"
)

type Winner struct {
	Board         bingo.Board
	WinningNumber int
}

//go:embed input.txt
var input string

func main() {
	list := strings.Split(input, "\n")

	boardsRawList := list[1:]
	boards := generateBoardsFromList(boardsRawList)

	numbersToCall := strings.Split(list[0], ",")

	var winners []Winner

	for _, rawNumber := range numbersToCall {
		number, _ := strconv.Atoi(rawNumber)
		for idx, board := range boards {
			if board.GameComplete {
				continue
			}
			board.Mark(number)
			if board.IsBingo() {
				winners = append(winners, Winner{
					Board:         board,
					WinningNumber: number,
				})
				boards[idx] = board
			}
		}
	}

	lastWinner := winners[len(winners)-1]

	sumUnpickedNumber := Sum(lastWinner.Board.UnselectedNumbers())
	fmt.Println("------")
	fmt.Println(sumUnpickedNumber)
	fmt.Println(lastWinner.WinningNumber)
	fmt.Println(sumUnpickedNumber * lastWinner.WinningNumber)
}

func generateBoardsFromList(list []string) []bingo.Board {
	var result [][]string
	tmpBoard := []string{}
	for _, row := range list {
		if len(row) == 0 && len(tmpBoard) == 0 {
			continue
		}
		if len(row) == 0 && len(tmpBoard) > 0 {
			result = append(result, tmpBoard)
			tmpBoard = []string{}
			continue
		}
		tmpBoard = append(tmpBoard, row)
	}
	var boards []bingo.Board
	for _, board := range result {
		b := buildBoard(board)
		boards = append(boards, b)
	}
	return boards
}

func buildBoard(boardRaw []string) bingo.Board {
	b := bingo.New(len(boardRaw))
	for _, row := range boardRaw {
		numbersString := strings.Split(strings.TrimSpace(row), " ")
		for _, item := range numbersString {
			if item == "" {
				continue
			}
			number, _ := strconv.Atoi(item)
			b.Add(number)
		}

	}
	return b
}

func Sum(list []bingo.Number) int {
	var result int
	for _, item := range list {
		result += item.Number
	}
	return result
}
