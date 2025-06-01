package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Animated
	Speed float64
	State int
}

func (p *Player) Update() {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.Speed
	}

	p.Animated.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Animated.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", p.State))
}

func newPlayer(img *ebiten.Image) *Player {
	return &Player{
		Animated: Animated{
			Entity: Entity{
				X:      100,
				Y:      100,
				Sprite: img,
			},
			frameOX:     16,
			frameOY:     16,
			frameWidth:  32,
			frameHeight: 32,
			frameCount:  12,
			frameSpeed:  5,
		},
		Speed: 2.5,
	}
}
