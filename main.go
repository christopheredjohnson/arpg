package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	player *Player
	camera rl.Camera2D
)

func main() {
	// Initialize window
	rl.InitWindow(screenWidth, screenHeight, "2D ARPG - Raylib + Go")
	rl.SetTargetFPS(60)

	// Load player texture (placeholder)
	playerTexture := rl.LoadTexture("assets/GUMDROP.E64.R.PNG")
	defer rl.UnloadTexture(playerTexture)

	player = NewPlayer(playerTexture)
	player.Position = rl.Vector2{X: screenWidth / 2, Y: screenHeight / 2}
	// Player state
	playerSpeed := float32(200)

	camera = rl.Camera2D{}
	camera.Zoom = 1.0

	for !rl.WindowShouldClose() {
		// === Update ===
		// Basic WASD movement
		if rl.IsKeyDown(rl.KeyW) {
			player.Position.Y -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) {
			player.Position.Y += playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyA) {
			player.Position.X -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.Position.X += playerSpeed * rl.GetFrameTime()
		}

		// Update camera to follow player
		camera.Target = rl.Vector2{
			X: player.Position.X,
			Y: player.Position.Y,
		}
		camera.Offset = rl.Vector2{X: screenWidth / 2, Y: screenHeight / 2}

		// === Draw ===
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		player.Draw()
		rl.EndMode2D()

		rl.DrawText("WASD to move", 10, 10, 20, rl.LightGray)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
