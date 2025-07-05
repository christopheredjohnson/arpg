package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var game *Game

func main() {
	rl.InitWindow(screenWidth, screenHeight, "ARPG - Diablo Style")
	rl.SetTargetFPS(60)

	game = NewGame()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		game.Update(dt)
		game.Draw()
	}

	rl.CloseWindow()
}
