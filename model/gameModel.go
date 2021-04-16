package model

import (
	"math/rand"
)

type GameModel struct {
	Grid          Grid
	Current, Next Tetromino
	position      Point
}

func (model *GameModel) NewGame(width, height int) {
	model.Grid = NewGrid(Size{width, height})
	model.Next = *GetRandomTetromino()
}

func (model *GameModel) MoveDown() {
	if model.Current.Exists() {
		// first just try to move down
		model.Grid.Peel(model.position, model.Current)
		nextPosition := model.position.Down(&model.Grid.Size)
		if nextPosition != nil && model.Grid.Stamp(*nextPosition, model.Current) {
			// move succeeded!
			model.position = *nextPosition
		} else {
			// move failed, return to previous position permanently!
			model.Grid.Stamp(model.position, model.Current)
			// TODO: Check for lines
			model.Current = Tetromino{}
		}
	} else {
		// Pull next tetromino
		model.Current = model.Next
		model.Next = *GetRandomTetromino()
		model.position = Point{model.Grid.Size.Width/2 - 2, 0}
		// TODO: Game over?
		model.Grid.Stamp(model.position, model.Current)
	}
}

func (model *GameModel) RandomStamp() {
	tetromino := GetRandomTetromino()
	position := Point{rand.Intn(model.Grid.Size.Width - 4), rand.Intn(model.Grid.Size.Height - 4)}
	model.Grid.Stamp(position, *tetromino)
}
