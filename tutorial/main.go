package main

import (
	"image"
	"image/color"
	"log"
	"tutorial/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Some globals
var PIXEL_WIDTH = 16
var PIXEL_HEIGHT = 16

var WINDOW_WIDTH = 320
var WINDOW_HEIGHT = 240

// Engine object.
type Game struct {
	player      *entities.Player
	enemies     []*entities.Enemy
	potions     []*entities.Potion
	tilemapJSON *TilemapJSON
	tilemapImg  *ebiten.Image
	cam         *Camera
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

	// Draw Camera with the sprite at the center.
	g.cam.FollowTarget(g.player.X+(float64(PIXEL_WIDTH)/2), g.player.Y+(float64(PIXEL_HEIGHT)/2), float64(WINDOW_WIDTH), float64(WINDOW_HEIGHT))
	g.cam.ConstrainCamera(
		float64(g.tilemapJSON.Layers[0].Width)*float64(PIXEL_WIDTH),
		float64(g.tilemapJSON.Layers[0].Height)*float64(PIXEL_HEIGHT),
		float64(WINDOW_HEIGHT),
		float64(WINDOW_WIDTH))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	// loop over latyers
	for _, layer := range g.tilemapJSON.Layers {
		for index, id := range layer.Data {
			// Get location to print to
			x := (index % layer.Width) * PIXEL_WIDTH
			y := (index / layer.Width) * PIXEL_HEIGHT

			srcX := ((id - 1) % 22) * PIXEL_WIDTH  // fix the 22 later, its the width of the tilemap.tsx
			srcY := ((id - 1) / 22) * PIXEL_HEIGHT // Same

			opts.GeoM.Translate(float64(x), float64(y))

			opts.GeoM.Translate(g.cam.X, g.cam.Y)

			screen.DrawImage(
				g.tilemapImg.SubImage(image.Rect(srcX, srcY, srcX+PIXEL_WIDTH, srcY+PIXEL_HEIGHT)).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}

	opts.GeoM.Translate(g.player.X, g.player.Y)
	opts.GeoM.Translate(g.cam.X, g.cam.Y)

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
		opts.GeoM.Translate(g.cam.X, g.cam.Y)

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
		opts.GeoM.Translate(g.cam.X, g.cam.Y)

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

	tilemapJSON, err := NewTilemapJSON("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}
	tilemapImg, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	// Construct game obj
	game := Game{
		player: &entities.Player{
			Sprite: &entities.Sprite{
				Img: playerImg,
				X:   100,
				Y:   100,
			},
			Health: 10,
		},
		enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Img: skeletonImg,
					X:   0,
					Y:   0,
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &entities.Sprite{
					Img: skeletonImg,
					X:   200,
					Y:   0,
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &entities.Sprite{
					Img: skeletonImg,
					X:   300,
					Y:   0,
				},
				FollowsPlayer: false,
			},
		},
		potions: []*entities.Potion{
			{
				Sprite: &entities.Sprite{
					Img: potionImg,
					X:   200,
					Y:   200,
				},
				AmtHeal: 1,
			},
		},
		tilemapJSON: tilemapJSON,
		tilemapImg:  tilemapImg,
		cam:         NewCamera(0.0, 0.0),
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
