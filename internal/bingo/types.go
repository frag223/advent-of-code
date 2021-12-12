package bingo

type Direction string

const (
	Row    Direction = "row"
	Column Direction = "column"
)

type Number struct {
	Number   int
	Selected bool
}

type WinningNumbers struct {
	Numbers   []Number
	Direction Direction
}
