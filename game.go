package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player      *Player
	tileMap     *TileMap
	tilesetData *TilesetData
	tileImage   *ebiten.Image
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	drawTileLayer(
		screen,
		g.tileMap.Layers[0], // first visible layer
		g.tilesetData,       // parsed tileset data
		g.tileImage,         // loaded image
		g.tileMap.TileWidth,
		g.tileMap.TileHeight,
	)

	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
