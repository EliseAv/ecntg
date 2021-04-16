package model

import (
	"fmt"
	"math/rand"
	"strings"
)

type Tetromino struct {
	Minos       [4]Mino
	Right, Left *Tetromino
}

type Mino struct {
	Offset Point
	Cell   Cell
}

var universe = genUniverse()

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

func (tetromino Tetromino) String() string {
	i := int(0)
	for _, mino := range tetromino.Minos {
		i <<= 2
		i += mino.Offset.X
		i <<= 2
		i += mino.Offset.Y
	}
	return fmt.Sprintf("%x", i)
}

func genUniverse() []*Tetromino {
	rotatingSets := strings.Split(rawData, "/")
	result := make([]*Tetromino, len(rotatingSets))
	for iSet, set := range rotatingSets {
		tetromino := genTetrominoSet(strings.TrimSpace(set))
		result[iSet] = tetromino
	}
	return result
}

func genTetrominoSet(set string) *Tetromino {
	lines := strings.Split(set, "\n")
	setSize := len(strings.Split(lines[0], " "))
	pieces := make([]Tetromino, setSize)
	minoIndexes := make([]int, setSize)
	// Build linked list
	for i, piece := range pieces {
		piece.Right = &pieces[(i+1)%setSize]
		piece.Right.Left = &piece
	}
	// Transcribe characters
	for y, line := range lines {
		for i, row := range strings.Split(line, " ") {
			for x, char := range row {
				if char == '*' {
					pieces[i].Minos[minoIndexes[i]] = Mino{Point{x, y}, SingleCell}
					minoIndexes[i]++
				}
			}
		}
	}
	// Yield head of linked list
	return &pieces[0]
}

func GetRandomTetromino() *Tetromino {
	index := rand.Intn(len(universe))
	result := universe[index]
	return result
}

func (tetromino Tetromino) Exists() bool {
	return tetromino.Left != nil
}
