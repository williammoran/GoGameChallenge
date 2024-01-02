package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	hSpeed = float64(1)
	vSpeed = float64(0.3)
)

type Baddie struct {
	x, y, direction float64
	image           *ebiten.Image
}

func (a *Baddie) Update() {
	// This initializes the baddie at game startup
	if a.x < radius {
		a.x = radius
		a.y = radius
		a.direction = hSpeed
	}
	// Move
	a.x += a.direction
	a.y += vSpeed
	// Change direction hit either screen edge
	if a.x > screenWidth-radius {
		a.x = screenWidth - radius
		a.direction = -hSpeed
	}
	if a.x < radius {
		a.x = radius
		a.direction = hSpeed
	}
	// If the baddie has landed, game over
	if a.y > playerVerticalPosition {
		fmt.Println("You lose")
		os.Exit(0)
	}
}

// Draw draws the baddie to the screen
func (a *Baddie) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(a.x-radius, a.y-radius)
	screen.DrawImage(a.image, &op)
}

// Die sets the baddie's location back to the start
// to simulate another baddie attacking
func (a *Baddie) Die() {
	a.x = float64(tileSize / 2)
	a.y = float64(tileSize / 2)
	a.direction = hSpeed
}
