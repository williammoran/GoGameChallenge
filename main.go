package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 640
	screenHeight = 480
	tileSize     = 32
	radius       = float64(tileSize / 2)
)

func main() {
	// Create our game object and tell Ebiten to run it
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
