package game

import (
	"errors"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureResource struct {
	*sdl.Texture
	// width & height
	W, H int32
}

type ResourceManager struct {
	renderer *sdl.Renderer
	textures map[string]TextureResource
}

func (rm *ResourceManager) UnloadAll() {
	for _, t := range rm.textures {
		t.Destroy()
	}
}

func (rm *ResourceManager) GetTexture(key string) (TextureResource, error) {
	t, found := rm.textures[key]
	if found {
		return t, nil
	}
	return t, errors.New("Not found")
}

func (rm *ResourceManager) LoadTexture(key string, path string) (TextureResource, error) {
	var t TextureResource

	s, err := img.Load(path)
	if err != nil {
		return t, err
	}
	s.SetColorKey(true, sdl.MapRGB(s.Format, 0, 18, 25))
	defer s.Free()

	tex, err := rm.renderer.CreateTextureFromSurface(s)
	// this can be used to add tints
	// tex.SetColorMod(127, 0, 0)
	if err != nil {
		return t, err
	}

	if old, found := rm.textures[key]; found {
		old.Destroy()
	}

	t.Texture = tex
	t.W = s.W
	t.H = s.H
	rm.textures[key] = t
	return t, nil
}

func NewResourceManager(r *sdl.Renderer) *ResourceManager {
	return &ResourceManager{
		renderer: r,
		textures: make(map[string]TextureResource, 8),
	}
}
