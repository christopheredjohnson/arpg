package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationKey struct {
	State     int
	Direction int
}

type AnimationRow struct {
	RowIndex   int
	FrameCount int
	FrameSpeed int
}
type Animated struct {
	Entity
	Vertical        bool
	count           int
	frameOX         int
	frameOY         int
	frameWidth      int
	frameHeight     int
	paddingX        int // horizontal padding between frames
	paddingY        int // vertical padding between rows
	direction       int
	state           int
	AnimationFrames map[AnimationKey]AnimationRow
	Spritesheets    map[int]*ebiten.Image // key = state
}

func (a *Animated) Update() {
	key := AnimationKey{State: a.state, Direction: a.direction}
	row, ok := a.AnimationFrames[key]
	if !ok {
		row = AnimationRow{FrameCount: 1, FrameSpeed: 10}
	}
	a.count = (a.count + 1) % (row.FrameSpeed * row.FrameCount)
}

func (a *Animated) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(a.frameWidth)/2, -float64(a.frameHeight)/2)
	op.GeoM.Translate(a.X, a.Y)

	key := AnimationKey{State: a.state, Direction: a.direction}
	row, ok := a.AnimationFrames[key]
	if !ok {
		row = AnimationRow{RowIndex: 0, FrameCount: 1, FrameSpeed: 10}
	}

	i := (a.count / row.FrameSpeed) % row.FrameCount
	sx := a.frameOX + i*(a.frameWidth+a.paddingX)
	sy := a.frameOY + row.RowIndex*(a.frameHeight+a.paddingY)

	sheet, ok := a.Spritesheets[a.state]
	if !ok || sheet == nil {
		return // fallback: skip drawing if no spritesheet found
	}

	sub := sheet.SubImage(image.Rect(sx, sy, sx+a.frameWidth, sy+a.frameHeight)).(*ebiten.Image)
	screen.DrawImage(sub, op)
}
