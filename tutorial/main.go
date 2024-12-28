package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Some globals
var WINDOW_WIDTH = 320
var WINDOW_HEIGHT = 240

// Engine object.
type Game struct {
	PlayerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed((ebiten.KeyRight)) {
		g.X += 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyLeft)) {
		g.X -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyUp)) {
		g.Y -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyDown)) {
		g.Y += 2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.X, g.Y)

	screen.DrawImage(
		g.PlayerImage.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image), // image.Rect can cast to a *ebiten.Image :)
		&opts,
	)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/knight.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{PlayerImage: playerImg, X: 100, Y: 100}); err != nil {
		log.Fatal(err)
	}
}
