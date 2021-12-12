package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/frag223/advent-of-code/internal/bingo"
)

//go:embed test.txt
var input string

func main() {
	list := strings.Split(input, "\n")

	boardsRawList := list[1:]
	boards := generateBoardsFromList(boardsRawList)

	numbersToCall := strings.Split(list[0], ",")

	var winner bingo.Board
	winningNumber := 0

numbersToCallLoop:
	for _, rawNumber := range numbersToCall {
		for _, board := range boards {
			number, _ := strconv.Atoi(rawNumber)
			board.Mark(number)
			if board.IsBingo() {
				winner = board
				winningNumber = number
				break numbersToCallLoop
			}
		}
	}

	fmt.Println(winner.Grid[0])
	fmt.Println(winningNumber)
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
		numbersString := strings.Split(row, " ")
		for _, item := range numbersString {
			number, _ := strconv.Atoi(item)
			b.Add(number)
		}

	}
	return b
}
