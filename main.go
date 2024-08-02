package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/masfuulaji/go-rogue/pkg/game"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		panic(err)
	}
}
