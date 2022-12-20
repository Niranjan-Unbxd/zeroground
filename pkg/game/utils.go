package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func TextureFromText(r *sdl.Renderer, font *ttf.Font, color sdl.Color, line string) (*sdl.Texture, int32, int32) {
	if font == nil {
		return nil, 0, 0
	}

	s, err := font.RenderUTF8Solid(line, color)
	if err != nil {
		panic(err)
	}
	defer s.Free()
	tex, err := r.CreateTextureFromSurface(s)
	if err != nil {
		panic(err)
	}

	return tex, s.W, s.H
}
