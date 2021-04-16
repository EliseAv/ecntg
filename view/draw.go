package view

import (
	"ecntg/model"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const cellSize = 8

var imagePoint = genImagePoint()

func Layout(board model.Size) (width, height int) {
	width = (board.Width + 6) * cellSize
	height = (board.Height + 2) * cellSize
	return
}

func DrawGame(screen *ebiten.Image, game *model.GameModel) {
	count := 0
	for y, row := range game.Grid.Values {
		for x, cell := range row {
			if cell.Exists() {
				options := ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(x+2)*cellSize, float64(y+1)*cellSize)
				screen.DrawImage(imagePoint, &options)
				count++
			}
		}
	}
	ebitenutil.DebugPrint(screen, "Cells="+fmt.Sprint(count))
}

func genImagePoint() *ebiten.Image {
	const b = cellSize - 1
	image := ebiten.NewImage(cellSize, cellSize)
	image.Fill(hexColor(0x2188b5))
	light := hexColor(0x21b54d)
	dark := hexColor(0x3721b5)
	for i := 0; i < cellSize; i++ {
		image.Set(b, b-i, dark)
		image.Set(0, i, light)
		image.Set(i, 0, light)
		image.Set(b-i, b, dark)
	}
	return image
}

func hexColor(hex int) color.RGBA {
	blue := byte(hex)
	hex >>= 8
	green := byte(hex)
	hex >>= 8
	red := byte(hex)
	return color.RGBA{red, green, blue, 255}
}
