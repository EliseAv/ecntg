package model

type Grid struct {
	Width, Height int
	Values        [][]*Cell
}

func NewGrid(width, height int) Grid {
	grid := Grid{
		Width:  width,
		Height: height,
		Values: make([][]*Cell, height),
	}
	for y := 0; y < height; y++ {
		grid.Values[y] = make([]*Cell, width)
	}
	return grid
}

func (grid *Grid) Stamp(x, y int, tetromino *Tetromino) bool {
	for ty, row := range tetromino.Cells {
		for tx, cell := range row {
			if cell != nil {
				cx := x + tx
				cy := y + ty
				if cx < grid.Width && cy < grid.Height && grid.Values[cy][cx] != nil {
					grid.Values[cy][cx] = cell
				} else {
					return false
				}
			}
		}
	}
	return true
}

func (grid *Grid) Peel(x, y int, tetromino *Tetromino) {
	for ty, row := range tetromino.Cells {
		for tx, cell := range row {
			if cell != nil {
				grid.Values[y+ty][x+tx] = nil
			}
		}
	}
}

func (grid *Grid) ClearLines(pos, amount int) {
	destination := pos + amount - 1
	for source := pos - 1; source >= 0; source-- {
		grid.Values[destination] = grid.Values[source]
		destination--
	}
	for ; destination >= 0; destination-- {
		grid.Values[destination] = make([]*Cell, grid.Width)
	}
}
