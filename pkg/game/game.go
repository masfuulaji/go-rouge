package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/masfuulaji/go-rogue/pkg/config"
	"github.com/masfuulaji/go-rogue/pkg/systems"
)

type Game struct {
	Tiles []systems.MapTile
}

func NewGame() *Game {
	return &Game{
		Tiles: systems.CreateTiles(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gd := config.NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := g.Tiles[systems.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 800
}
