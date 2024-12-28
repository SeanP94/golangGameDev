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

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

type Potion struct {
	*Sprite
	AmtHeal uint
}

// Engine object.
type Game struct {
	player  *Player
	enemies []*Enemy
	potions []*Potion
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed((ebiten.KeyRight)) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyLeft)) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyUp)) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyDown)) {
		g.player.Y += 2
	}

	for _, sprite := range g.enemies {
		if !sprite.FollowsPlayer {
			break
		}
		if sprite.X < g.player.X {
			sprite.X++
		} else if sprite.X > g.player.X {
			sprite.X--
		}
		if sprite.Y < g.player.Y {
			sprite.Y++
		} else if sprite.Y > g.player.Y {
			sprite.Y--
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image), // image.Rect can cast to a *ebiten.Image :)
		&opts,
	)

	// Reset for the next.
	opts.GeoM.Reset()

	for _, sprite := range g.enemies {
		opts.GeoM.Translate(sprite.X, sprite.Y)
		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image), // image.Rect can cast to a *ebiten.Image :)
			&opts,
		)

		opts.GeoM.Reset()
	}

	for _, sprite := range g.potions {
		opts.GeoM.Translate(sprite.X, sprite.Y)
		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image), // image.Rect can cast to a *ebiten.Image :)
			&opts,
		)

		opts.GeoM.Reset()
	}

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
	potionImg, _, err := ebitenutil.NewImageFromFile("assets/images/lifepot.png")
	if err != nil {
		log.Fatal(err)
	}
	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	// Construct game obj
	game := Game{
		player: &Player{
			&Sprite{
				Img: playerImg,
				X:   100,
				Y:   100,
			},
			10,
		},
		enemies: []*Enemy{
			{
				&Sprite{
					Img: skeletonImg,
					X:   0,
					Y:   0,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImg,
					X:   200,
					Y:   0,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImg,
					X:   300,
					Y:   0,
				},
				false,
			},
		},
		potions: []*Potion{
			{
				&Sprite{
					Img: potionImg,
					X:   200,
					Y:   200,
				},
				1,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
