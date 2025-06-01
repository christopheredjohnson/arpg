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
	Inventory
}

func (p *Player) Update() {

	moving := false

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.Speed
		moving = true
		p.Animated.direction = dirUp
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.Speed
		moving = true
		p.Animated.direction = dirDown
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.Speed
		moving = true
		p.Animated.direction = dirLeft
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.Speed
		moving = true
		p.Animated.direction = dirRight
	}

	if moving {
		p.Animated.state = stateWalk
	} else {
		p.Animated.state = stateIdle
	}
	p.Animated.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Animated.Draw(screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", len(p.Inventory.Items)), int(p.X)-6, int(p.Y)-32)
}

func newPlayer(
	walkImg *ebiten.Image,
	attackImg *ebiten.Image,
	idleImg *ebiten.Image,
) (*Player, error) {
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
			state:       stateIdle,
			Spritesheets: map[int]*ebiten.Image{
				stateWalk:   walkImg,
				stateAttack: attackImg,
				stateIdle:   idleImg,
			},
			AnimationFrames: map[AnimationKey]AnimationRow{
				// walking
				{stateWalk, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 15},
				{stateWalk, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 15},
				{stateWalk, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 15},
				{stateWalk, dirUp}:    {RowIndex: 3, FrameCount: 6, FrameSpeed: 15},
				// attacking
				{stateAttack, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 15},
				{stateAttack, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 15},
				{stateAttack, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 15},
				{stateAttack, dirUp}:    {RowIndex: 3, FrameCount: 6, FrameSpeed: 15},
				// idle
				{stateIdle, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 15},
				{stateIdle, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 15},
				{stateIdle, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 15},
				{stateIdle, dirUp}:    {RowIndex: 3, FrameCount: 4, FrameSpeed: 15},
			},
		},
		Speed: 2.5,
		Inventory: Inventory{
			Items: []*Item{},
			Size:  10,
		},
	}, nil
}
