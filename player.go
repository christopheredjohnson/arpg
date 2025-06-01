package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Animated
	Speed float64
}

func (p *Player) Update() {
	p.Animated.Update()

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += p.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= p.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= p.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Animated.Draw(screen)
}
