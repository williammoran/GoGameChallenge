package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const playerSpeed = 2

type Player struct {
	bullet *Bullet
	x      float64
	image  *ebiten.Image
}

func (a *Player) Update() {
	// Accept input from the player and move the player's
	// sprite around or fire the weapan as appropriate
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		a.x -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		a.x += playerSpeed
	}
	// Prevent the player from moving offscreen
	if a.x < radius {
		a.x = radius
	}
	if a.x > screenWidth-radius {
		a.x = screenWidth - radius
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		a.bullet.Fire(a.x, playerVerticalPosition-radius)
	}
}

// Draw draws the player on the screen
func (a *Player) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(a.x-radius, playerVerticalPosition)
	screen.DrawImage(a.image, &op)
}
