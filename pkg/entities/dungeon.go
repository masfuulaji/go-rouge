package entities

import "github.com/masfuulaji/go-rogue/pkg/systems"

type Dungeon struct {
	Name   string
	Levels []systems.Level
}
