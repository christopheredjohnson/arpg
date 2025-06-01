package main

import (
	"encoding/json"
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileMap struct {
	Height     int                `json:"height"`
	Width      int                `json:"width"`
	TileWidth  int                `json:"tilewidth"`
	TileHeight int                `json:"tileheight"`
	Tilesets   []TilesetReference `json:"tilesets"`
	Layers     []Layer            `json:"layers"`
}

type TilesetReference struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

type TilesetData struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Columns     int    `json:"columns"`
	TileCount   int    `json:"tilecount"`
	TileWidth   int    `json:"tilewidth"`
	TileHeight  int    `json:"tileheight"`
	ImageWidth  int    `json:"imagewidth"`
	ImageHeight int    `json:"imageheight"`
}

type Layer struct {
	Data   []int  `json:"data"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func loadTilesetData(path string) (*TilesetData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tileset TilesetData
	err = json.Unmarshal(data, &tileset)
	if err != nil {
		return nil, err
	}
	return &tileset, nil
}

func loadTileMap(path string) (*TileMap, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	var tileMap TileMap
	err = json.Unmarshal(data, &tileMap)

	fmt.Println(tileMap)

	if err != nil {
		return nil, err
	}
	return &tileMap, nil
}

func loadTilesetImage(path string) (*ebiten.Image, error) {

	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		return nil, err
	}

	return img, nil
}

func drawTileLayer(screen *ebiten.Image, layer Layer, tileset *TilesetData, tilesetImage *ebiten.Image, mapTileWidth, mapTileHeight int) {
	tileWidth := tileset.TileWidth
	tileHeight := tileset.TileHeight
	columns := tileset.Columns

	for i := len(layer.Data) - 1; i >= 0; i-- {
		tileID := layer.Data[i]
		if tileID == 0 {
			continue // empty tile
		}

		tileID -= 1 // because tile IDs in Tiled are 1-based

		sx := (tileID % columns) * tileWidth
		sy := (tileID / columns) * tileHeight

		tile := tilesetImage.SubImage(image.Rect(sx, sy, sx+tileWidth, sy+tileHeight)).(*ebiten.Image)

		x := (i % layer.Width) * mapTileWidth
		y := (i / layer.Width) * mapTileHeight

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(tile, op)
	}
}
