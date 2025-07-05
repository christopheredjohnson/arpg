package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Player *Player
}

func NewGame() *Game {
	return &Game{
		Player: NewPlayer(),
	}
}

func (g *Game) Update(dt float32) {
	g.Player.Update(dt)
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// TODO: Draw map first
	g.Player.Draw()

	rl.EndDrawing()
}
