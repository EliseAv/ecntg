package main

import (
	"ecntg/model"
	"ecntg/view"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var frameCount uint64

type Game struct {
	model model.GameModel
}

func (g *Game) Update() error {
	frameCount++
	if frameCount%20 == 0 || frameCount > 600 {
		g.model.RandomStamp()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	view.DrawGame(screen, &g.model)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return view.Layout(g.model.Grid.Size)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := Game{}
	game.model.NewGame(10, 15)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
