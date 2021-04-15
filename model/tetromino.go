package model

import (
	"math/rand"
	"strings"
)

type Tetromino struct {
	Cells      [4][4]*Cell
	Prev, Next *Tetromino
}

var universe []*Tetromino = genUniverse()

var rawData = `
---- ---- ---- --*- 
-*-- -**- ***- --*-
-*** -*-- --*- -**-
---- -*-- ---- ----
/
---- -*-- ---- ----
--*- -*-- -*** -**-
***- -**- -*-- --*-
---- ---- ---- --*- 
/
--*- --*- ---- --*-
-*** --** -*** -**-
---- --*- --*- --*-
---- ---- ---- ----
/
---- --*-
---- --*-
**** --*-
---- --*-
/
----
-**-
-**-
----
/
---- -*--
--** -**-
-**- --*-
---- ----
/
---- --*-
-**- -**-
--** -*--
---- ----
`

func genUniverse() []*Tetromino {
	rotatingSets := strings.Split(rawData, "/")
	result := make([]*Tetromino, len(rotatingSets))
	for iSet, set := range rotatingSets {
		result[iSet] = genTetrominoSet(strings.TrimSpace(set))
	}
	return result
}

func genTetrominoSet(set string) *Tetromino {
	lines := strings.Split(set, "\n")
	setSize := len(strings.Split(lines[0], " "))
	pieces := make([]Tetromino, setSize)
	// Build linked list
	for i, piece := range pieces {
		piece.Next = &pieces[(i+1)%setSize]
		piece.Next.Prev = &piece
	}
	// Transcribe characters
	for y, line := range lines {
		for i, row := range strings.Split(line, " ") {
			for x, char := range row {
				if char == '*' {
					pieces[i].Cells[y][x] = &Cell{}
				}
			}
		}
	}
	// Yield head of linked list
	return &pieces[0]
}

func GetRandomTetromino() *Tetromino {
	index := rand.Intn(len(universe))
	return universe[index]
}
