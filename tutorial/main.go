package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Some globals
var WINDOW_WIDTH = 320
var WINDOW_HEIGHT = 240

// Engine object.
type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, "Hello, World!")

	screen.Fill(color.RGBA{120, 180, 255, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
