package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {

	swordIdle, _, err := ebitenutil.NewImageFromFile("assets/images/player/Sword_Walk_full.png")

	if err != nil {
		log.Fatal(err)
	}

	p, err := newPlayer(swordIdle)

	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		Player: p,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
