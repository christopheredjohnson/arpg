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

func newPlayer(walkImg *ebiten.Image) (*Player, error) {
	return &Player{
		Animated: Animated{
			Entity: Entity{
				X: 100,
				Y: 100,
			},
			frameOX:     0,
			frameOY:     0,
			frameWidth:  64,
			frameHeight: 64,
			direction:   dirDown,
			state:       stateWalk,
			Spritesheets: map[int]*ebiten.Image{
				// stateIdle:   idleSheet,
				stateWalk: walkImg,
				// stateAttack: attackSheet,
			},
			AnimationFrames: map[AnimationKey]AnimationRow{
				{stateWalk, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 6},
				{stateWalk, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 6},
				{stateWalk, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 6},
				{stateWalk, dirUp}:    {RowIndex: 3, FrameCount: 6, FrameSpeed: 6},
			},
		},
		Speed: 2.5,
	}, nil
}
