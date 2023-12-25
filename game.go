package main

import (
	"bytes"
	"image"
	"image/color"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const playerVerticalPosition = screenHeight - tileSize*2

// These embeds pull the data in from the files

//go:embed baddie.png
var baddiePNG []byte

//go:embed bullet.png
var bulletPNG []byte

//go:embed player.png
var playerPNG []byte

func NewGame() *Game {
	game := Game{}
	// Turn the PNG images into image types that can be drawn
	// and create objects for each
	img, _, err := image.Decode(bytes.NewReader(baddiePNG))
	if err != nil {
		panic(err)
	}
	game.baddie = &Baddie{
		image: ebiten.NewImageFromImage(img),
	}
	img, _, err = image.Decode(bytes.NewReader(bulletPNG))
	if err != nil {
		panic(err)
	}
	game.bullet = &Bullet{
		image:  ebiten.NewImageFromImage(img),
		baddie: game.baddie,
	}
	img, _, err = image.Decode(bytes.NewReader(playerPNG))
	if err != nil {
		panic(err)
	}
	game.player = &Player{
		image:  ebiten.NewImageFromImage(img),
		bullet: game.bullet,
		x:      screenWidth / 2,
	}
	return &game
}

type Game struct {
	baddie *Baddie
	bullet *Bullet
	player *Player
}

// Layout returns the logical screen size. Use a fixed
// size for simplicity
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Update is where all the game logic should occur.
// This includes checking for commands from the player,
// as well as executing any AI decisions
func (g *Game) Update() error {
	g.baddie.Update()
	g.bullet.Update()
	g.player.Update()
	return nil
}

// All we do here is draw to the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	g.baddie.Draw(screen)
	g.bullet.Draw(screen)
	g.player.Draw(screen)
}
