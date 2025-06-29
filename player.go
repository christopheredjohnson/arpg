package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position rl.Vector2
	Texture  rl.Texture2D
}

func NewPlayer(texture rl.Texture2D) *Player {
	return &Player{
		Texture: texture,
	}
}

func (p *Player) Draw() {
	source := rl.NewRectangle(0, 0, 64, 64)
	rl.DrawTextureRec(p.Texture, source, rl.Vector2{
		X: p.Position.X - source.Width/2,
		Y: p.Position.Y - source.Height/2,
	}, rl.White)
}
