package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	// set
	f, err := os.OpenFile("arpg.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	walkImg, _, err := ebitenutil.NewImageFromFile("assets/images/player/Sword_Walk_full.png")

	if err != nil {
		log.Fatal(err)
	}

	attackImg, _, err := ebitenutil.NewImageFromFile("assets/images/player/Sword_attack_full.png")

	if err != nil {
		log.Fatal(err)
	}

	idleImg, _, err := ebitenutil.NewImageFromFile("assets/images/player/Sword_Idle_full.png")

	if err != nil {
		log.Fatal(err)
	}

	p, err := newPlayer(
		walkImg,
		attackImg,
		idleImg,
	)

	if err != nil {
		log.Fatal(err)
	}

	mapPath := "assets/maps/spawn.json"
	tileMap, err := loadTileMap(mapPath)

	if err != nil {
		log.Fatal(err)
	}

	tilesetRef := tileMap.Tilesets[0] // first tileset reference
	cp := cleanPath(tilesetRef.Source)
	tilesetData, err := loadTilesetData(path.Join("assets/maps/", cp))

	if err != nil {
		log.Fatal(err)
	}
	cp = cleanPath(tilesetData.Image)
	tilesetImage, err := loadTilesetImage(path.Join("assets/maps/tilesets/", cp))

	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		Player:      p,
		tileMap:     tileMap,
		tileImage:   tilesetImage,
		tilesetData: tilesetData,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func cleanPath(raw string) string {
	// Replace Tiled's escaped slashes and normalize
	fixed := filepath.Clean(strings.ReplaceAll(raw, "\\", "/"))
	return fixed
}
