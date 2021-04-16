package model

type Grid struct {
	Size   Size
	Values [][]Cell
}

func NewGrid(size Size) Grid {
	grid := Grid{
		Size:   size,
		Values: make([][]Cell, size.Height),
	}
	for y := 0; y < size.Height; y++ {
		grid.Values[y] = make([]Cell, size.Width)
	}
	return grid
}

func (grid Grid) At(pos Point) Cell {
	return grid.Values[pos.Y][pos.X]
}

func (grid *Grid) SetAt(pos Point, cell Cell) {
	grid.Values[pos.Y][pos.X] = cell
}

func (grid *Grid) Stamp(pos Point, tetromino Tetromino) bool {
	// Can we stamp it here?
	for _, mino := range tetromino.Minos {
		point := pos.Add(mino.Offset)
		if !point.In(grid.Size) || grid.At(point).Exists() {
			return false
		}
	}
	// We can!
	for _, mino := range tetromino.Minos {
		point := pos.Add(mino.Offset)
		grid.SetAt(point, mino.Cell)
	}
	return true
}

func (grid *Grid) Peel(pos Point, tetromino Tetromino) {
	for _, mino := range tetromino.Minos {
		point := pos.Add(mino.Offset)
		grid.SetAt(point, Cell{})
	}
}

func (grid *Grid) ClearLines(pos, amount int) {
	destination := pos + amount - 1
	// Shift lines down
	for source := pos - 1; source >= 0; source-- {
		grid.Values[destination] = grid.Values[source]
		destination--
	}
	// Add new lines at the top
	for ; destination >= 0; destination-- {
		grid.Values[destination] = make([]Cell, grid.Size.Width)
	}
}
