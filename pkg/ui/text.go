package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Basic interface {
	Render(*sdl.Renderer)
}

type Dynamic interface {
	Render(*sdl.Renderer)
	Update()
}

type Interactive interface {
	Render(*sdl.Renderer)
	Handle(sdl.Event)
	Update()
}

type Text struct {
	text    string
	texture *sdl.Texture
	dest    sdl.Rect
}

func (t *Text) Render(r *sdl.Renderer) {
	r.Copy(t.texture, nil, &t.dest)
}

func NewText(x, y int32, text string, color sdl.Color, r *sdl.Renderer, font *ttf.Font) *Text {
	s, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		panic(err)
	}
	defer s.Free()
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		panic(err)
	}

	ui := &Text{
		text:    text,
		texture: t,
		dest:    sdl.Rect{X: x, Y: y, W: s.W, H: s.H},
	}

	return ui
}
