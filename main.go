package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Entity struct {
	X, Y              float64
	Sprite            *ebiten.Image
	Health, MaxHealth int
}

type Game struct {
	Player *Player
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")

	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		Player: &Player{
			Animated: Animated{
				Entity: Entity{
					X:      100,
					Y:      100,
					Sprite: playerImg,
				},
				frameOX:     0,
				frameOY:     0,
				frameWidth:  16,
				frameHeight: 16,
				frameCount:  2,
				frameSpeed:  10,
				Vertical:    true,
			},
			Speed: 1.0,
		},
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
