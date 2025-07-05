package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos       rl.Vector2
	Dest      rl.Vector2
	Speed     float32
	Tex       rl.Texture2D
	Frame     int
	FrameTime float32
	Timer     float32
}

func NewPlayer() *Player {
	tex := rl.LoadTexture("assets/player.png") // Assume 3x1 animation strip
	defer rl.UnloadTexture(tex)
	return &Player{
		Pos:       rl.NewVector2(400, 300),
		Dest:      rl.NewVector2(400, 300),
		Speed:     200,
		Tex:       tex,
		Frame:     0,
		FrameTime: 0.2,
	}
}

func (p *Player) Update(dt float32) {
	// Click to move
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		p.Dest = rl.GetMousePosition()
	}

	delta := rl.Vector2Subtract(p.Dest, p.Pos)
	distance := rl.Vector2Length(delta)

	if distance > 2 {
		move := rl.Vector2Scale(rl.Vector2Normalize(delta), p.Speed*dt)
		p.Pos = rl.Vector2Add(p.Pos, move)

		// Animation update
		p.Timer += dt
		if p.Timer > p.FrameTime {
			p.Frame = (p.Frame + 1) % 3
			p.Timer = 0
		}
	}
}

func (p *Player) Draw() {
	frameWidth := float32(p.Tex.Width / 3)
	src := rl.NewRectangle(float32(p.Frame)*frameWidth, 0, frameWidth, float32(p.Tex.Height))
	dst := rl.NewRectangle(p.Pos.X, p.Pos.Y, frameWidth, float32(p.Tex.Height))
	origin := rl.NewVector2(frameWidth/2, float32(p.Tex.Height)/2)
	rl.DrawTexturePro(p.Tex, src, dst, origin, 0, rl.White)
}
