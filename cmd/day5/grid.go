package main

type Position struct {
	X            int
	Y            int
	TimesCrossed int
}

type Diagram struct {
	Grid       [][]Position
	Size       int
	StrictMode bool
}

func NewDiagram(size int, strict bool) Diagram {
	grid := make([][]Position, size)
	for idx := range grid {
		for i := 0; i < size; i++ {
			grid[idx] = append(grid[idx], Position{})
		}
	}
	return Diagram{
		Grid:       grid,
		Size:       size,
		StrictMode: strict,
	}
}

func (d *Diagram) PlaceLine(start, end Position) {
	if d.StrictMode && !isStraightLine(start, end) {
		return
	}
	d.Grid[start.X][start.Y] = start
	line := generateLine(start, end)
	for _, item := range line {
		d.Grid[item.X][item.Y] = Position{
			X:            item.X,
			Y:            item.Y,
			TimesCrossed: d.Grid[item.X][item.Y].TimesCrossed + 1,
		}
	}
}

func isStraightLine(start, end Position) bool {
	return start.X == end.X || start.Y == end.Y
}

func generateLine(start Position, end Position) []Position {
	var result []Position
	result = append(result, start)
	next := start
	for next != end {
		next = incrementPosition(next, end)
		result = append(result, next)
	}
	return result
}

func incrementPosition(start, end Position) Position {
	x := start.X
	y := start.Y

	if x > end.X {
		x--
	} else if x < end.X {
		x++
	}

	if y > end.Y {
		y--
	} else if y < end.Y {
		y++
	}
	return Position{X: x, Y: y}
}
