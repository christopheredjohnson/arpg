package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	dirUp = iota
	dirDown
	dirLeft
	dirRight
)

type Player struct {
	Animated
	Speed float64
	State int
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.Speed
		p.Animated.direction = dirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.Speed
		p.Animated.direction = dirRight
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.Speed
		p.Animated.direction = dirUp
		// up only idle only  has 4 frames
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.Speed
		p.Animated.direction = dirDown
	}

	p.Animated.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Animated.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", p.Animated.direction))
}

func newPlayer(img *ebiten.Image) *Player {
	return &Player{
		Animated: Animated{
			Entity: Entity{
				X:      100,
				Y:      100,
				Sprite: img,
			},
			frameOX:     0,
			frameOY:     0,
			frameWidth:  64,
			frameHeight: 64,
			direction:   dirDown,
			AnimationRows: map[int]AnimationRow{
				dirDown:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 10},
				dirLeft:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 10},
				dirRight: {RowIndex: 2, FrameCount: 6, FrameSpeed: 10},
				dirUp:    {RowIndex: 3, FrameCount: 6, FrameSpeed: 10},
			},
		},
		Speed: 2.5,
	}
}
