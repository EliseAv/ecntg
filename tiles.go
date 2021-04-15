package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const Size = 8

var Background *ebiten.Image = genBlock(color.RGBA{0, 0, 0, 255})

func genBlock(color color.RGBA) *ebiten.Image {
	image := ebiten.NewImage(Size, Size)
	image.Fill(color)
	return image
}
