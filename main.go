package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	swordIdle, _, err := ebitenutil.NewImageFromFile("assets/images/player/Sword_Idle_full.png")

	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		Player: newPlayer(swordIdle),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
