package objects

import (
	"zeroground/physics"

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

type Drawer interface {
	Draw(*sdl.Renderer)
}

type Object interface {
	Position() sdl.Point
	Update()
	Draw(*sdl.Renderer)
	// Rect() []sdl.Rect
}

type Snake interface {
	Object
	Hitbox() []physics.Plane2D
	SetDirection(dir Direction)
}

type Player interface {
	Snake
	Handle(sdl.Event)
	Eat(food Food)
}

type Food interface {
	Object
	Hitbox() []physics.Plane2D
	State() freshness
}

// Done: enemies which can move -> Tower
// Done: food & poison reset after fixed time (food, poison and tower)
// Done: snake implements Snake interface
// Done: player is a snake controlled by inputs
// Done: make small snakes autonamous
// Done: snakes should die/reset when hit boundries
// Done: small snake enemies spawns from spawner
// Done: Snake can speed up, fix bug where it speeds up instantly
// Done: food turns stale then rotten
//	1. fresh food is good and gives +1 body
//	2. stale food no benifits
//	3. a new food spawns as soon as current food goes stale or a fixed time is passed
// 	4. rotten food kills
//	5. invalid foods should be ignored for all operations and removed ASAP

// IN PROGRESS:
// TODO: food should move (imagine rabbits or mice) - ???
// TODO: venom upgrades like (increase initial damage, dps, fire speed, range, cooldown, etc) - ???
