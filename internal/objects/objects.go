package objects

import (
	"zeroground/pkg/objects"
	"zeroground/pkg/physics"

	"github.com/veandco/go-sdl2/sdl"
)

type Direction uint8

const (
	Dir_Invalid Direction = iota
	Dir_Left
	Dir_Right
	Dir_Up
	Dir_Down
)

type Snake interface {
	objects.Object
	Hitbox() []physics.Plane2D
	SetDirection(dir Direction)
	HasIntersection(...*sdl.Rect) (int, bool)
	IsAlive() bool
}

type Player interface {
	Snake
	Handle(sdl.Event)
	Eat(food Food)
}

type Food interface {
	objects.Object
	Hitbox() []physics.Plane2D
	State() freshness
}
