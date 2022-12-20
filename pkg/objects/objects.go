package objects

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Drawer interface {
	Draw(*sdl.Renderer)
}

type Object interface {
	Position() sdl.Point
	Update()
	Draw(*sdl.Renderer)
}
