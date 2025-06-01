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
	p.Animated.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Animated.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", p.Animated.direction))
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
				{stateWalk, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 6},
				{stateWalk, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 6},
				{stateWalk, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 6},
				{stateWalk, dirUp}:    {RowIndex: 3, FrameCount: 6, FrameSpeed: 6},
				// attacking
				{stateAttack, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 6},
				{stateAttack, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 6},
				{stateAttack, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 6},
				{stateAttack, dirUp}:    {RowIndex: 3, FrameCount: 6, FrameSpeed: 6},
				// idle
				{stateIdle, dirDown}:  {RowIndex: 0, FrameCount: 6, FrameSpeed: 10},
				{stateIdle, dirLeft}:  {RowIndex: 1, FrameCount: 6, FrameSpeed: 10},
				{stateIdle, dirRight}: {RowIndex: 2, FrameCount: 6, FrameSpeed: 10},
				{stateIdle, dirUp}:    {RowIndex: 3, FrameCount: 4, FrameSpeed: 20},
			},
		},
		Speed: 2.5,
	}, nil
}
