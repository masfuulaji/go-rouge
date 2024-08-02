package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/masfuulaji/go-rogue/pkg/config"
	"github.com/masfuulaji/go-rogue/pkg/entities"
)

type Game struct {
	Map entities.GameMap
}

func NewGame() *Game {
	return &Game{
		Map: entities.NewGameMap(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gd := config.NewGameData()
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	gd := config.NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}
