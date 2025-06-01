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
	paddingX    int // horizontal padding between frames
	paddingY    int // vertical padding between rows
}

func (a *Animated) Update() {
	a.count = (a.count + 1) % (a.frameSpeed * a.frameCount)
}

func (a *Animated) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(a.frameWidth)/2, -float64(a.frameHeight)/2)
	op.GeoM.Translate(a.X, a.Y)

	i := (a.count / a.frameSpeed) % a.frameCount
	sx, sy := a.frameOX, a.frameOY
	if a.Vertical {
		sy += i * (a.frameHeight + a.paddingY)
	} else {
		sx += i * (a.frameWidth + a.paddingX)
	}

	sub := a.Sprite.SubImage(image.Rect(sx, sy, sx+a.frameWidth, sy+a.frameHeight)).(*ebiten.Image)
	screen.DrawImage(sub, op)
}
