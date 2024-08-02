package entities

import (
	"github.com/masfuulaji/go-rogue/pkg/systems"
)

type GameMap struct {
	Dungeons []Dungeon
}

func NewGameMap() GameMap {
	l := systems.NewLevel()

	levels := make([]systems.Level, 0)
	levels = append(levels, *l)
	d := Dungeon{
		Name:   "Dungeon",
		Levels: levels,
	}

	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	return GameMap{
		Dungeons: dungeons,
	}
}
