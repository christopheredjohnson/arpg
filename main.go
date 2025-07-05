package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FrameWidth   = 64
	FrameHeight  = 64
	FramesPerRow = 10
)

type Animation struct {
	Row        int
	FrameCount int
}

type Player struct {
	Pos            rl.Vector2
	Dest           rl.Vector2
	Speed          float32
	Tex            rl.Texture2D
	Frame          int
	FrameTime      float32
	Timer          float32
	AnimKey        string
	Animations     map[string]Animation
	IsAttacking    bool
	Facing         string  // "down", "up", "left", "right"
	AttackIndex    int     // which attack animation in the chain
	QueuedAttack   bool    // whether player pressed again during current attack
	AttackCooldown float32 // optional cooldown per combo hit
	AttackTimer    float32 // track time within attack
}

func NewPlayer() *Player {
	tex := rl.LoadTexture("assets/player.png")

	animationMap := map[string]Animation{
		"idle_down":  {0, 8},
		"idle_right": {1, 8},
		"idle_up":    {2, 8},
		"idle_left":  {3, 8},

		"walk_down":  {4, 8},
		"walk_right": {5, 8},
		"walk_up":    {6, 8},
		"walk_left":  {7, 8},

		"attack_down":  {12, 6},
		"attack_right": {13, 6},
		"attack_up":    {14, 6},
		"attack_left":  {15, 6},

		"attack_down_1":  {12, 6},
		"attack_right_1": {13, 6},
		"attack_up_1":    {14, 6},
		"attack_left_1":  {15, 6},

		"attack_down_2":  {20, 10},
		"attack_right_2": {21, 10},
		"attack_up_2":    {22, 10},
		"attack_left_2":  {23, 10},
	}

	return &Player{
		Pos:        rl.NewVector2(400, 300),
		Dest:       rl.NewVector2(400, 300),
		Speed:      200,
		Tex:        tex,
		FrameTime:  69.0 / 1000.0,
		AnimKey:    "idle_down",
		Animations: animationMap,
	}
}

func (p *Player) Update(dt float32) {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		p.Dest = rl.GetMousePosition()
	}

	// Start or queue attack
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		if p.IsAttacking {
			// Queue combo follow-up
			p.QueuedAttack = true
		} else {
			// Start attack from beginning
			p.AttackIndex = 1
			p.IsAttacking = true
			p.QueuedAttack = false
			p.Frame = 0
			p.Timer = 0
			p.AttackTimer = 0
			p.AnimKey = "attack_" + p.Facing + "_1"
		}
	}

	if p.IsAttacking {
		p.Timer += dt
		p.AttackTimer += dt

		if p.Timer > p.FrameTime {
			p.Frame++
			p.Timer = 0
		}

		// Get current animation
		anim := p.Animations[p.AnimKey]
		if p.Frame >= anim.FrameCount {
			if p.QueuedAttack && p.AttackIndex < 2 {
				p.AttackIndex++
				p.AnimKey = "attack_" + p.Facing + "_" + itoa(p.AttackIndex)
				p.Frame = 0
				p.Timer = 0
				p.AttackTimer = 0
				p.QueuedAttack = false
			} else {
				// End of attack sequence
				p.IsAttacking = false
				p.AttackIndex = 0
				p.AnimKey = "idle_" + p.Facing
				p.Frame = 0
			}
		}

		return // skip movement during attack
	}

	// Movement
	delta := rl.Vector2Subtract(p.Dest, p.Pos)
	distance := rl.Vector2Length(delta)

	if distance > 2 {
		move := rl.Vector2Scale(rl.Vector2Normalize(delta), p.Speed*dt)
		p.Pos = rl.Vector2Add(p.Pos, move)

		angle := rl.Vector2Angle(rl.NewVector2(1, 0), delta) * (180.0 / rl.Pi)
		if delta.Y < 0 {
			angle = 360 - angle
		}

		if abs(delta.X) > abs(delta.Y) {
			if delta.X > 0 {
				p.AnimKey = "walk_right"
				p.Facing = "right"
			} else {
				p.AnimKey = "walk_left"
				p.Facing = "left"
			}
		} else {
			if delta.Y > 0 {
				p.AnimKey = "walk_down"
				p.Facing = "down"
			} else {
				p.AnimKey = "walk_up"
				p.Facing = "up"
			}
		}

		p.Timer += dt
		if p.Timer > p.FrameTime {
			p.Frame = (p.Frame + 1) % p.Animations[p.AnimKey].FrameCount
			p.Timer = 0
		}
	} else {
		switch p.AnimKey {
		case "walk_down":
			p.AnimKey = "idle_down"
		case "walk_up":
			p.AnimKey = "idle_up"
		case "walk_left":
			p.AnimKey = "idle_left"
		case "walk_right":
			p.AnimKey = "idle_right"
		}

		// Animate idle
		p.Timer += dt
		if p.Timer > p.FrameTime {
			p.Frame = (p.Frame + 1) % p.Animations[p.AnimKey].FrameCount
			p.Timer = 0
		}
	}
}

func (p *Player) Draw() {
	anim, ok := p.Animations[p.AnimKey]
	if !ok {
		println("[ERROR] Missing animation key:", p.AnimKey)
		p.IsAttacking = false
		p.AnimKey = "idle_down"
		return
	}
	src := rl.NewRectangle(
		float32(p.Frame*FrameWidth),
		float32(anim.Row*FrameHeight),
		FrameWidth,
		FrameHeight,
	)
	dst := rl.NewRectangle(p.Pos.X, p.Pos.Y, FrameWidth, FrameHeight)
	origin := rl.NewVector2(FrameWidth/2, FrameHeight/2)
	rl.DrawTexturePro(p.Tex, src, dst, origin, 0, rl.White)
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(600)
	rl.InitWindow(screenWidth, screenHeight, "Player Animation System")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := NewPlayer()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		player.Update(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		player.Draw()

		rl.EndDrawing()
	}
}

func abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
