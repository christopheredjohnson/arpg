package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animated struct {
	Entity
	Vertical    bool
	count       int
	frameOX     int
	frameOY     int
	frameWidth  int
	frameHeight int
	frameCount  int
	frameSpeed  int
}

func (a *Animated) Update() {
	a.count = (a.count + 1) % (a.frameSpeed * a.frameCount)
}

func (a *Animated) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(a.frameWidth)/2, -float64(a.frameHeight)/2)
	op.GeoM.Translate(a.X, a.Y)

	i := (a.count / a.frameSpeed) % a.frameCount
	var sx, sy int

	if a.Vertical {
		sx = a.frameOX
		sy = a.frameOY + i*a.frameHeight
	} else {
		sx = a.frameOX + i*a.frameWidth
		sy = a.frameOY
	}

	sub := a.Sprite.SubImage(image.Rect(sx, sy, sx+a.frameWidth, sy+a.frameHeight)).(*ebiten.Image)
	screen.DrawImage(sub, op)
}
