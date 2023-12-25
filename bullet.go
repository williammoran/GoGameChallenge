package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	baddie *Baddie
	x, y   float64
	active bool
	image  *ebiten.Image
}

func (a *Bullet) Update() {
	if a.active {
		a.y -= 5
		center := float64(tileSize / 2)
		if a.y-center < 1 {
			a.active = false
			return
		}
		// calculate distance between bullet and baddie
		dx := math.Abs(a.x - a.baddie.x)
		dy := math.Abs(a.y - a.baddie.y)
		dist := math.Sqrt(dx*dx + dy*dy)
		// If the distance is less than the sum of the
		// radii, then they are colliding
		if dist <= 18 {
			a.baddie.Die()
			a.active = false
		}
	}
}

// drawSprite draws this sprite on the screen
func (a *Bullet) Draw(screen *ebiten.Image) {
	if a.active {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(a.x-radius, a.y-radius)
		screen.DrawImage(a.image, &op)
	}
}

// Fire sets the bullet to start moving from the
// specified coordinates
func (a *Bullet) Fire(x, y float64) {
	a.x = x
	a.y = y
	a.active = true
}
