package objects

import (
	"math/rand"
	"zeroground/pkg/ui"

	"github.com/veandco/go-sdl2/sdl"
)

type simpleObject struct {
	Size     int32
	Pos      sdl.Point
	Color    sdl.Color
	Lifespan uint64

	createdAt uint64
}

func (so *simpleObject) Position() sdl.Point {
	return so.Pos
}

func (so *simpleObject) Rect() []sdl.Rect {
	rects := make([]sdl.Rect, 1)
	return append(rects, sdl.Rect{
		X: so.Pos.X - so.Size/2,
		Y: so.Pos.Y - so.Size/2,
		W: so.Size,
		H: so.Size,
	})
}

func (so *simpleObject) Handle(sdl.Event) {}

func (so *simpleObject) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(so.Color.R, so.Color.G, so.Color.B, sdl.ALPHA_TRANSPARENT)
	for _, rect := range so.Rect() {
		renderer.FillRect(&rect)
	}
}

func (so *simpleObject) Reset() {
	so.Pos.X = rand.Int31() % (ui.WindowWidth - so.Size)
	so.Pos.Y = rand.Int31() % (ui.WindowHeight - so.Size)
	so.createdAt = sdl.GetTicks64()
}

func (so *simpleObject) Update() {
	if sdl.GetTicks64()-so.createdAt >= so.Lifespan {
		so.Reset()
	}
}

func NewSimpleObject(size int32, color sdl.Color, lifespan uint64) Object {
	so := &simpleObject{
		Color:    color,
		Size:     size,
		Lifespan: lifespan,
	}
	so.Reset()
	return so
}
