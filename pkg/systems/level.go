package systems

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/masfuulaji/go-rogue/pkg/config"
)

type MapTile struct {
	Image   *ebiten.Image
	PixelX  int
	PixelY  int
	Blocked bool
}

func GetIndexFromXY(x, y int) int {
	g := config.NewGameData()
	return y*g.ScreenWidth + x
}

func CreateTiles() []MapTile {
	g := config.NewGameData()
	tiles := make([]MapTile, 0)
	for x := 0; x < g.ScreenWidth; x++ {
		for y := 0; y < g.ScreenHeight; y++ {
			if x == 0 || x == g.ScreenWidth-1 || y == 0 || y == g.ScreenHeight-1 {

				wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					Image:   wall,
					PixelX:  x * g.TileWidth,
					PixelY:  y * g.TileHeight,
					Blocked: true,
				}
				tiles = append(tiles, tile)
			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					Image:   floor,
					PixelX:  x * g.TileWidth,
					PixelY:  y * g.TileHeight,
					Blocked: false,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return tiles
}
